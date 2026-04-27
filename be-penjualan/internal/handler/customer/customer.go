package customer

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/response"
	"insightflow/be-penjualan/internal/service"
)

// Handler holds customer service dependency.
type Handler struct {
	svc service.CustomerService
}

// New creates a new customer Handler.
func New() *Handler {
	repo := repository.NewCustomerRepository()
	svc := service.NewCustomerService(repo)
	return &Handler{svc: svc}
}

// List handles GET /api/v1/customer
//
//	@Summary		List semua customer
//	@Tags			Customer
//	@Security		JWT
//	@Produce		json
//	@Success		200	{object}	response.Standard
//	@Router			/customer [get]
func (h *Handler) List(c *fiber.Ctx) error {
	list, err := h.svc.List(c.Context())
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.OK(c, "Berhasil mendapatkan daftar customer", list)
}

// GetByID handles GET /api/v1/customer/:id
//
//	@Summary		Detail customer by ID
//	@Tags			Customer
//	@Security		JWT
//	@Produce		json
//	@Param			id	path		int	true	"Customer ID"
//	@Success		200	{object}	response.Standard
//	@Failure		404	{object}	response.Standard
//	@Router			/customer/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID customer tidak valid", nil)
	}
	cust, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, err.Error())
	}
	return response.OK(c, "Berhasil mendapatkan data customer", cust)
}

// CreateRequest is the payload for creating a customer.
type CreateRequest struct {
	KodeCust string `json:"kode_cust"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Telepon  string `json:"telepon"`
	Alamat   string `json:"alamat"`
}

// Create handles POST /api/v1/customer
//
//	@Summary		Tambah customer baru
//	@Tags			Customer
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateRequest	true	"Data customer"
//	@Success		201		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		409		{object}	response.Standard
//	@Router			/customer [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	cust := &domain.Customer{
		KodeCust: req.KodeCust,
		Nama:     req.Nama,
		Email:    req.Email,
		Telepon:  req.Telepon,
		Alamat:   req.Alamat,
	}

	if err := h.svc.Create(c.Context(), cust); err != nil {
		if err.Error() == "email sudah terdaftar" {
			return response.Conflict(c, err.Error())
		}
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.Created(c, "Customer berhasil ditambahkan", cust)
}

// UpdateRequest is the payload for updating a customer.
type UpdateRequest struct {
	KodeCust string `json:"kode_cust"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Telepon  string `json:"telepon"`
	Alamat   string `json:"alamat"`
	Aktif    bool   `json:"aktif"`
}

// Update handles PUT /api/v1/customer/:id
//
//	@Summary		Update customer
//	@Tags			Customer
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Customer ID"
//	@Param			request	body		UpdateRequest	true	"Data customer"
//	@Success		200		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		404		{object}	response.Standard
//	@Router			/customer/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID customer tidak valid", nil)
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	cust := &domain.Customer{
		ID:       id,
		KodeCust: req.KodeCust,
		Nama:     req.Nama,
		Email:    req.Email,
		Telepon:  req.Telepon,
		Alamat:   req.Alamat,
		Aktif:    req.Aktif,
	}

	if err := h.svc.Update(c.Context(), cust); err != nil {
		if err.Error() == "customer tidak ditemukan" {
			return response.NotFound(c, err.Error())
		}
		if err.Error() == "email sudah digunakan customer lain" {
			return response.Conflict(c, err.Error())
		}
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.OK(c, "Customer berhasil diupdate", cust)
}
