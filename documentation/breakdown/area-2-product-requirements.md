# Area 2 — Product Requirements (Klarifikasi & Finalisasi)

> **Tujuan:** Memastikan semua pertanyaan bisnis yang berdampak pada implementasi teknis sudah dijawab oleh stakeholder sebelum development dimulai.

---

## 2.1 Pertanyaan yang Harus Dijawab Stakeholder

> [!IMPORTANT]
> Semua pertanyaan di bawah ini **wajib dijawab** sebelum development dimulai. Setiap jawaban berdampak langsung pada desain database, UI, dan logika backend.

| # | Pertanyaan | Dampak Teknis |
|---|---|---|
| 1 | Varian produk: satu SKU per kombinasi ukuran+warna, atau satu produk dengan field ukuran/warna? | Desain tabel DB & UI form |
| 2 | Customer bisa buat akun sendiri, atau order selalu via sales? | Auth flow, tabel customer |
| 3 | Daftar laporan di dropdown dashboard: mana yang wajib MVP? | Backend query, n8n prompt |
| 4 | Role `viewer`: apa saja yang bisa dilihat? | Middleware RoleGuard |
| 5 | Threshold anomali: global atau bisa per-metrik? | Tabel config, n8n logic |
| 6 | Daily summary: dikirim ke grup Telegram atau per-user individual? | n8n workflow |
| 7 | Sales bisa lihat semua order tim atau hanya order sendiri di dashboard web? | Query + RoleGuard |

---

## 2.2 Daftar Laporan Dashboard (Usulan MVP)

Laporan-laporan ini tersedia sebagai pilihan di dropdown dashboard. AI akan menentukan tipe chart yang paling tepat secara otomatis.

| Kode | Nama Laporan | Tipe Chart (Default) |
|---|---|---|
| `daily-sales` | Penjualan Harian | Line Chart |
| `monthly-sales` | Penjualan Bulanan | Line Chart |
| `top-products` | Produk Terlaris | Bar Chart |
| `sales-by-person` | Penjualan per Sales | Bar Chart |
| `order-funnel` | Funnel Status Order | Funnel Chart |
| `category-breakdown` | Penjualan per Kategori Pakaian | Pie Chart |
| `low-stock` | Produk Stok Rendah | Table |
| `revenue-trend` | Tren Pendapatan | Line Chart |

---

## 2.3 Definisi Role & Hak Akses

| Role | Hak Akses |
|---|---|
| `admin` | Full access: CRUD semua data, kelola user, konfigurasi Telegram, lihat semua laporan |
| `manager` | Lihat semua laporan, semua data order tim, tidak bisa CRUD user |
| `sales` | Input order, lihat order sendiri, lihat laporan data sendiri via Telegram |
| `viewer` | *(Perlu klarifikasi dari stakeholder — lihat pertanyaan no. 4)* |

---

## 2.4 Daftar Laporan yang Disarankan untuk MVP

> [!TIP]
> Berdasarkan analisis kebutuhan bisnis, **minimal 5 laporan** harus tersedia saat go-live MVP:
> 1. `daily-sales` — Kebutuhan monitoring harian manajer
> 2. `top-products` — Identifikasi produk unggulan
> 3. `sales-by-person` — Evaluasi performa sales
> 4. `order-funnel` — Deteksi bottleneck dalam proses order
> 5. `low-stock` — Mencegah kehabisan stok

---

## Catatan Penting

> [!WARNING]
> Jika pertanyaan no. 1 (varian produk) dijawab dengan **"satu produk dengan field ukuran/warna"**, maka satu record di `tbl_produk` merepresentasikan satu kombinasi ukuran+warna. Ini mempengaruhi cara stok dihitung dan cara produk ditampilkan di katalog.

> [!NOTE]
> Keputusan pada pertanyaan no. 6 menentukan apakah `telegram_config` perlu menyimpan beberapa `chat_id` (grup) atau `user_id` (personal). Desain tabel saat ini mendukung **grup Telegram** sebagai target utama.
