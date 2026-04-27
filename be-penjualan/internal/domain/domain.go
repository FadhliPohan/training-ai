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
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Nama           string    `json:"nama" gorm:"type:varchar(100);not null"`
	Email          string    `json:"email" gorm:"type:varchar(150);not null;uniqueIndex"`
	Password       string    `json:"-" gorm:"type:text;not null"` // bcrypt hash — never serialised to JSON
	Role           string    `json:"role" gorm:"type:varchar(20);not null"`
	TelegramUserID *int64    `json:"telegram_user_id,omitempty" gorm:"index:idx_users_telegram_user_id"`
	Aktif          bool      `json:"aktif" gorm:"not null;default:true"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;default:now()"`
}

// Produk represents a clothing product in the catalogue.
type Produk struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KodeProduk      string    `json:"kode_produk" gorm:"type:varchar(50);not null;uniqueIndex"`
	Nama            string    `json:"nama" gorm:"type:varchar(150);not null"`
	KategoriPakaian string    `json:"kategori_pakaian" gorm:"type:varchar(100);not null;index:idx_tbl_produk_kategori"`
	Ukuran          string    `json:"ukuran" gorm:"type:varchar(20);not null"`
	Warna           string    `json:"warna" gorm:"type:varchar(50);not null"`
	Bahan           string    `json:"bahan" gorm:"type:varchar(100);not null"`
	Harga           float64   `json:"harga" gorm:"type:numeric(14,2);not null"`
	Stok            int       `json:"stok" gorm:"not null;default:0"`
	Aktif           bool      `json:"aktif" gorm:"not null;default:true;index:idx_tbl_produk_aktif"`
	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"not null;default:now()"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"not null;default:now()"`
}

// Customer represents a buyer who can register and place orders.
type Customer struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	KodeCust  string    `json:"kode_cust" gorm:"type:varchar(50);not null;uniqueIndex"`
	Nama      string    `json:"nama" gorm:"type:varchar(150);not null"`
	Email     string    `json:"email" gorm:"type:varchar(150);not null;uniqueIndex"`
	Password  string    `json:"-" gorm:"type:text"`
	Telepon   string    `json:"telepon" gorm:"type:varchar(30)"`
	Alamat    string    `json:"alamat" gorm:"type:text"`
	Aktif     bool      `json:"aktif" gorm:"not null;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:now()"`
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
	ID          int         `json:"id" gorm:"primaryKey;autoIncrement"`
	NoOrder     string      `json:"no_order" gorm:"type:varchar(50);not null;uniqueIndex"`
	CustomerID  int         `json:"customer_id" gorm:"not null;index:idx_tbl_order_customer_id"`
	SalesID     uuid.UUID   `json:"sales_id" gorm:"type:uuid;not null;index:idx_tbl_order_sales_id"`
	Tanggal     time.Time   `json:"tanggal" gorm:"type:date;not null;default:CURRENT_DATE;index:idx_tbl_order_tanggal"`
	Status      OrderStatus `json:"status" gorm:"type:varchar(20);not null;default:'pending';index:idx_tbl_order_status"`
	Total       float64     `json:"total" gorm:"type:numeric(14,2);not null;default:0"`
	AlasanBatal *string     `json:"alasan_batal,omitempty" gorm:"type:text"`
	CreatedAt   time.Time   `json:"created_at" gorm:"not null;default:now()"`

	// Populated via JOIN when needed
	Customer *Customer     `json:"customer,omitempty" gorm:"-"`
	Sales    *User         `json:"sales,omitempty" gorm:"-"`
	Details  []OrderDetail `json:"details,omitempty" gorm:"-"`
}

// OrderDetail is a single line-item inside an Order.
type OrderDetail struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int     `json:"order_id" gorm:"not null;index:idx_tbl_order_detail_order_id"`
	ProdukID  int     `json:"produk_id" gorm:"not null;index:idx_tbl_order_detail_produk_id"`
	Qty       int     `json:"qty" gorm:"not null"`
	HargaSaat float64 `json:"harga_saat" gorm:"type:numeric(14,2);not null"` // snapshot price at time of order
	Subtotal  float64 `json:"subtotal" gorm:"type:numeric(14,2);not null"`

	Produk *Produk `json:"produk,omitempty" gorm:"-"`
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
	ID      int              `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID int              `json:"order_id" gorm:"not null;index:idx_tbl_pembayaran_order_id"`
	Jumlah  float64          `json:"jumlah" gorm:"type:numeric(14,2);not null"`
	Metode  string           `json:"metode" gorm:"type:varchar(20);not null"` // transfer | tunai | kartu
	Status  PembayaranStatus `json:"status" gorm:"type:varchar(20);not null;default:'pending'"`
	Tanggal time.Time        `json:"tanggal" gorm:"not null;default:now()"`
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
	ID      int              `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID int              `json:"order_id" gorm:"not null;uniqueIndex"`
	Kurir   string           `json:"kurir" gorm:"type:varchar(50);not null"`
	NoResi  string           `json:"no_resi" gorm:"type:varchar(100);not null"`
	Status  PengirimanStatus `json:"status" gorm:"type:varchar(20);not null;default:'proses'"`
	Tanggal time.Time        `json:"tanggal" gorm:"not null;default:now()"`
}

// TelegramConfig holds the Telegram group configuration for a division.
type TelegramConfig struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	NamaGrup   string    `json:"nama_grup" gorm:"type:varchar(100);not null;default:'Default Group'"`
	ChatID     int64     `json:"chat_id" gorm:"not null;default:0"`
	Aktif      bool      `json:"aktif" gorm:"not null;default:true"`
	JamSummary string    `json:"jam_summary" gorm:"type:varchar(5);not null;default:'07:00'"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;default:now()"`
}

// AnomalyConfig stores the anomaly detection threshold for a specific metric.
type AnomalyConfig struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	MetricKey    string    `json:"metric_key" gorm:"type:varchar(100);not null;uniqueIndex"` // e.g. daily_revenue, order_count
	ThresholdPct float64   `json:"threshold_pct" gorm:"type:numeric(5,2);not null"`          // e.g. 10.00 = 10%
	Aktif        bool      `json:"aktif" gorm:"not null;default:true"`
}

// SavedDashboard stores user-saved dashboard preset configuration.
type SavedDashboard struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     *uuid.UUID `json:"user_id,omitempty" gorm:"type:uuid;index"`
	Nama       string     `json:"nama" gorm:"type:varchar(100);not null"`
	ConfigJSON string     `json:"config_json" gorm:"type:jsonb;not null;default:'{}'"`
	CreatedAt  time.Time  `json:"created_at" gorm:"not null;default:now()"`
}

func (User) TableName() string {
	return "app.users"
}

func (TelegramConfig) TableName() string {
	return "app.telegram_config"
}

func (AnomalyConfig) TableName() string {
	return "app.anomaly_config"
}

func (SavedDashboard) TableName() string {
	return "app.saved_dashboards"
}

func (Produk) TableName() string {
	return "bisnis.tbl_produk"
}

func (Customer) TableName() string {
	return "bisnis.tbl_customer"
}

func (Order) TableName() string {
	return "bisnis.tbl_order"
}

func (OrderDetail) TableName() string {
	return "bisnis.tbl_order_detail"
}

func (Pembayaran) TableName() string {
	return "bisnis.tbl_pembayaran"
}

func (Pengiriman) TableName() string {
	return "bisnis.tbl_pengiriman"
}
