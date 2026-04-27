package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/response"
	"insightflow/be-penjualan/internal/service"
)

// Handler holds user service dependency.
type Handler struct {
	svc service.UserService
}

// New creates a new users Handler.
func New() *Handler {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	return &Handler{svc: svc}
}

// List handles GET /api/v1/users
//
//	@Summary		List semua user
//	@Tags			Users
//	@Security		JWT
//	@Produce		json
//	@Success		200	{object}	response.Standard
//	@Router			/users [get]
func (h *Handler) List(c *fiber.Ctx) error {
	list, err := h.svc.List(c.Context())
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.OK(c, "Berhasil mendapatkan daftar user", list)
}

// GetByID handles GET /api/v1/users/:id
//
//	@Summary		Detail user by ID
//	@Tags			Users
//	@Security		JWT
//	@Produce		json
//	@Param			id	path		string	true	"User UUID"
//	@Success		200	{object}	response.Standard
//	@Failure		404	{object}	response.Standard
//	@Router			/users/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID user tidak valid", nil)
	}
	u, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, err.Error())
	}
	return response.OK(c, "Berhasil mendapatkan data user", u)
}

// CreateRequest is the payload for creating a user.
type CreateRequest struct {
	Nama           string `json:"nama"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	TelegramUserID *int64 `json:"telegram_user_id,omitempty"`
}

// Create handles POST /api/v1/users
//
//	@Summary		Buat user baru (admin only)
//	@Tags			Users
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateRequest	true	"Data user"
//	@Success		201		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		409		{object}	response.Standard
//	@Router			/users [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	u := &domain.User{
		Nama:           req.Nama,
		Email:          req.Email,
		Role:           req.Role,
		TelegramUserID: req.TelegramUserID,
	}

	if err := h.svc.Create(c.Context(), u, req.Password); err != nil {
		if err.Error() == "email sudah terdaftar" {
			return response.Conflict(c, err.Error())
		}
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.Created(c, "User berhasil dibuat", fiber.Map{
		"id":    u.ID,
		"nama":  u.Nama,
		"email": u.Email,
		"role":  u.Role,
	})
}

// UpdateRequest is the payload for updating a user.
type UpdateRequest struct {
	Nama           string `json:"nama"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	TelegramUserID *int64 `json:"telegram_user_id,omitempty"`
	Aktif          bool   `json:"aktif"`
}

// Update handles PUT /api/v1/users/:id
//
//	@Summary		Update user (admin only)
//	@Tags			Users
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"User UUID"
//	@Param			request	body		UpdateRequest	true	"Data user"
//	@Success		200		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		404		{object}	response.Standard
//	@Router			/users/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID user tidak valid", nil)
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	u := &domain.User{
		ID:             id,
		Nama:           req.Nama,
		Email:          req.Email,
		Role:           req.Role,
		TelegramUserID: req.TelegramUserID,
		Aktif:          req.Aktif,
	}

	if err := h.svc.Update(c.Context(), u); err != nil {
		if err.Error() == "user tidak ditemukan" {
			return response.NotFound(c, err.Error())
		}
		if err.Error() == "email sudah digunakan user lain" {
			return response.Conflict(c, err.Error())
		}
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.OK(c, "User berhasil diupdate", u)
}

// Deactivate handles PATCH /api/v1/users/:id
//
//	@Summary		Nonaktifkan user (admin only)
//	@Tags			Users
//	@Security		JWT
//	@Produce		json
//	@Param			id	path		string	true	"User UUID"
//	@Success		200	{object}	response.Standard
//	@Failure		404	{object}	response.Standard
//	@Router			/users/{id} [patch]
func (h *Handler) Deactivate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID user tidak valid", nil)
	}

	if err := h.svc.Deactivate(c.Context(), id); err != nil {
		return response.NotFound(c, err.Error())
	}
	return response.OK(c, "User berhasil dinonaktifkan", nil)
}
