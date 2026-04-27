package database

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedData inserts minimum required seed data if not already present.
// Idempotent: uses INSERT ... ON CONFLICT DO NOTHING.
func SeedData() {
	if Pool == nil {
		log.Fatal("[seed] database pool is not initialised")
	}

	ctx := context.Background()

	log.Println("[seed] starting seed data insertion...")

	seedUsers(ctx)
	seedCustomers(ctx)
	seedProduk(ctx)
	seedOrders(ctx)
	seedTelegramConfig(ctx)

	log.Println("[seed] seed data insertion complete")
}

func hashPassword(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("[seed] failed to hash password: %v", err)
	}
	return string(hash)
}

func seedUsers(ctx context.Context) {
	type seedUser struct {
		nama     string
		email    string
		password string
		role     string
	}

	users := []seedUser{
		{"Administrator", "admin@insightflow.id", "Admin@12345", "admin"},
		{"Budi Manager", "manager@insightflow.id", "Manager@12345", "manager"},
		{"Citra Sales", "sales@insightflow.id", "Sales@12345", "sales"},
		{"Doni Viewer", "viewer@insightflow.id", "Viewer@12345", "viewer"},
	}

	const q = `
		INSERT INTO app.users (nama, email, password, role, aktif)
		VALUES ($1, $2, $3, $4, true)
		ON CONFLICT (email) DO NOTHING
	`

	for _, u := range users {
		hashed := hashPassword(u.password)
		_, err := Pool.Exec(ctx, q, u.nama, u.email, hashed, u.role)
		if err != nil {
			log.Printf("[seed] failed to insert user %s: %v", u.email, err)
		} else {
			fmt.Printf("[seed] user seeded: %s (%s)\n", u.email, u.role)
		}
	}
}

func seedCustomers(ctx context.Context) {
	type seedCustomer struct {
		kode    string
		nama    string
		email   string
		telepon string
		alamat  string
	}

	customers := []seedCustomer{
		{"CUST-001", "PT Maju Bersama", "maju@example.com", "021-12345678", "Jl. Merdeka No. 1, Jakarta"},
		{"CUST-002", "CV Sejahtera Jaya", "sejahtera@example.com", "022-87654321", "Jl. Pahlawan No. 5, Bandung"},
		{"CUST-003", "Toko Busana Indah", "busana@example.com", "031-11223344", "Jl. Pemuda No. 10, Surabaya"},
	}

	const q = `
		INSERT INTO bisnis.tbl_customer (kode_cust, nama, email, telepon, alamat, aktif)
		VALUES ($1, $2, $3, $4, $5, true)
		ON CONFLICT (email) DO NOTHING
	`

	for _, c := range customers {
		_, err := Pool.Exec(ctx, q, c.kode, c.nama, c.email, c.telepon, c.alamat)
		if err != nil {
			log.Printf("[seed] failed to insert customer %s: %v", c.email, err)
		} else {
			fmt.Printf("[seed] customer seeded: %s\n", c.nama)
		}
	}
}

func seedProduk(ctx context.Context) {
	type seedProduk struct {
		kode      string
		nama      string
		kategori  string
		ukuran    string
		warna     string
		bahan     string
		harga     float64
		stok      int
	}

	produks := []seedProduk{
		{"PRD-001", "Kemeja Batik Lengan Panjang", "Kemeja", "L", "Biru Navy", "Katun", 185000, 50},
		{"PRD-002", "Kaos Polos Premium", "Kaos", "M", "Putih", "Cotton Combed 30s", 75000, 100},
		{"PRD-003", "Celana Chino Slim Fit", "Celana", "32", "Khaki", "Twill Cotton", 220000, 30},
		{"PRD-004", "Jaket Bomber Unisex", "Jaket", "XL", "Hitam", "Polyester", 350000, 20},
		{"PRD-005", "Dress Batik Sogan", "Dress", "S", "Coklat Sogan", "Sutera", 450000, 15},
	}

	const q = `
		INSERT INTO bisnis.tbl_produk (kode_produk, nama, kategori_pakaian, ukuran, warna, bahan, harga, stok, aktif)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, true)
		ON CONFLICT (kode_produk) DO NOTHING
	`

	for _, p := range produks {
		_, err := Pool.Exec(ctx, q, p.kode, p.nama, p.kategori, p.ukuran, p.warna, p.bahan, p.harga, p.stok)
		if err != nil {
			log.Printf("[seed] failed to insert produk %s: %v", p.kode, err)
		} else {
			fmt.Printf("[seed] produk seeded: %s\n", p.nama)
		}
	}
}

func seedOrders(ctx context.Context) {
	// 1. Get Sales User ID
	var salesID string
	err := Pool.QueryRow(ctx, "SELECT id FROM app.users WHERE role = 'sales' LIMIT 1").Scan(&salesID)
	if err != nil {
		log.Printf("[seed] failed to get sales user: %v", err)
		return
	}

	// 2. Get Customer ID
	var customerID int
	err = Pool.QueryRow(ctx, "SELECT id FROM bisnis.tbl_customer WHERE kode_cust = 'CUST-001' LIMIT 1").Scan(&customerID)
	if err != nil {
		log.Printf("[seed] failed to get customer: %v", err)
		return
	}

	// 3. Get Produk ID and Harga
	var produkID int
	var harga float64
	err = Pool.QueryRow(ctx, "SELECT id, harga FROM bisnis.tbl_produk WHERE kode_produk = 'PRD-001' LIMIT 1").Scan(&produkID, &harga)
	if err != nil {
		log.Printf("[seed] failed to get produk: %v", err)
		return
	}

	// 4. Insert Order
	const qOrder = `
		INSERT INTO bisnis.tbl_order (no_order, customer_id, sales_id, tanggal, status, total)
		VALUES ($1, $2, $3, CURRENT_DATE, $4, $5)
		ON CONFLICT (no_order) DO NOTHING
		RETURNING id
	`
	var orderID int
	qty := 2
	total := harga * float64(qty)
	
	err = Pool.QueryRow(ctx, qOrder, "ORD-2026-001", customerID, salesID, "paid", total).Scan(&orderID)
	if err != nil {
		// Jika order sudah ada, err adalah no rows in result set, kita bisa cari ID-nya
		err = Pool.QueryRow(ctx, "SELECT id FROM bisnis.tbl_order WHERE no_order = 'ORD-2026-001' LIMIT 1").Scan(&orderID)
		if err != nil {
			log.Printf("[seed] failed to insert/get order: %v", err)
			return
		}
	} else {
		fmt.Printf("[seed] order seeded: ORD-2026-001\n")
		
		// 5. Insert Order Detail
		const qDetail = `
			INSERT INTO bisnis.tbl_order_detail (order_id, produk_id, qty, harga_saat, subtotal)
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = Pool.Exec(ctx, qDetail, orderID, produkID, qty, harga, total)
		if err != nil {
			log.Printf("[seed] failed to insert order detail: %v", err)
		}

		// 6. Insert Pembayaran
		const qBayar = `
			INSERT INTO bisnis.tbl_pembayaran (order_id, jumlah, metode, status)
			VALUES ($1, $2, $3, $4)
		`
		_, err = Pool.Exec(ctx, qBayar, orderID, total, "transfer", "verified")
		if err != nil {
			log.Printf("[seed] failed to insert pembayaran: %v", err)
		}

		// 7. Insert Pengiriman
		const qKirim = `
			INSERT INTO bisnis.tbl_pengiriman (order_id, kurir, no_resi, status)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (order_id) DO NOTHING
		`
		_, err = Pool.Exec(ctx, qKirim, orderID, "JNE", "JNE1234567890", "proses")
		if err != nil {
			log.Printf("[seed] failed to insert pengiriman: %v", err)
		}
	}
}

func seedTelegramConfig(ctx context.Context) {
	type seedTele struct {
		namaGrup   string
		chatID     int64
		jamSummary string
	}

	configs := []seedTele{
		{"Sales Daily Updates", -1001234567890, "07:00"},
		{"Management Dashboard Alerts", -1009876543210, "08:00"},
	}

	// Menggunakan chat_id sebagai basis pengecekan duplicate
	const q = `
		INSERT INTO app.telegram_config (nama_grup, chat_id, jam_summary, aktif)
		VALUES ($1, $2, $3, true)
		ON CONFLICT DO NOTHING
	`

	for _, c := range configs {
		// PostgreSQL uuid tidak punya default unique constraints untuk chat_id,
		// jadi kita gunakan pengecekan manual atau jalankan query sederhana
		var exists bool
		err := Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM app.telegram_config WHERE chat_id = $1)", c.chatID).Scan(&exists)
		if err == nil && !exists {
			_, err = Pool.Exec(ctx, q, c.namaGrup, c.chatID, c.jamSummary)
			if err != nil {
				log.Printf("[seed] failed to insert telegram config %s: %v", c.namaGrup, err)
			} else {
				fmt.Printf("[seed] telegram config seeded: %s\n", c.namaGrup)
			}
		} else if err != nil {
			log.Printf("[seed] failed to check telegram config: %v", err)
		}
	}
}
