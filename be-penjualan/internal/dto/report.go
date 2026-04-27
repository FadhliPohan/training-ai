package dto

import "time"

// ==================== REPORT DTOs ====================

// ReportRequest represents the request payload for generating a report
type ReportRequest struct {
	Type     string    `query:"type" validate:"required,oneof=daily-sales monthly-sales top-products sales-by-person order-funnel category-breakdown low-stock revenue-trend" example:"daily-sales"`
	FromDate time.Time `query:"from" validate:"required" example:"2026-04-01T00:00:00Z"`
	ToDate   time.Time `query:"to" validate:"required" example:"2026-04-30T23:59:59Z"`
	SalesID  *string   `query:"sales_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// ReportResponse represents the response payload for a report
type ReportResponse struct {
	Type        string      `json:"type" example:"daily-sales"`
	Title       string      `json:"title" example:"Penjualan Harian"`
	Description string      `json:"description" example:"Grafik penjualan harian dalam periode tertentu"`
	ChartType   string      `json:"chart_type" example:"line"`
	Data        interface{} `json:"data"`
	Summary     string      `json:"summary" example:"Total penjualan meningkat 15% dibanding bulan sebelumnya."`
	Anomalies   []Anomaly   `json:"anomalies,omitempty"`
	Recommendation string   `json:"recommendation,omitempty" example:"Pertimbangkan promosi untuk produk dengan penjualan rendah.""`
}

// Anomaly represents an anomaly detected in the report data
type Anomaly struct {
	Metric     string  `json:"metric" example:"daily_revenue"`
	Value      float64 `json:"value" example:"5000000"`
	Expected   float64 `json:"expected" example:"6000000"`
	Deviation  float64 `json:"deviation" example:"-16.67"`
	Threshold  float64 `json:"threshold" example:"10.00"`
	Message    string  `json:"message" example:"Pendapatan harian turun 16.67% dari ekspektasi."`
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
	ProductName string `json:"product_name" example:"Kaos Polos Cotton Combed 30s"`
	CurrentStock int   `json:"current_stock" example:"5"`
	MinStock    int    `json:"min_stock" example:"10"`
}

// RevenueTrendData represents data for revenue trend report
type RevenueTrendData struct {
	Period  string  `json:"period" example:"Week 1"`
	Revenue float64 `json:"revenue" example:"25000000"`
}