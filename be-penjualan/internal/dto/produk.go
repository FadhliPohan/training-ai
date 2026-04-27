package dto

// ==================== PRODUK DTOs ====================

// ProdukRequest represents the request payload for creating/updating a product
type ProdukRequest struct {
	KodeProduk      string  `json:"kode_produk" validate:"required" example:"KAO-001-M-HTM"`
	Nama            string  `json:"nama" validate:"required,min=1,max=200" example:"Kaos Polos Cotton Combed 30s"`
	KategoriPakaian string  `json:"kategori_pakaian" validate:"required,oneof=atasan bawahan dress outerwear aksesoris" example:"atasan"`
	Ukuran          string  `json:"ukuran" validate:"required" example:"M"`
	Warna           string  `json:"warna" validate:"required" example:"Hitam"`
	Bahan           string  `json:"bahan" validate:"required" example:"Katun Combed 30s"`
	Harga           float64 `json:"harga" validate:"required,min=0" example:"89000"`
	Stok            int     `json:"stok" validate:"required,min=0" example:"50"`
}

// ProdukResponse represents the response payload for a product
type ProdukResponse struct {
	ID              int     `json:"id" example:"1"`
	KodeProduk      string  `json:"kode_produk" example:"KAO-001-M-HTM"`
	Nama            string  `json:"nama" example:"Kaos Polos Cotton Combed 30s"`
	KategoriPakaian string  `json:"kategori_pakaian" example:"atasan"`
	Ukuran          string  `json:"ukuran" example:"M"`
	Warna           string  `json:"warna" example:"Hitam"`
	Bahan           string  `json:"bahan" example:"Katun Combed 30s"`
	Harga           float64 `json:"harga" example:"89000"`
	Stok            int     `json:"stok" example:"50"`
	Aktif           bool    `json:"aktif" example:"true"`
}

// ProdukListResponse represents the response payload for a list of products
type ProdukListResponse struct {
	Produk []ProdukResponse `json:"produk"`
}

// ProdukFilter represents the filter parameters for listing products
type ProdukFilter struct {
	Aktif           *bool  `query:"aktif" validate:"omitempty"`
	KategoriPakaian string `query:"kategori_pakaian" validate:"omitempty,oneof=atasan bawahan dress outerwear aksesoris"`
	Ukuran          string `query:"ukuran" validate:"omitempty"`
	Warna           string `query:"warna" validate:"omitempty"`
	Bahan           string `query:"bahan" validate:"omitempty"`
	Search          string `query:"search" validate:"omitempty"`
	Page            int    `query:"page" validate:"omitempty,min=1" example:"1"`
	Limit           int    `query:"limit" validate:"omitempty,min=1,max=100" example:"10"`
}

// ProdukIDResponse represents the response payload containing only a product ID
type ProdukIDResponse struct {
	ID int `json:"id" example:"1"`
}