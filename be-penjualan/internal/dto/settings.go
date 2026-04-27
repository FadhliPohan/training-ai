package dto

// ==================== SETTINGS DTOs ====================

// TelegramConfigRequest represents the request payload for Telegram configuration
type TelegramConfigRequest struct {
	NamaGrup   string `json:"nama_grup" validate:"required,min=1,max=100" example:"Sales Team"`
	ChatID     int64  `json:"chat_id" validate:"required" example:"-1001234567890"`
	JamSummary string `json:"jam_summary" validate:"required,datetime=15:04" example:"07:00"`
}

// TelegramConfigResponse represents the response payload for Telegram configuration
type TelegramConfigResponse struct {
	ID         string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	NamaGrup   string `json:"nama_grup" example:"Sales Team"`
	ChatID     int64  `json:"chat_id" example:"-1001234567890"`
	Aktif      bool   `json:"aktif" example:"true"`
	JamSummary string `json:"jam_summary" example:"07:00"`
}

// UpdateTelegramConfigRequest represents the request payload for updating Telegram configuration
type UpdateTelegramConfigRequest struct {
	NamaGrup   *string `json:"nama_grup" validate:"omitempty,min=1,max=100"`
	ChatID     *int64  `json:"chat_id" validate:"omitempty"`
	JamSummary *string `json:"jam_summary" validate:"omitempty,datetime=15:04"`
	Aktif      *bool   `json:"aktif" validate:"omitempty"`
}

// AnomalyConfigRequest represents the request payload for anomaly configuration
type AnomalyConfigRequest struct {
	MetricKey    string  `json:"metric_key" validate:"required" example:"daily_revenue"`
	ThresholdPct float64 `json:"threshold_pct" validate:"required,min=0,max=100" example:"10.00"`
	Aktif        bool    `json:"aktif" example:"true"`
}

// AnomalyConfigResponse represents the response payload for anomaly configuration
type AnomalyConfigResponse struct {
	ID           string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	MetricKey    string  `json:"metric_key" example:"daily_revenue"`
	ThresholdPct float64 `json:"threshold_pct" example:"10.00"`
	Aktif        bool    `json:"aktif" example:"true"`
}

// UpdateAnomalyConfigRequest represents the request payload for updating anomaly configuration
type UpdateAnomalyConfigRequest struct {
	ThresholdPct *float64 `json:"threshold_pct" validate:"omitempty,min=0,max=100"`
	Aktif        *bool    `json:"aktif" validate:"omitempty"`
}