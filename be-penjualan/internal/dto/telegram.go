package dto

// TelegramWebhookRequest is the payload sent by Telegram to our webhook endpoint.
// Telegram sends only what's relevant to the update type; we only parse what we need.
// Reference: https://core.telegram.org/bots/api#update
type TelegramWebhookRequest struct {
	UpdateID int64           `json:"update_id"`
	Message  *TelegramMessage `json:"message,omitempty"`
}

// TelegramMessage represents an incoming Telegram message.
type TelegramMessage struct {
	MessageID int64          `json:"message_id"`
	From      *TelegramUser  `json:"from,omitempty"`
	Chat      TelegramChat   `json:"chat"`
	Date      int64          `json:"date"`
	Text      string         `json:"text"`
}

// TelegramUser represents the sender of a Telegram message.
type TelegramUser struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

// TelegramChat represents the chat context of an incoming message.
type TelegramChat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"` // private | group | supergroup | channel
}

// TelegramWebhookResponse is the payload forwarded from our backend to n8n.
// n8n receives this and drives the intent → endpoint → AI → Telegram reply pipeline.
type TelegramWebhookResponse struct {
	// User info resolved from app.users via telegram_user_id
	UserID   string `json:"user_id"`
	UserNama string `json:"user_nama"`
	Role     string `json:"role"`

	// Original Telegram context — n8n uses ChatID to send the reply
	TelegramUserID int64  `json:"telegram_user_id"`
	ChatID         int64  `json:"chat_id"`
	MessageText    string `json:"message_text"`
	MessageID      int64  `json:"message_id"`
}
