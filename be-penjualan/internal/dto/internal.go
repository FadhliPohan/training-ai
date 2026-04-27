package dto

import "time"

// InternalSummaryResponse is the payload returned by GET /api/internal/reports/summary.
// Designed to be directly consumable by n8n without further parsing.
// n8n will forward this to LLM to generate the Telegram daily summary message.
type InternalSummaryResponse struct {
	// Period info
	Date      string `json:"date"`       // "2026-04-27"
	DayOfWeek string `json:"day_of_week"` // "Senin"

	// Core sales metrics
	TotalRevenue    float64 `json:"total_revenue"`    // total omzet hari ini
	TotalOrders     int     `json:"total_orders"`     // total semua order
	CompletedOrders int     `json:"completed_orders"` // status: closed
	PendingOrders   int     `json:"pending_orders"`   // status: pending
	ConfirmedOrders int     `json:"confirmed_orders"` // status: confirmed
	AvgOrderValue   float64 `json:"avg_order_value"`  // rata-rata nilai order

	// Top performers
	TopProduct  string `json:"top_product"`   // nama produk terlaris
	TopCategory string `json:"top_category"`  // kategori terlaris
	TopSales    string `json:"top_sales"`     // nama sales terbaik hari ini

	// Comparison vs yesterday (for AI to highlight trend)
	RevenueVsYesterday float64 `json:"revenue_vs_yesterday_pct"` // +12.5 = naik 12.5%
	OrdersVsYesterday  float64 `json:"orders_vs_yesterday_pct"`

	// Low stock alert (summary)
	LowStockCount    int      `json:"low_stock_count"`    // jumlah produk stok rendah
	LowStockProducts []string `json:"low_stock_products"` // nama produk stok rendah

	// Anomalies (if any — pre-detected by backend)
	HasAnomaly bool   `json:"has_anomaly"`
	AnomalyMsg string `json:"anomaly_msg,omitempty"` // human-readable jika ada anomali
}

// InternalAnomalyResponse is the payload returned by GET /api/internal/reports/anomaly.
// n8n will check this every 15 minutes to decide whether to send an alert.
type InternalAnomalyResponse struct {
	CheckedAt    time.Time         `json:"checked_at"`
	ThresholdPct float64           `json:"threshold_pct"` // threshold yang digunakan
	HasAnomaly   bool              `json:"has_anomaly"`   // true = perlu alert
	Anomalies    []AnomalyDetected `json:"anomalies"`     // kosong jika tidak ada
}

// AnomalyDetected describes a single detected anomaly.
type AnomalyDetected struct {
	Metric      string  `json:"metric"`       // "daily_revenue", "order_count", dll
	MetricLabel string  `json:"metric_label"` // "Omzet Harian"
	Actual      float64 `json:"actual"`
	Expected    float64 `json:"expected"`     // rata-rata 7 hari terakhir
	VariancePct float64 `json:"variance_pct"` // negatif = turun, positif = naik
	Direction   string  `json:"direction"`    // "turun" | "naik"
	Severity    string  `json:"severity"`     // "warning" | "critical"
}

// InternalUserByTelegramResponse is returned by GET /api/internal/users/by-telegram.
// n8n uses this to resolve a Telegram user's role and inject sales_id into queries.
type InternalUserByTelegramResponse struct {
	UserID string `json:"user_id"`
	Nama   string `json:"nama"`
	Role   string `json:"role"` // "sales" | "manager" | "admin"
}
