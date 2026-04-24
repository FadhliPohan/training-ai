# Task Plan Implementasi — InsightFlow Self-Service AI Dashboard

> **Status Keseluruhan:** 🟡 Dalam Persiapan — `~2%`  
> **Fokus Saat Ini:** Backend (Golang Fiber) → Frontend (Next.js) → *(Testing: ditunda)*  
> **Terakhir Diperbarui:** 24 April 2026

---

## Legenda Status

| Simbol | Arti |
|---|---|
| `[ ]` | Belum dikerjakan |
| `[/]` | Sedang dikerjakan |
| `[x]` | Selesai |
| `%` | Persentase penyelesaian per fase |

---

## 📊 Ringkasan Progres

| Fase | Area | Progress | Status |
|---|---|---|---|
| **Fase 0** | Persiapan & Klarifikasi | `33%` (7/21) | 🟡 Berjalan |
| **Fase 1** | Backend — Setup & Fondasi | `0%` | 🔴 Belum |
| **Fase 2** | Backend — Auth & Master Data | `0%` | 🔴 Belum |
| **Fase 3** | Backend — Transaksi | `0%` | 🔴 Belum |
| **Fase 4** | Backend — Dashboard & AI | `0%` | 🔴 Belum |
| **Fase 5** | n8n — Setup & Workflows | `0%` | 🔴 Belum |
| **Fase 6** | Frontend — Setup & Design System | `0%` | 🔴 Belum |
| **Fase 7** | Frontend — Halaman Toko & Auth | `0%` | 🔴 Belum |
| **Fase 8** | Frontend — Admin & Transaksi | `0%` | 🔴 Belum |
| **Fase 9** | Frontend — Dashboard & AI | `0%` | 🔴 Belum |
| **Fase 10** | Security Hardening | `0%` | 🔴 Belum |
| ~~**Fase 11**~~ | ~~Testing~~ | ~~ditunda~~ | ⏸️ Ditunda |

**Total Keseluruhan: `7 / ~128 task` selesai**

---

## FASE 0 — Persiapan & Klarifikasi PRD

> **Target:** Semua pertanyaan stakeholder terjawab, API docs tersusun, environment siap.  
> **Progress: `33%` (7/21 task)** ✅ 0.1 selesai

### 0.1 Klarifikasi Stakeholder ✅

- [x] **#1 Varian produk:** Satu produk dengan field `ukuran` & `warna` *(bukan SKU terpisah per kombinasi)*
- [x] **#2 Customer auth:** Customer dapat register & login sendiri *(perlu endpoint `POST /auth/register`)*
- [x] **#3 Laporan MVP:** Semua 8 laporan masuk MVP (`daily-sales`, `monthly-sales`, `top-products`, `sales-by-person`, `order-funnel`, `category-breakdown`, `low-stock`, `revenue-trend`)
- [x] **#4 Role `viewer`:** Akses sama seperti admin tetapi **read-only** — tidak bisa tambah/edit/hapus data apapun
- [x] **#5 Threshold anomali:** **Per metrik** *(butuh tabel `app.anomaly_config` terpisah, bukan satu kolom threshold di `telegram_config`)*
- [x] **#6 Daily summary:** Dikirim ke **grup Telegram** (bukan per-user individual)
- [x] **#7 Sales di web:** Sales dapat melihat **semua order tim** *(tidak dibatasi `WHERE sales_id = user_id` di web)*

> [!IMPORTANT]
> **Implikasi Teknis dari Jawaban di Atas:**
> - **#2:** Tambahkan endpoint `POST /auth/register` untuk customer. Tabel `bisnis.tbl_customer` perlu field `password` (bcrypt) dan `aktif`.
> - **#4:** Middleware `RoleGuard` perlu membedakan antara akses **write** dan **read**. Semua endpoint `POST/PUT/PATCH/DELETE` harus diblokir untuk role `viewer`.
> - **#5:** Skema DB perlu tabel baru `app.anomaly_config` dengan kolom `metric_key` dan `threshold_pct` per metrik. Kolom `threshold` di `telegram_config` tidak cukup.

### 0.2 API Documentation

- [ ] Setup Postman Collection atau OpenAPI 3.0
- [ ] Definisikan format respons standar `{ success, message, data, errors }`
- [ ] Definisikan semua error code (400, 401, 403, 404, 409, 500)
- [ ] Dokumentasikan semua 21 endpoint dengan contoh request + response
- [ ] Hosting Swagger UI di `/api/docs`
- [ ] Commit Postman Collection `.json` ke repo

### 0.3 Setup Repository & Environment

- [ ] Inisialisasi Git repository dengan struktur folder `be-penjualan/` dan `fe-penjualan/`
- [ ] Buat `.gitignore` (pastikan `.env` masuk)
- [ ] Buat `docker-compose.yml` untuk PostgreSQL + n8n (development)
- [ ] Buat `.env.example` dengan semua variable yang dibutuhkan

---

## FASE 1 — Backend: Setup & Fondasi

> **Target:** Go Fiber project berjalan, database schema terbuat, migration tools siap.  
> **Progress: `0%`** | Lokasi: `be-penjualan/`

### 1.1 Init Project Go Fiber

- [ ] `go mod init` dengan nama module yang sesuai
- [ ] Install dependency utama:
  - `github.com/gofiber/fiber/v2`
  - `github.com/jackc/pgx/v5` (pgxpool)
  - `github.com/golang-migrate/migrate/v4`
  - `github.com/golang-jwt/jwt/v5`
  - `github.com/go-playground/validator/v10`
  - `github.com/rs/zerolog`
  - `github.com/joho/godotenv`
- [ ] Setup struktur folder:
  ```
  /cmd
  /internal
    /handler
    /service
    /repository
    /middleware
    /domain
  /config
  /db/migrations
  ```
- [ ] Setup konfigurasi dari `.env` (DB DSN, JWT Secret, n8n URL, Telegram Token)
- [ ] Setup PostgreSQL connection pool (`pgxpool`)

### 1.2 Database Migration

- [ ] Setup migration tool (`golang-migrate`)
- [ ] Migration: `CREATE SCHEMA app`
- [ ] Migration: `CREATE SCHEMA bisnis`
- [ ] Migration: `app.users` (id, nama, email, password, role, telegram_user_id, aktif, created_at)
- [ ] Migration: `app.telegram_config` (id, nama_grup, chat_id, aktif, jam_summary, threshold)
- [ ] Migration: `app.saved_dashboards` (id, user_id, nama, konfigurasi JSONB, created_at)
- [ ] Migration: `bisnis.tbl_produk` (id, kode_produk, nama, kategori_pakaian, ukuran, warna, bahan, harga, stok, aktif)
- [ ] Migration: `bisnis.tbl_customer` (id, kode_cust, nama, email, telepon, alamat, created_at)
- [ ] Migration: `bisnis.tbl_order` (id, no_order, customer_id, sales_id, tanggal, status, total, created_at)
- [ ] Migration: `bisnis.tbl_order_detail` (id, order_id, produk_id, qty, harga_saat, subtotal)
- [ ] Migration: `bisnis.tbl_pembayaran` (id, order_id, jumlah, metode, status, tanggal)
- [ ] Migration: `bisnis.tbl_pengiriman` (id, order_id, kurir, no_resi, status, tanggal)

### 1.3 Indexing Database

- [ ] `CREATE INDEX ON bisnis.tbl_order(tanggal)`
- [ ] `CREATE INDEX ON bisnis.tbl_order(sales_id)`
- [ ] `CREATE INDEX ON bisnis.tbl_order(status)`
- [ ] `CREATE INDEX ON bisnis.tbl_produk(kategori_pakaian)`
- [ ] `CREATE INDEX ON app.users(telegram_user_id)`

### 1.4 Seeder Data Development

- [ ] Seeder: user admin + manager + 3 sales (dengan bcrypt password)
- [ ] Seeder: 20 produk pakaian dummy (berbagai kategori, ukuran, warna)
- [ ] Seeder: 10 customer dummy
- [ ] Seeder: sample order + detail + pembayaran + pengiriman

### 1.5 Non-Functional Setup

- [ ] `GET /health` — health check endpoint
- [ ] Structured logging dengan `zerolog`
- [ ] Rate limiting middleware (Fiber)
- [ ] Request timeout 30 detik
- [ ] Graceful shutdown handler (SIGTERM)
- [ ] CORS middleware (whitelist frontend URL)

---

## FASE 2 — Backend: Auth & Master Data

> **Target:** Login/logout berfungsi, CRUD Produk/Customer/User selesai.  
> **Progress: `0%`**

### 2.1 Modul Auth

- [ ] Domain struct: `User`, JWT Claims
- [ ] Repository: `FindUserByEmail`, `FindUserByID`
- [ ] Service: `Login` (validasi email+password, generate JWT), `GetProfile`
- [ ] Handler: `POST /auth/login`
- [ ] Handler: `POST /auth/logout`
- [ ] Handler: `GET /auth/me`
- [ ] Middleware: `AuthRequired` (validasi JWT dari header/cookie)
- [ ] Middleware: `RoleGuard(roles ...string)` (cek role dari JWT claim)
- [ ] JWT: HS256, expiry 8 jam, secret dari env

### 2.2 Modul Produk (Master Data)

- [ ] Domain struct: `Produk`
- [ ] Repository: `FindAll`, `FindByID`, `FindAktif`, `Create`, `Update`, `SoftDelete`
- [ ] Service: validasi input, business logic
- [ ] Handler: `GET /produk` (dengan filter `?aktif=true` untuk dropdown order)
- [ ] Handler: `GET /produk/:id`
- [ ] Handler: `POST /produk` (Admin only)
- [ ] Handler: `PUT /produk/:id` (Admin only)
- [ ] Handler: `PATCH /produk/:id` (soft-delete: `aktif = false`)
- [ ] Validasi: nama, harga, kategori_pakaian wajib diisi

### 2.3 Modul Customer (Master Data)

- [ ] Domain struct: `Customer`
- [ ] Repository: `FindAll`, `FindByID`, `Create`, `Update`
- [ ] Handler: `GET /customer`
- [ ] Handler: `GET /customer/:id`
- [ ] Handler: `POST /customer`
- [ ] Handler: `PUT /customer/:id`

### 2.4 Modul User/Sales (Master Data)

- [ ] Repository: `FindAll`, `FindByID`, `Create`, `Update`, `Deactivate`
- [ ] Handler: `GET /users` (Admin only)
- [ ] Handler: `GET /users/:id` (Admin only)
- [ ] Handler: `POST /users` — buat user baru + set role + `telegram_user_id`
- [ ] Handler: `PUT /users/:id` — edit profil + role
- [ ] Handler: `PATCH /users/:id` — nonaktifkan akun
- [ ] Hash password baru dengan bcrypt cost factor ≥ 12

### 2.5 Settings Telegram

- [ ] Handler: `GET /settings/telegram` (Admin only)
- [ ] Handler: `PUT /settings/telegram` — update chat_id, jam_summary, threshold

---

## FASE 3 — Backend: Transaksi

> **Target:** Seluruh siklus order dari pending → closed dapat diproses via API.  
> **Progress: `0%`**

### 3.1 Modul Order

- [ ] Domain struct: `Order`, `OrderDetail`, `OrderStatus`
- [ ] Repository: `FindAll` (dengan filter status, sales_id, tanggal), `FindByID`, `Create`
- [ ] Service: `CreateOrder` — validasi stok, **atomic transaction** (insert order + semua detail sekaligus)
- [ ] Handler: `GET /orders` (filter: status, from, to, sales_id)
- [ ] Handler: `GET /orders/:id`
- [ ] Handler: `POST /orders` — buat order baru
- [ ] Handler: `POST /orders/:id/confirm` → status: `confirmed`
- [ ] Handler: `POST /orders/:id/cancel` → status: `cancelled` + field alasan

### 3.2 Modul Pembayaran

- [ ] Domain struct: `Pembayaran`
- [ ] Repository: `FindByOrderID`, `Create`, `UpdateStatus`
- [ ] Handler: `POST /payments` — catat pembayaran
- [ ] Handler: `POST /payments/:id/verify` → status: `verified`, order status → `paid`

### 3.3 Modul Pengiriman

- [ ] Domain struct: `Pengiriman`
- [ ] Repository: `FindByOrderID`, `Create`, `Update`
- [ ] Handler: `POST /shipments` — catat nomor resi + kurir, order status → `shipped`
- [ ] Handler: `PUT /shipments/:id` → status: `diterima`, order status → `closed`

---

## FASE 4 — Backend: Dashboard & AI Integration

> **Target:** Endpoint laporan aggregasi berfungsi dan terintegrasi dengan n8n.  
> **Progress: `0%`**

### 4.1 Modul Laporan (Reports)

- [ ] Implementasi query aggregasi per tipe laporan:
  - [ ] `daily-sales` — penjualan per hari dalam rentang tanggal
  - [ ] `monthly-sales` — penjualan per bulan
  - [ ] `top-products` — produk terlaris berdasarkan qty/revenue
  - [ ] `sales-by-person` — performa per sales
  - [ ] `order-funnel` — count order per status
  - [ ] `category-breakdown` — penjualan per kategori pakaian
  - [ ] `low-stock` — produk dengan stok di bawah threshold
  - [ ] `revenue-trend` — tren pendapatan per periode
- [ ] Handler: `GET /reports?type=&from=&to=&sales_id=`
- [ ] Whitelist validasi nilai `type` (hanya nilai yang diizinkan)
- [ ] Kirim data aggregat ke n8n webhook
- [ ] Terima respons n8n: `chart_type`, `summary`, `anomalies[]`, `recommendation`
- [ ] Gabungkan data + AI response → return JSON ke frontend

### 4.2 Modul AI Chat (SSE)

- [ ] Handler: `GET /chat/stream?message=`
- [ ] Set SSE headers: `Content-Type: text/event-stream`, `Cache-Control: no-cache`
- [ ] Query produk relevan dari DB berdasarkan keyword dari pesan
- [ ] POST ke n8n chat workflow dengan pesan + context produk
- [ ] Forward stream response dari n8n ke client via SSE

---

## FASE 5 — n8n: Setup & Semua Workflow

> **Target:** 5 workflow AI berjalan dan terkoneksi ke LLM + Telegram.  
> **Progress: `0%`**

### 5.1 Setup n8n

- [ ] Deploy n8n via Docker Compose (`restart: always`)
- [ ] Konfigurasi: hanya dapat diakses via internal network
- [ ] Setup credentials di n8n:
  - [ ] OpenAI API Key / Gemini API Key
  - [ ] Telegram Bot Token
  - [ ] PostgreSQL connection string
- [ ] Set environment variables: backend base URL, webhook secret

### 5.2 Workflow 1 — Dashboard AI Analysis

- [ ] HTTP Webhook node (terima data dari Golang)
- [ ] Code node — format data menjadi prompt
- [ ] OpenAI/Gemini node — kirim ke LLM
- [ ] Code node — parse & validasi JSON respons LLM
- [ ] Respond to Webhook — kembalikan JSON ke Golang
- [ ] Error Trigger node + fallback message

### 5.3 Workflow 2 — AI Chat Assistant

- [ ] HTTP Webhook node (terima pesan + context produk)
- [ ] Code node — susun system prompt + context
- [ ] OpenAI node — generate jawaban
- [ ] Respond to Webhook — kembalikan jawaban ke Golang
- [ ] Error Trigger node + fallback message

### 5.4 Workflow 3 — Telegram Daily Summary

- [ ] Schedule Trigger — `0 0 * * *` (07:00 WIB)
- [ ] PostgreSQL node — aggregasi penjualan 24 jam terakhir
- [ ] Code node — format data + susun prompt
- [ ] OpenAI/Gemini node — generate ringkasan harian
- [ ] Code node — format pesan Telegram
- [ ] Telegram node — `sendMessage` ke semua chat_id aktif
- [ ] Error Trigger node + retry logic

### 5.5 Workflow 4 — Telegram Anomaly Alert

- [ ] Schedule Trigger — `*/15 * * * *` (setiap 15 menit)
- [ ] PostgreSQL node — data terbaru vs rata-rata 7 hari
- [ ] Code node — hitung varians vs threshold dari DB
- [ ] IF node — varians > threshold?
- [ ] OpenAI/Gemini node — formulasikan pesan alert
- [ ] Telegram node — `sendMessage` ke grup
- [ ] Error Trigger node + retry 2x delay 30 detik

### 5.6 Workflow 5 — Telegram Q&A Per-Role

- [ ] Telegram Trigger node (webhook dari Telegram API)
- [ ] PostgreSQL node — lookup `telegram_user_id` di `app.users`
- [ ] IF node — user terdaftar?
- [ ] Switch node — route berdasarkan role (sales/manager)
- [ ] Code node — susun prompt + data sesuai scope role
- [ ] OpenAI/Gemini node — generate jawaban
- [ ] Telegram node — `sendMessage` ke `chat_id` pengirim
- [ ] Error Trigger node + fallback message

---

## FASE 6 — Frontend: Setup & Design System

> **Target:** Next.js project berjalan, design system selesai, semua base component tersedia.  
> **Progress: `0%`** | Lokasi: `fe-penjualan/`  
> ⚠️ *Fase ini dimulai setelah Fase 1-3 Backend selesai*

### 6.1 Init Project Next.js

- [ ] `npx create-next-app@latest ./` dengan TypeScript + Tailwind + App Router
- [ ] Install dependency:
  - `shadcn/ui` — base component library
  - `@tanstack/react-query` — data fetching
  - `axios` — HTTP client
  - `recharts` atau `echarts-for-react` — charting library
  - `html2canvas` + `jspdf` — export PDF
  - `js-cookie` — cookie management
- [ ] Setup Tailwind config dengan custom token warna (biru-putih + dark navy)
- [ ] Setup Axios instance dengan JWT interceptor
- [ ] Setup React Query provider
- [ ] Setup font Inter dari Google Fonts

### 6.2 Design System & Base Components

- [ ] `Button` — variant: primary, secondary, danger, ghost (dengan loading state)
- [ ] `Input` — label + error state + helper text
- [ ] `Select` — label + error state
- [ ] `Textarea` — label + error state
- [ ] `Modal` / `Dialog` — dengan backdrop + close button
- [ ] `Badge` / `StatusBadge` — untuk semua status order
- [ ] `DataTable` — sortable + pagination + loading skeleton
- [ ] `Alert` / `Toast` — notifikasi sukses/error
- [ ] `Skeleton` — loader placeholder
- [ ] `EmptyState` — tampilan data kosong
- [ ] `PageHeader` — judul + breadcrumb + action slot
- [ ] `Sidebar` — collapsible, dengan active state per route
- [ ] `TopNavbar` — user info + logout

---

## FASE 7 — Frontend: Halaman Toko & Auth

> **Target:** Halaman publik toko berjalan, login/logout berfungsi, route protection aktif.  
> **Progress: `0%`**

### 7.1 Halaman Toko (Public)

- [ ] `/` — Beranda: hero section, grid produk, filter kategori/ukuran/warna
- [ ] `/produk/:id` — Detail produk: foto, info lengkap, pilih ukuran & warna, tombol order
- [ ] `ChatWidget` — floating button pojok kanan bawah
  - [ ] Panel chat slide-up/slide-down
  - [ ] Input pesan + tombol kirim
  - [ ] SSE streaming: tampilkan teks karakter demi karakter
  - [ ] Typing indicator saat AI memproses
  - [ ] Mobile responsive

### 7.2 Auth

- [ ] `/login` — Form email + password + validasi inline Bahasa Indonesia
- [ ] Protected Route HOC — redirect ke `/login` jika tidak ada token valid
- [ ] Auto redirect ke `/login` saat token expired (via Axios interceptor)
- [ ] Logout: hapus token + redirect ke `/login`

---

## FASE 8 — Frontend: Admin & Transaksi

> **Target:** Seluruh halaman CRUD Master Data dan alur transaksi order selesai.  
> **Progress: `0%`**

### 8.1 Admin — Master Data

- [ ] `/admin/produk` — DataTable + tombol Tambah/Edit/Nonaktifkan
- [ ] `/admin/produk/baru` — Form produk baru (nama, kategori, ukuran, warna, bahan, harga, stok)
- [ ] `/admin/produk/:id/edit` — Form edit produk
- [ ] `/admin/customer` — DataTable customer + CRUD modal
- [ ] `/admin/users` — DataTable user + form tambah + set role + telegram_user_id
- [ ] `/settings/telegram` — Form konfigurasi Telegram (chat_id, jam_summary, threshold)

### 8.2 Sales — Transaksi

- [ ] `/sales/orders` — DataTable order + filter status (dropdown)
- [ ] `/sales/orders/baru` — Form order:
  - [ ] Select customer (autocomplete/searchable)
  - [ ] Select produk (searchable + tampilkan stok)
  - [ ] Field qty, ukuran, warna
  - [ ] Validasi stok tidak cukup secara real-time
- [ ] `/sales/orders/:id` — Detail order:
  - [ ] Stepper visual: Pending → Confirmed → Paid → Shipped → Closed
  - [ ] Tombol aksi sesuai status saat ini
  - [ ] Modal konfirmasi setiap perubahan status
  - [ ] Form pembayaran (jika status confirmed)
  - [ ] Form pengiriman / nomor resi (jika status paid)

---

## FASE 9 — Frontend: Dashboard & AI

> **Target:** Dashboard analitik dengan chart + AI insight berfungsi penuh.  
> **Progress: `0%`**

- [ ] `/dashboard` — Halaman utama dashboard:
  - [ ] `ReportSelector` — dropdown pilih laporan + deskripsi
  - [ ] `FilterBar` — date range picker, filter sales, filter kategori
  - [ ] Tombol "Tampilkan Laporan"
  - [ ] `ChartRenderer` — render Line/Bar/Pie/Funnel sesuai `chart_type` dari AI
  - [ ] `AIInsightCard` — card biru dengan ringkasan AI 2-3 kalimat
  - [ ] `AnomalyFlag` — banner ⚠️ + penjelasan + rekomendasi (tampil jika ada anomali)
  - [ ] Skeleton loader saat data sedang diproses
  - [ ] Tombol "Download PDF" (`html2canvas` + `jsPDF`)

---

## FASE 10 — Security Hardening

> **Target:** Semua checklist security terpenuhi sebelum UAT.  
> **Progress: `0%`**

- [ ] Audit: semua endpoint melewati `AuthRequired` middleware
- [ ] Audit: endpoint sensitif melewati `RoleGuard`
- [ ] Audit: query sales selalu filter `WHERE sales_id = user_id`
- [ ] Audit: tidak ada secret yang ter-commit ke repository
- [ ] Test: rate limit login aktif (coba login > 5x dalam 1 menit)
- [ ] Test: token yang expired redirect ke halaman login
- [ ] Audit: payload yang dikirim ke n8n/LLM tidak mengandung data personal
- [ ] Konfigurasi: security headers di Nginx (`X-Frame-Options`, `X-Content-Type-Options`, `HSTS`)
- [ ] Konfigurasi: CORS whitelist hanya domain frontend
- [ ] Konfigurasi: DB user hanya punya privilege `SELECT`, `INSERT`, `UPDATE`
- [ ] Verifikasi: n8n tidak bisa diakses dari luar internal network

---

## ~~FASE 11 — Testing~~ *(Ditunda)*

> **Status: ⏸️ Ditunda — akan dijadwalkan ulang setelah semua fitur core selesai.**

Item yang akan dikerjakan saat testing:
- ~~Unit test untuk service layer (auth, transaksi, laporan)~~
- ~~Integration test untuk endpoint kritis~~
- ~~User Acceptance Testing (UAT) dengan minimal 3 pengguna non-teknis~~
- ~~Load testing untuk endpoint dashboard dan SSE chat~~

---

## Urutan Pengerjaan yang Direkomendasikan

```
Minggu 1-2 (Fase 0 + 1):
  → Klarifikasi PRD + API Docs + Setup Go Fiber + DB Migration

Minggu 3-4 (Fase 2 + 3):
  → Backend Auth + CRUD Master Data + Modul Transaksi

Minggu 5-6 (Fase 4 + 5):
  → Backend Dashboard + SSE Chat + n8n Setup + 5 Workflows

Minggu 7-8 (Fase 6 + 7):
  → Frontend Setup + Design System + Halaman Toko + Auth

Minggu 9-10 (Fase 8 + 9):
  → Frontend Admin + Transaksi + Dashboard AI

Minggu 11 (Fase 10):
  → Security Hardening + Pre-launch checklist

Minggu 12:
  → Buffer / Bug Fix / Go-Live
```

---

*Dokumen ini diperbarui setiap ada progress pengerjaan. Tandai task dengan `[x]` saat selesai dan update persentase di tabel ringkasan.*
