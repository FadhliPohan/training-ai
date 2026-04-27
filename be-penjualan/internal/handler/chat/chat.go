package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/middleware"
	"insightflow/be-penjualan/internal/response"
)

// systemPrompt is the AI assistant persona injected at the start of every conversation.
const systemPrompt = `Kamu adalah InsightFlow AI, asisten analitik penjualan cerdas untuk dashboard InsightFlow PT PUSRI.
Kamu membantu tim sales, manager, dan viewer dengan:

1. **Analisis Penjualan**: Tren omzet harian/mingguan/bulanan, top produk, performa per sales
2. **Monitoring KPI**: Pencapaian target, anomali penjualan, alert stok menipis
3. **Laporan Cepat**: Ringkasan performa, komparasi periode, proyeksi
4. **Rekomendasi**: Saran strategis berdasarkan data, identifikasi peluang & risiko

Panduan menjawab:
- Gunakan Bahasa Indonesia yang profesional namun ramah
- Jawaban singkat, padat, dan actionable (max 3-4 paragraf kecuali diminta detail)
- Sertakan angka dan metrik jika relevan
- Gunakan emoji secukupnya untuk keterbacaan (📈 📉 ⚠️ ✅ dll)
- Jika data spesifik tidak tersedia, berikan panduan umum yang berguna
- Selalu akhiri dengan saran tindak lanjut jika relevan

Konteks sistem: InsightFlow v1.0, PT PUSRI, dashboard self-service analitik penjualan.`

// Handler holds chat handler dependencies.
type Handler struct {
	httpClient *http.Client
}

// New creates a new chat Handler.
func New() *Handler {
	return &Handler{
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// ChatRequest is the request body for the chat endpoint.
type ChatRequest struct {
	Messages []ChatMessage `json:"messages" validate:"required,min=1"`
}

// ChatMessage represents a single message in the conversation.
type ChatMessage struct {
	Role    string `json:"role"`    // "user" | "assistant"
	Content string `json:"content"` // message text
}

// openAIRequest is the payload sent to OpenAI API.
type openAIRequest struct {
	Model       string              `json:"model"`
	Messages    []openAIMessage     `json:"messages"`
	MaxTokens   int                 `json:"max_tokens"`
	Temperature float64             `json:"temperature"`
	Stream      bool                `json:"stream"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// ChatResponse is the response body for the chat endpoint.
type ChatResponse struct {
	Reply string `json:"reply"`
}

// Chat handles POST /api/v1/chat.
//
//	@Summary		AI Chat for sales & monitoring
//	@Description	Send a conversation to OpenAI GPT-4o with InsightFlow context and get a reply
//	@Tags			Chat
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			body	body		ChatRequest		true	"Conversation messages"
//	@Success		200		{object}	response.Standard{data=ChatResponse}
//	@Failure		400		{object}	response.Standard
//	@Failure		401		{object}	response.Standard
//	@Failure		500		{object}	response.Standard
//	@Router			/chat [post]
func (h *Handler) Chat(c *fiber.Ctx) error {
	// Parse request body
	var req ChatRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format body tidak valid", nil)
	}
	if len(req.Messages) == 0 {
		return response.BadRequest(c, "Messages tidak boleh kosong", nil)
	}

	// Enrich with user identity from JWT claims (if available)
	userInfo := ""
	if claims, ok := middleware.GetClaims(c); ok {
		userInfo = fmt.Sprintf("\n[Pengguna: ID=%s, Role=%s]", claims.UserID, claims.Role)
	}

	// Build OpenAI messages: system prompt + conversation history
	msgs := make([]openAIMessage, 0, len(req.Messages)+1)
	msgs = append(msgs, openAIMessage{
		Role:    "system",
		Content: systemPrompt + userInfo,
	})
	for _, m := range req.Messages {
		role := strings.ToLower(m.Role)
		if role != "user" && role != "assistant" {
			role = "user"
		}
		msgs = append(msgs, openAIMessage{
			Role:    role,
			Content: m.Content,
		})
	}

	// Call OpenAI API
	reply, err := h.callOpenAI(msgs)
	if err != nil {
		c.Context().Logger().Printf("[chat] openai error: %v", err)
		return response.InternalServerError(c)
	}

	return response.OK(c, "Berhasil mendapatkan balasan AI", ChatResponse{Reply: reply})
}

// callOpenAI sends messages to OpenAI API and returns the assistant reply.
func (h *Handler) callOpenAI(msgs []openAIMessage) (string, error) {
	payload := openAIRequest{
		Model:       "gpt-4o",
		Messages:    msgs,
		MaxTokens:   1024,
		Temperature: 0.7,
		Stream:      false,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, config.App.OpenAIURL, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.App.OpenAIAPIKey)

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	var oaiResp openAIResponse
	if err := json.Unmarshal(respBody, &oaiResp); err != nil {
		return "", fmt.Errorf("unmarshal response: %w", err)
	}

	if oaiResp.Error != nil {
		return "", fmt.Errorf("openai api error: %s", oaiResp.Error.Message)
	}

	if len(oaiResp.Choices) == 0 {
		return "", fmt.Errorf("openai returned no choices")
	}

	return oaiResp.Choices[0].Message.Content, nil
}
