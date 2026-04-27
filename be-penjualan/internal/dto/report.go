package dto

// ReportRequest represents query parameters for report generation.
type ReportRequest struct {
	Type    string `query:"type" example:"daily-sales"`
	From    string `query:"from" example:"2026-04-01"`
	To      string `query:"to" example:"2026-04-30"`
	SalesID string `query:"sales_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Mode    string `query:"mode" example:"raw"`
}

// ReportResponse represents dashboard report output consumed by frontend and n8n.
type ReportResponse struct {
	Type           string                 `json:"type" example:"daily-sales"`
	Title          string                 `json:"title" example:"Penjualan Harian"`
	Description    string                 `json:"description" example:"Ringkasan penjualan harian dalam periode tertentu"`
	ChartType      string                 `json:"chart_type" example:"line"`
	Data           interface{}            `json:"data"`
	Metrics        map[string]interface{} `json:"metrics,omitempty"`
	Summary        string                 `json:"summary" example:"Penjualan dalam 7 hari terakhir cenderung naik."`
	Anomalies      []Anomaly              `json:"anomalies,omitempty"`
	Recommendation string                 `json:"recommendation,omitempty" example:"Pertahankan stok produk terlaris untuk 3 hari ke depan."`
	AISource       string                 `json:"ai_source" example:"n8n"`
	FallbackReason string                 `json:"fallback_reason,omitempty" example:"n8n timeout, menggunakan summary default backend"`
}

// Anomaly represents anomaly output after normalization.
type Anomaly struct {
	Metric      string  `json:"metric" example:"daily_revenue"`
	Actual      float64 `json:"actual" example:"2000000"`
	Expected    float64 `json:"expected" example:"5000000"`
	VariancePct float64 `json:"variance_pct" example:"-60"`
	Description string  `json:"description" example:"Pendapatan harian turun signifikan dari baseline."`
}

// DailySalesData represents data for daily sales report
type DailySalesData struct {
	Date  string  `json:"date" example:"2026-04-24"`
	Value float64 `json:"value" example:"5000000"`
}

// MonthlySalesData represents data for monthly sales report
type MonthlySalesData struct {
	Month string  `json:"month" example:"April 2026"`
	Value float64 `json:"value" example:"150000000"`
}

// TopProductData represents data for top products report
type TopProductData struct {
	ProductName string  `json:"product_name" example:"Kaos Polos Cotton Combed 30s"`
	Quantity    int     `json:"quantity" example:"150"`
	Revenue     float64 `json:"revenue" example:"13350000"`
}

// SalesPersonData represents data for sales by person report
type SalesPersonData struct {
	SalesName string  `json:"sales_name" example:"Jane Doe"`
	Revenue   float64 `json:"revenue" example:"75000000"`
	Orders    int     `json:"orders" example:"50"`
}

// OrderFunnelData represents data for order funnel report
type OrderFunnelData struct {
	Status string `json:"status" example:"pending"`
	Count  int    `json:"count" example:"100"`
}

// CategoryBreakdownData represents data for category breakdown report
type CategoryBreakdownData struct {
	Category string  `json:"category" example:"atasan"`
	Revenue  float64 `json:"revenue" example:"100000000"`
	Percent  float64 `json:"percent" example:"66.67"`
}

// LowStockData represents data for low stock report
type LowStockData struct {
	ProductName  string `json:"product_name" example:"Kaos Polos Cotton Combed 30s"`
	CurrentStock int    `json:"current_stock" example:"5"`
	MinStock     int    `json:"min_stock" example:"10"`
}

// RevenueTrendData represents data for revenue trend report
type RevenueTrendData struct {
	Period  string  `json:"period" example:"Week 1"`
	Revenue float64 `json:"revenue" example:"25000000"`
}
