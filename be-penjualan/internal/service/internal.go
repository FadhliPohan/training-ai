package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"insightflow/be-penjualan/internal/dto"
)

// InternalService provides data endpoints for n8n workflows.
// All methods return flat, AI-friendly structs — no JWT, no pagination.
type InternalService interface {
	Summary(ctx context.Context) (*dto.InternalSummaryResponse, error)
	Anomaly(ctx context.Context, thresholdPct float64) (*dto.InternalAnomalyResponse, error)
}

type internalService struct{}

// NewInternalService creates a new InternalService.
func NewInternalService() InternalService {
	return &internalService{}
}

// Summary returns aggregated sales data for today.
// Currently uses dummy data; replace SQL queries here when live data is ready.
func (s *internalService) Summary(ctx context.Context) (*dto.InternalSummaryResponse, error) {
	now := time.Now()

	// --- TODAY dummy metrics ---
	totalRevenue := 4_750_000.0
	totalOrders := 23
	completedOrders := 17
	pendingOrders := 3
	confirmedOrders := 3
	avgOrderValue := totalRevenue / float64(totalOrders)

	// --- YESTERDAY dummy (for comparison) ---
	yesterdayRevenue := 4_200_000.0
	yesterdayOrders := 20

	revenueVsPct := pctChange(yesterdayRevenue, totalRevenue)
	ordersVsPct := pctChange(float64(yesterdayOrders), float64(totalOrders))

	// --- Top performers ---
	topProduct := "Kemeja Batik Lengan Panjang"
	topCategory := "Atasan"
	topSales := "Citra Sales"

	// --- Low stock ---
	lowStockProducts := []string{"Jaket Bomber Unisex", "Dress Batik Sogan"}
	lowStockCount := len(lowStockProducts)

	// --- Anomaly hint ---
	hasAnomaly := math.Abs(revenueVsPct) > 15
	anomalyMsg := ""
	if hasAnomaly {
		dir := "naik"
		if revenueVsPct < 0 {
			dir = "turun"
		}
		anomalyMsg = fmt.Sprintf("Omzet hari ini %s %.1f%% dibanding kemarin.", dir, math.Abs(revenueVsPct))
	}

	dayNames := map[time.Weekday]string{
		time.Sunday:    "Minggu",
		time.Monday:    "Senin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jumat",
		time.Saturday:  "Sabtu",
	}

	return &dto.InternalSummaryResponse{
		Date:               now.Format("2006-01-02"),
		DayOfWeek:          dayNames[now.Weekday()],
		TotalRevenue:       totalRevenue,
		TotalOrders:        totalOrders,
		CompletedOrders:    completedOrders,
		PendingOrders:      pendingOrders,
		ConfirmedOrders:    confirmedOrders,
		AvgOrderValue:      avgOrderValue,
		TopProduct:         topProduct,
		TopCategory:        topCategory,
		TopSales:           topSales,
		RevenueVsYesterday: revenueVsPct,
		OrdersVsYesterday:  ordersVsPct,
		LowStockCount:      lowStockCount,
		LowStockProducts:   lowStockProducts,
		HasAnomaly:         hasAnomaly,
		AnomalyMsg:         anomalyMsg,
	}, nil
}

// Anomaly checks today's metrics against a 7-day baseline and returns detected anomalies.
// Currently uses dummy data; replace with real SQL aggregation when live data is ready.
func (s *internalService) Anomaly(ctx context.Context, thresholdPct float64) (*dto.InternalAnomalyResponse, error) {
	if thresholdPct <= 0 {
		thresholdPct = 10.0 // default 10%
	}

	now := time.Now()

	// --- Dummy: today vs 7-day average ---
	checks := []struct {
		metric      string
		label       string
		today       float64
		sevenDayAvg float64
	}{
		{"daily_revenue", "Omzet Harian", 4_750_000, 5_300_000},
		{"order_count", "Jumlah Order", 23, 25},
		{"avg_order_value", "Rata-rata Nilai Order", 206_521, 212_000},
	}

	var anomalies []dto.AnomalyDetected
	for _, c := range checks {
		variance := pctChange(c.sevenDayAvg, c.today)
		if math.Abs(variance) < thresholdPct {
			continue
		}

		direction := "naik"
		if variance < 0 {
			direction = "turun"
		}

		severity := "warning"
		if math.Abs(variance) >= thresholdPct*2 {
			severity = "critical"
		}

		anomalies = append(anomalies, dto.AnomalyDetected{
			Metric:      c.metric,
			MetricLabel: c.label,
			Actual:      c.today,
			Expected:    c.sevenDayAvg,
			VariancePct: math.Round(variance*100) / 100,
			Direction:   direction,
			Severity:    severity,
		})
	}

	return &dto.InternalAnomalyResponse{
		CheckedAt:    now,
		ThresholdPct: thresholdPct,
		HasAnomaly:   len(anomalies) > 0,
		Anomalies:    anomalies,
	}, nil
}

// pctChange calculates percentage change from baseline to current.
// Returns positive for increase, negative for decrease.
func pctChange(baseline, current float64) float64 {
	if baseline == 0 {
		return 0
	}
	raw := ((current - baseline) / baseline) * 100
	return math.Round(raw*100) / 100
}
