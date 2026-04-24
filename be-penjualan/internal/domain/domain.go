package domain

import (
	"time"

	"github.com/google/uuid"
)

// Role constants define valid user roles in the system.
const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleSales   = "sales"
	RoleViewer  = "viewer"
)

// User represents an internal application user (admin, manager, sales, viewer).
type User struct {
	ID             uuid.UUID  `json:"id"`
	Nama           string     `json:"nama"`
	Email          string     `json:"email"`
	Password       string     `json:"-"` // bcrypt hash — never serialised to JSON
	Role           string     `json:"role"`
	TelegramUserID *int64     `json:"telegram_user_id,omitempty"`
	Aktif          bool       `json:"aktif"`
	CreatedAt      time.Time  `json:"created_at"`
}

// Produk represents a clothing product in the catalogue.
type Produk struct {
	ID              int        `json:"id"`
	KodeProduk      string     `json:"kode_produk"`
	Nama            string     `json:"nama"`
	KategoriPakaian string     `json:"kategori_pakaian"`
	Ukuran          string     `json:"ukuran"`
	Warna           string     `json:"warna"`
	Bahan           string     `json:"bahan"`
	Harga           float64    `json:"harga"`
	Stok            int        `json:"stok"`
	Aktif           bool       `json:"aktif"`
}

// Customer represents a buyer who can register and place orders.
type Customer struct {
	ID        int       `json:"id"`
	KodeCust  string    `json:"kode_cust"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email"`
	Telepon   string    `json:"telepon"`
	Alamat    string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderStatus defines valid order lifecycle states.
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusConfirmed OrderStatus = "confirmed"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusClosed    OrderStatus = "closed"
	StatusCancelled OrderStatus = "cancelled"
)

// Order is the header record for a sales transaction.
type Order struct {
	ID         int         `json:"id"`
	NoOrder    string      `json:"no_order"`
	CustomerID int         `json:"customer_id"`
	SalesID    uuid.UUID   `json:"sales_id"`
	Tanggal    time.Time   `json:"tanggal"`
	Status     OrderStatus `json:"status"`
	Total      float64     `json:"total"`
	CreatedAt  time.Time   `json:"created_at"`

	// Populated via JOIN when needed
	Customer *Customer `json:"customer,omitempty"`
	Sales    *User     `json:"sales,omitempty"`
	Details  []OrderDetail `json:"details,omitempty"`
}

// OrderDetail is a single line-item inside an Order.
type OrderDetail struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProdukID  int     `json:"produk_id"`
	Qty       int     `json:"qty"`
	HargaSaat float64 `json:"harga_saat"` // snapshot price at time of order
	Subtotal  float64 `json:"subtotal"`

	Produk *Produk `json:"produk,omitempty"`
}

// PembayaranStatus defines payment verification states.
type PembayaranStatus string

const (
	PembayaranPending  PembayaranStatus = "pending"
	PembayaranVerified PembayaranStatus = "verified"
	PembayaranRejected PembayaranStatus = "rejected"
)

// Pembayaran records a payment transaction for an order.
type Pembayaran struct {
	ID      int              `json:"id"`
	OrderID int              `json:"order_id"`
	Jumlah  float64          `json:"jumlah"`
	Metode  string           `json:"metode"` // transfer | tunai | kartu
	Status  PembayaranStatus `json:"status"`
	Tanggal time.Time        `json:"tanggal"`
}

// PengirimanStatus defines shipment tracking states.
type PengirimanStatus string

const (
	PengirimanProses   PengirimanStatus = "proses"
	PengirimanDikirim  PengirimanStatus = "dikirim"
	PengirimanDiterima PengirimanStatus = "diterima"
)

// Pengiriman records shipment/courier information for an order.
type Pengiriman struct {
	ID      int              `json:"id"`
	OrderID int              `json:"order_id"`
	Kurir   string           `json:"kurir"`
	NoResi  string           `json:"no_resi"`
	Status  PengirimanStatus `json:"status"`
	Tanggal time.Time        `json:"tanggal"`
}

// TelegramConfig holds the Telegram group configuration for a division.
type TelegramConfig struct {
	ID        string    `json:"id"`
	NamaGrup  string    `json:"nama_grup"`
	ChatID    int64     `json:"chat_id"`
	Aktif     bool      `json:"aktif"`
	JamSummary string   `json:"jam_summary"`
	CreatedAt time.Time `json:"created_at"`
}

// AnomalyConfig stores the anomaly detection threshold for a specific metric.
type AnomalyConfig struct {
	ID           string  `json:"id"`
	MetricKey    string  `json:"metric_key"`    // e.g. daily_revenue, order_count
	ThresholdPct float64 `json:"threshold_pct"` // e.g. 10.00 = 10%
	Aktif        bool    `json:"aktif"`
}
