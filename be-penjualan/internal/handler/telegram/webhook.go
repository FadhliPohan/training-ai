package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/dto"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/response"
)

// Handler handles inbound Telegram webhook requests.
type Handler struct {
	userRepo repository.UserRepository
}

// New creates a new Telegram webhook handler.
func New() *Handler {
	return &Handler{
		userRepo: repository.NewUserRepository(),
	}
}

// Webhook godoc
//
//	@Summary		Inbound Telegram webhook
//	@Description	Receives Telegram Update payloads, resolves the sender's role from app.users
//	@Description	via telegram_user_id, then forwards the enriched payload to n8n for AI processing.
//	@Description	Telegram requires a 200 OK response quickly — all heavy processing is async via n8n.
//	@Tags			Telegram
//	@Accept			json
//	@Produce		json
//	@Param			X-Telegram-Bot-Api-Secret-Token	header	string					false	"Optional Telegram webhook secret (set in setWebhook)"
//	@Param			body							body	dto.TelegramWebhookRequest	true	"Telegram Update object"
//	@Success		200								{object}	response.Standard
//	@Failure		400								{object}	response.Standard
//	@Router			/api/v1/telegram/webhook [post]
func (h *Handler) Webhook(c *fiber.Ctx) error {
	// 1. Parse Telegram Update payload
	var req dto.TelegramWebhookRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Payload Telegram tidak valid.", nil)
	}

	// 2. Ignore non-message updates (callback_query, edited_message, etc.)
	if req.Message == nil {
		// Return 200 immediately — Telegram expects a quick OK for all update types
		return response.OK(c, "ok", nil)
	}

	msg := req.Message

	// 3. Ignore messages without text (stickers, files, etc.)
	if msg.Text == "" {
		return response.OK(c, "ok", nil)
	}

	// 4. Ignore bot messages
	if msg.From == nil || msg.From.IsBot {
		return response.OK(c, "ok", nil)
	}

	telegramUserID := msg.From.ID

	// 5. Resolve user from telegram_user_id in app.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.userRepo.FindByTelegramUserID(ctx, telegramUserID)
	if err != nil {
		// User not registered — do NOT return error to Telegram (would cause retries).
		// Fire-and-forget: ask n8n to send "not registered" reply.
		go h.forwardUnregistered(telegramUserID, msg.Chat.ID, msg.MessageID)
		return response.OK(c, "ok", nil)
	}

	// 6. Build enriched payload for n8n
	payload := dto.TelegramWebhookResponse{
		UserID:         user.ID.String(),
		UserNama:       user.Nama,
		Role:           user.Role,
		TelegramUserID: telegramUserID,
		ChatID:         msg.Chat.ID,
		MessageText:    msg.Text,
		MessageID:      msg.MessageID,
	}

	// 7. Forward to n8n asynchronously — respond to Telegram immediately
	go h.forwardToN8N(payload)

	return response.OK(c, "ok", nil)
}

// forwardToN8N posts the enriched payload to the n8n Telegram Q&A workflow webhook.
func (h *Handler) forwardToN8N(payload dto.TelegramWebhookResponse) {
	n8nURL := fmt.Sprintf("%s%s",
		config.App.N8NBaseURL,
		config.App.N8NTelegramWebhookPath,
	)

	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[telegram/webhook] failed to marshal payload: %v", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, n8nURL, bytes.NewReader(body))
	if err != nil {
		log.Printf("[telegram/webhook] failed to build n8n request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if secret := config.App.N8NWebhookSecret; secret != "" {
		req.Header.Set("X-N8N-Webhook-Secret", secret)
	}

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[telegram/webhook] n8n request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Printf("[telegram/webhook] n8n responded with status %d", resp.StatusCode)
	}
}

// forwardUnregistered notifies n8n to reply to an unregistered Telegram user.
func (h *Handler) forwardUnregistered(telegramUserID, chatID, messageID int64) {
	n8nURL := fmt.Sprintf("%s%s",
		config.App.N8NBaseURL,
		config.App.N8NTelegramWebhookPath,
	)

	// Send minimal payload — n8n will detect missing user_id and reply accordingly
	payload := dto.TelegramWebhookResponse{
		TelegramUserID: telegramUserID,
		ChatID:         chatID,
		MessageID:      messageID,
		// UserID empty → n8n treats as unregistered
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, n8nURL, bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if secret := config.App.N8NWebhookSecret; secret != "" {
		req.Header.Set("X-N8N-Webhook-Secret", secret)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, _ := client.Do(req)
	if resp != nil {
		resp.Body.Close()
	}
}
