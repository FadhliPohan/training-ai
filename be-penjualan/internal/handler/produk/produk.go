package produk

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/repository"
	"insightflow/be-penjualan/internal/response"
	"insightflow/be-penjualan/internal/service"
)

// Handler holds produk service dependency.
type Handler struct {
	svc service.ProdukService
}

// New creates a new produk Handler.
func New() *Handler {
	repo := repository.NewProdukRepository()
	svc := service.NewProdukService(repo)
	return &Handler{svc: svc}
}

// ListPublic handles GET /api/v1/produk (public catalogue — aktif only)
//
//	@Summary		List produk aktif (public)
//	@Tags			Produk
//	@Produce		json
//	@Success		200	{object}	response.Standard
//	@Router			/produk [get]
func (h *Handler) ListPublic(c *fiber.Ctx) error {
	list, err := h.svc.List(c.Context(), true)
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.OK(c, "Berhasil mendapatkan daftar produk", list)
}

// List handles GET /api/v1/produk (protected — all including non-aktif)
//
//	@Summary		List semua produk (admin/manager/sales)
//	@Tags			Produk
//	@Security		JWT
//	@Produce		json
//	@Param			aktif	query	bool	false	"Filter aktif saja"
//	@Success		200		{object}	response.Standard
//	@Router			/produk [get]
func (h *Handler) List(c *fiber.Ctx) error {
	aktifOnly := c.QueryBool("aktif", false)
	list, err := h.svc.List(c.Context(), aktifOnly)
	if err != nil {
		return response.InternalServerError(c)
	}
	return response.OK(c, "Berhasil mendapatkan daftar produk", list)
}

// GetByID handles GET /api/v1/produk/:id
//
//	@Summary		Detail produk by ID
//	@Tags			Produk
//	@Security		JWT
//	@Produce		json
//	@Param			id	path		int	true	"Produk ID"
//	@Success		200	{object}	response.Standard
//	@Failure		404	{object}	response.Standard
//	@Router			/produk/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID produk tidak valid", nil)
	}

	p, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, err.Error())
	}
	return response.OK(c, "Berhasil mendapatkan data produk", p)
}

// CreateRequest is the payload for creating a new product.
type CreateRequest struct {
	KodeProduk      string  `json:"kode_produk"`
	Nama            string  `json:"nama"`
	KategoriPakaian string  `json:"kategori_pakaian"`
	Ukuran          string  `json:"ukuran"`
	Warna           string  `json:"warna"`
	Bahan           string  `json:"bahan"`
	Harga           float64 `json:"harga"`
	Stok            int     `json:"stok"`
}

// Create handles POST /api/v1/produk
//
//	@Summary		Buat produk baru
//	@Tags			Produk
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateRequest	true	"Data produk"
//	@Success		201		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		409		{object}	response.Standard
//	@Router			/produk [post]
func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	p := &domain.Produk{
		KodeProduk:      req.KodeProduk,
		Nama:            req.Nama,
		KategoriPakaian: req.KategoriPakaian,
		Ukuran:          req.Ukuran,
		Warna:           req.Warna,
		Bahan:           req.Bahan,
		Harga:           req.Harga,
		Stok:            req.Stok,
	}

	if err := h.svc.Create(c.Context(), p); err != nil {
		if isConflictErr(err.Error()) {
			return response.Conflict(c, err.Error())
		}
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.Created(c, "Produk berhasil dibuat", p)
}

// UpdateRequest is the payload for updating an existing product.
type UpdateRequest struct {
	KodeProduk      string  `json:"kode_produk"`
	Nama            string  `json:"nama"`
	KategoriPakaian string  `json:"kategori_pakaian"`
	Ukuran          string  `json:"ukuran"`
	Warna           string  `json:"warna"`
	Bahan           string  `json:"bahan"`
	Harga           float64 `json:"harga"`
	Stok            int     `json:"stok"`
	Aktif           bool    `json:"aktif"`
}

// Update handles PUT /api/v1/produk/:id
//
//	@Summary		Update produk
//	@Tags			Produk
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Produk ID"
//	@Param			request	body		UpdateRequest	true	"Data produk"
//	@Success		200		{object}	response.Standard
//	@Failure		400		{object}	response.Standard
//	@Failure		404		{object}	response.Standard
//	@Router			/produk/{id} [put]
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID produk tidak valid", nil)
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Format permintaan tidak valid", nil)
	}

	p := &domain.Produk{
		ID:              id,
		KodeProduk:      req.KodeProduk,
		Nama:            req.Nama,
		KategoriPakaian: req.KategoriPakaian,
		Ukuran:          req.Ukuran,
		Warna:           req.Warna,
		Bahan:           req.Bahan,
		Harga:           req.Harga,
		Stok:            req.Stok,
		Aktif:           req.Aktif,
	}

	if err := h.svc.Update(c.Context(), p); err != nil {
		if err.Error() == "produk tidak ditemukan" {
			return response.NotFound(c, err.Error())
		}
		if isConflictErr(err.Error()) {
			return response.Conflict(c, err.Error())
		}
		return response.BadRequest(c, err.Error(), nil)
	}
	return response.OK(c, "Produk berhasil diupdate", p)
}

// Deactivate handles PATCH /api/v1/produk/:id
//
//	@Summary		Nonaktifkan produk
//	@Tags			Produk
//	@Security		JWT
//	@Produce		json
//	@Param			id	path		int	true	"Produk ID"
//	@Success		200	{object}	response.Standard
//	@Failure		404	{object}	response.Standard
//	@Router			/produk/{id} [patch]
func (h *Handler) Deactivate(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.BadRequest(c, "ID produk tidak valid", nil)
	}

	if err := h.svc.Deactivate(c.Context(), id); err != nil {
		return response.NotFound(c, err.Error())
	}
	return response.OK(c, "Produk berhasil dinonaktifkan", nil)
}

func isConflictErr(msg string) bool {
	return len(msg) > 7 && msg[:7] == "kode_pr"
}
