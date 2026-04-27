package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/dto"
)

var (
	// ErrInvalidReportType indicates unsupported `type` query.
	ErrInvalidReportType = errors.New("jenis report tidak valid")
	// ErrInvalidDateFormat indicates unsupported date format.
	ErrInvalidDateFormat = errors.New("format tanggal tidak valid")
	// ErrInvalidDateRange indicates from date is after to date.
	ErrInvalidDateRange = errors.New("rentang tanggal tidak valid")
	// ErrInvalidReportMode indicates unsupported mode query.
	ErrInvalidReportMode = errors.New("mode report tidak valid: gunakan ai atau raw")
)

var reportCatalog = map[string]struct {
	title       string
	description string
	chartType   string
}{
	"daily-sales":        {title: "Penjualan Harian", description: "Ringkasan tren omzet harian", chartType: "line"},
	"monthly-sales":      {title: "Penjualan Bulanan", description: "Ringkasan tren omzet bulanan", chartType: "line"},
	"top-products":       {title: "Produk Terlaris", description: "Produk dengan kontribusi penjualan tertinggi", chartType: "bar"},
	"sales-by-person":    {title: "Penjualan per Sales", description: "Kontribusi omzet per sales", chartType: "bar"},
	"order-funnel":       {title: "Funnel Order", description: "Distribusi status order", chartType: "funnel"},
	"category-breakdown": {title: "Penjualan per Kategori", description: "Kontribusi kategori terhadap omzet", chartType: "pie"},
	"low-stock":          {title: "Stok Rendah", description: "Produk dengan stok di bawah ambang", chartType: "bar"},
	"revenue-trend":      {title: "Tren Pendapatan", description: "Tren revenue per periode", chartType: "line"},
}

// ReportService handles dashboard report generation and AI enrichment.
type ReportService interface {
	Generate(ctx context.Context, req dto.ReportRequest) (*dto.ReportResponse, error)
}

type reportService struct {
	httpClient *http.Client
}

// NewReportService creates ReportService with timeout-safe HTTP client.
func NewReportService() ReportService {
	return &reportService{
		httpClient: &http.Client{
			Timeout: 8 * time.Second,
		},
	}
}

func (s *reportService) Generate(ctx context.Context, req dto.ReportRequest) (*dto.ReportResponse, error) {
	reportType := strings.TrimSpace(req.Type)
	if reportType == "" {
		reportType = "daily-sales"
	}
	mode := strings.ToLower(strings.TrimSpace(req.Mode))
	if mode == "" {
		mode = "ai"
	}
	if mode != "ai" && mode != "raw" {
		return nil, ErrInvalidReportMode
	}

	meta, ok := reportCatalog[reportType]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrInvalidReportType, reportType)
	}

	from, to, err := parseDateRange(req.From, req.To)
	if err != nil {
		return nil, err
	}

	report := s.buildDummyReport(reportType, meta, from, to)
	if mode == "raw" {
		report.AISource = "raw"
		report.FallbackReason = ""
		report.Summary = ""
		report.Anomalies = nil
		report.Recommendation = ""
		return report, nil
	}

	ai, err := s.callN8N(ctx, req, report)
	if err != nil {
		report.AISource = "fallback"
		report.FallbackReason = err.Error()
		return report, nil
	}

	if ai.ChartType != "" {
		report.ChartType = ai.ChartType
	}
	if ai.Summary != "" {
		report.Summary = ai.Summary
	}
	if len(ai.Anomalies) > 0 {
		report.Anomalies = normalizeAnomalies(ai.Anomalies)
	}
	if ai.Recommendation != "" {
		report.Recommendation = ai.Recommendation
	}
	report.AISource = "n8n"
	report.FallbackReason = ""

	return report, nil
}

func parseDateRange(fromStr, toStr string) (time.Time, time.Time, error) {
	now := time.Now()
	loc := now.Location()

	if strings.TrimSpace(toStr) == "" {
		to := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
		from := to.AddDate(0, 0, -6)
		return from, to, nil
	}

	to, err := parseDateInput(toStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: gunakan YYYY-MM-DD atau RFC3339", ErrInvalidDateFormat)
	}

	if strings.TrimSpace(fromStr) == "" {
		from := to.AddDate(0, 0, -6)
		return from, to, nil
	}

	from, err := parseDateInput(fromStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: gunakan YYYY-MM-DD atau RFC3339", ErrInvalidDateFormat)
	}

	if from.After(to) {
		return time.Time{}, time.Time{}, fmt.Errorf("%w: from tidak boleh lebih besar dari to", ErrInvalidDateRange)
	}

	return from, to, nil
}

func parseDateInput(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	layouts := []string{
		"2006-01-02",
		time.RFC3339,
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			// Normalize date-only inputs to start/end of day by caller logic.
			if layout == "2006-01-02" {
				return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local), nil
			}
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", s)
}

func (s *reportService) buildDummyReport(reportType string, meta struct {
	title       string
	description string
	chartType   string
}, from, to time.Time) *dto.ReportResponse {
	days := int(to.Sub(from).Hours()/24) + 1
	if days < 1 {
		days = 1
	}
	if days > 31 {
		days = 31
	}

	var totalRevenue float64
	totalOrders := 0
	completedOrders := 0
	pendingOrders := 0

	daily := make([]dto.DailySalesData, 0, days)
	for i := 0; i < days; i++ {
		d := from.AddDate(0, 0, i)
		revenue := float64(1800000 + ((i*170000 + i*i*25000) % 900000))
		totalRevenue += revenue
		orders := 18 + (i % 9)
		totalOrders += orders
		completedOrders += orders - 3
		pendingOrders += 3
		daily = append(daily, dto.DailySalesData{
			Date:  d.Format("2006-01-02"),
			Value: revenue,
		})
	}

	avgOrderValue := 0.0
	if totalOrders > 0 {
		avgOrderValue = totalRevenue / float64(totalOrders)
	}

	metrics := map[string]interface{}{
		"total_revenue":    totalRevenue,
		"total_orders":     totalOrders,
		"completed_orders": completedOrders,
		"pending_orders":   pendingOrders,
		"avg_order_value":  avgOrderValue,
		"period_days":      days,
	}

	resp := &dto.ReportResponse{
		Type:           reportType,
		Title:          meta.title,
		Description:    meta.description,
		ChartType:      meta.chartType,
		Metrics:        metrics,
		Summary:        "Performa penjualan stabil dengan tren naik bertahap. Produk kategori atasan masih menjadi kontributor utama omzet.",
		Recommendation: "Jaga ketersediaan stok produk terlaris dan evaluasi promo untuk kategori dengan pertumbuhan rendah.",
		Anomalies: []dto.Anomaly{
			{
				Metric:      "daily_revenue",
				Actual:      daily[0].Value,
				Expected:    daily[0].Value * 1.18,
				VariancePct: -15.25,
				Description: "Pendapatan hari pertama periode berada di bawah baseline mingguan.",
			},
		},
		AISource: "fallback",
	}

	switch reportType {
	case "daily-sales":
		resp.Data = daily
	case "monthly-sales":
		resp.Data = []dto.MonthlySalesData{
			{Month: "Januari 2026", Value: 92000000},
			{Month: "Februari 2026", Value: 98000000},
			{Month: "Maret 2026", Value: 101000000},
			{Month: "April 2026", Value: 108500000},
		}
	case "top-products":
		resp.Data = []dto.TopProductData{
			{ProductName: "Kemeja Batik Lengan Panjang", Quantity: 132, Revenue: 24420000},
			{ProductName: "Kaos Polos Premium", Quantity: 258, Revenue: 19350000},
			{ProductName: "Celana Chino Slim Fit", Quantity: 88, Revenue: 19360000},
		}
	case "sales-by-person":
		resp.Data = []dto.SalesPersonData{
			{SalesName: "Citra Sales", Revenue: 46500000, Orders: 134},
			{SalesName: "Rudi Sales", Revenue: 39200000, Orders: 109},
			{SalesName: "Nina Sales", Revenue: 31800000, Orders: 87},
		}
	case "order-funnel":
		resp.Data = []dto.OrderFunnelData{
			{Status: "pending", Count: pendingOrders},
			{Status: "confirmed", Count: totalOrders - 8},
			{Status: "paid", Count: totalOrders - 15},
			{Status: "shipped", Count: totalOrders - 23},
			{Status: "closed", Count: completedOrders},
		}
	case "category-breakdown":
		resp.Data = []dto.CategoryBreakdownData{
			{Category: "Atasan", Revenue: 42500000, Percent: 39.2},
			{Category: "Bawahan", Revenue: 26700000, Percent: 24.6},
			{Category: "Dress", Revenue: 19800000, Percent: 18.3},
			{Category: "Outerwear", Revenue: 19400000, Percent: 17.9},
		}
	case "low-stock":
		resp.Data = []dto.LowStockData{
			{ProductName: "Jaket Bomber Unisex", CurrentStock: 4, MinStock: 10},
			{ProductName: "Dress Batik Sogan", CurrentStock: 3, MinStock: 8},
			{ProductName: "Celana Chino Slim Fit", CurrentStock: 6, MinStock: 10},
		}
	case "revenue-trend":
		resp.Data = []dto.RevenueTrendData{
			{Period: "Week 1", Revenue: 24700000},
			{Period: "Week 2", Revenue: 25900000},
			{Period: "Week 3", Revenue: 27200000},
			{Period: "Week 4", Revenue: 28900000},
		}
	}

	return resp
}

type n8nPayload struct {
	ReportType string                 `json:"report_type"`
	Data       interface{}            `json:"data"`
	Metrics    map[string]interface{} `json:"metrics"`
	Filters    map[string]string      `json:"filters"`
}

type n8nResult struct {
	ChartType      string       `json:"chart_type"`
	Summary        string       `json:"summary"`
	Recommendation string       `json:"recommendation"`
	Anomalies      []n8nAnomaly `json:"anomalies"`
}

type n8nAnomaly struct {
	Metric      string  `json:"metric"`
	Actual      float64 `json:"actual"`
	Expected    float64 `json:"expected"`
	VariancePct float64 `json:"variance_pct"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Deviation   float64 `json:"deviation"`
	Message     string  `json:"message"`
}

func normalizeAnomalies(items []n8nAnomaly) []dto.Anomaly {
	result := make([]dto.Anomaly, 0, len(items))
	for _, item := range items {
		actual := item.Actual
		if actual == 0 && item.Value != 0 {
			actual = item.Value
		}

		variance := item.VariancePct
		if variance == 0 && item.Deviation != 0 {
			variance = item.Deviation
		}

		description := item.Description
		if description == "" && item.Message != "" {
			description = item.Message
		}

		result = append(result, dto.Anomaly{
			Metric:      item.Metric,
			Actual:      actual,
			Expected:    item.Expected,
			VariancePct: variance,
			Description: description,
		})
	}
	return result
}

func (s *reportService) callN8N(ctx context.Context, req dto.ReportRequest, report *dto.ReportResponse) (*n8nResult, error) {
	baseURL := strings.TrimSpace(config.App.N8NBaseURL)
	path := strings.TrimSpace(config.App.N8NDashboardWebhookPath)
	if baseURL == "" || path == "" {
		return nil, errors.New("konfigurasi n8n belum lengkap")
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	payload := n8nPayload{
		ReportType: report.Type,
		Data:       report.Data,
		Metrics:    report.Metrics,
		Filters: map[string]string{
			"from":     req.From,
			"to":       req.To,
			"sales_id": req.SalesID,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal serialisasi payload n8n: %w", err)
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()

	url := strings.TrimRight(baseURL, "/") + path
	httpReq, err := http.NewRequestWithContext(ctxTimeout, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat request n8n: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if secret := strings.TrimSpace(config.App.N8NWebhookSecret); secret != "" {
		httpReq.Header.Set("X-Webhook-Secret", secret)
	}

	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("n8n tidak dapat diakses: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("gagal membaca respons n8n: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("n8n merespons status %d", resp.StatusCode)
	}

	parsed, err := parseN8NResult(raw)
	if err != nil {
		return nil, fmt.Errorf("respons n8n tidak valid: %w", err)
	}
	return parsed, nil
}

func parseN8NResult(raw []byte) (*n8nResult, error) {
	var direct n8nResult
	if err := json.Unmarshal(raw, &direct); err == nil {
		if direct.Summary != "" || direct.ChartType != "" || direct.Recommendation != "" || len(direct.Anomalies) > 0 {
			return &direct, nil
		}
	}

	var wrapped struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &wrapped); err == nil && len(wrapped.Data) > 0 {
		var nested n8nResult
		if err := json.Unmarshal(wrapped.Data, &nested); err == nil {
			return &nested, nil
		}
	}

	return nil, errors.New("field chart_type/summary tidak ditemukan")
}
