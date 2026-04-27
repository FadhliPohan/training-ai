# Sprint Task Planning — InsightFlow (Single Source of Truth)

> **Dokumen sprint resmi (satu-satunya):** gunakan file ini untuk planning dan tracking progress.  
> **Terakhir diperbarui:** 24 April 2026  
> **Status saat ini:** Pre-Sprint 1 (backlog sinkron dengan backend + auto-migrate env policy aktif)

---

## 1) Konsolidasi Dokumen

Dokumen ini merangkum materi dari:
- `Self_Service_AI_Dashboard_PRD.md`
- `work_breakdown.md`
- `SIMPLIFIED_SPRINT_PLAN.md`
- `implementation_sprint_plan.md`
- `SPRINT_1_PLANNING_MEETING.md`
- `TODAY_TASKS.md`

Catatan:
- PRD tetap dipakai sebagai referensi requirement produk.
- Untuk planning eksekusi sprint, hanya pakai file ini.

---

## 2) Backend Current State (Audit Kode Nyata)

Audit dilakukan langsung ke `be-penjualan/` pada 24 April 2026.

### Endpoint aktif saat ini

- `GET /health`
- `GET /swagger/*`
- `POST /api/v1/auth/login`

### Status implementasi backend

| Area | Status | Evidence |
|---|---|---|
| App bootstrap (Fiber, logger, CORS, timeout, graceful shutdown) | ✅ Selesai | `cmd/main.go` |
| Config env (`DATABASE_URL`, `JWT_SECRET`, n8n vars) | ✅ Selesai | `config/config.go` |
| DB connection pool (pgx) | ✅ Selesai | `internal/database/database.go` |
| Auto-migrate startup (GORM model + tag) | ✅ Selesai | `internal/database/database.go` + `internal/domain/domain.go` |
| Policy migration per environment | ✅ Selesai | `config/config.go` + `cmd/main.go` (`staging/production` run, `development` skip) |
| Domain model (`User`, `Produk`, `Order`, dll) | ✅ Selesai | `internal/domain/domain.go` |
| DTO auth/product/customer/order/payment/report/settings | ✅ Selesai (scaffold) | `internal/dto/*.go` |
| Response envelope standard | ✅ Selesai | `internal/response/response.go` |
| Middleware `AuthRequired`, `RoleGuard`, `ViewerReadOnly` | ✅ Selesai | `internal/middleware/auth.go` |
| Auth `POST /auth/login` | ✅ Selesai | `internal/handler/auth/login.go` |
| Auth `POST /auth/logout`, `GET /auth/me`, `POST /auth/register` | ❌ Belum | route masih comment |
| Produk, Customer, Users, Orders, Payments, Shipments handlers | ❌ Belum | route masih comment |
| Reports, Chat SSE, Settings Telegram | ❌ Belum | route masih comment |
| Unit/feature test stabil | ✅ Dasar stabil | `go test ./...` lulus, panic `database.Pool` nil sudah ditangani |

### Gap kritikal yang mempengaruhi sprint

- Layer repository/service belum dibentuk; login masih query DB langsung di handler.
- Belum ada validasi end-to-end auto-migrate di environment staging/production sesungguhnya.

---

## 3) Sprint Timeline (2 Minggu per Sprint)

| Sprint | Periode | Fokus | Target Kumulatif |
|---|---|---|---|
| Sprint 1 | 27 Apr 2026 - 08 Mei 2026 | Stabilkan fondasi backend + Auth lengkap + Master data dasar | 30% |
| Sprint 2 | 11 Mei 2026 - 22 Mei 2026 | Transaksi order/payment/shipment | 50% |
| Sprint 3 | 25 Mei 2026 - 05 Jun 2026 | Reports + Integrasi n8n + Telegram settings | 68% |
| Sprint 4 | 08 Jun 2026 - 19 Jun 2026 | Frontend foundation + auth + public catalog | 80% |
| Sprint 5 | 22 Jun 2026 - 03 Jul 2026 | Frontend admin/sales + dashboard AI | 92% |
| Sprint 6 | 06 Jul 2026 - 17 Jul 2026 | Security, QA, UAT, release candidate | 100% |

---

## 4) Sprint Backlog Terpadu

## Sprint 1 — Backend Foundation Hardening + Auth + Master Data Dasar

**Goal:** backend siap dipakai frontend untuk auth dan master data inti.

### A. Fondasi & stabilitas

- [x] Setup Fiber app + global middleware + graceful shutdown
- [x] Setup config + DB pool
- [x] Setup auto-migrate GORM model-based di startup
- [x] Terapkan policy env: staging/production run, development skip
- [ ] Validasi auto-migrate end-to-end di environment staging
- [ ] Seed manager + sales + customer dummy
- [x] Perbaiki test agar tidak panic jika DB belum init

### B. Auth module

- [x] `POST /auth/login`
- [ ] `POST /auth/logout`
- [ ] `GET /auth/me`
- [ ] `POST /auth/register` (customer self-registration sesuai PRD)
- [ ] Refactor login ke service/repository (hindari query langsung di handler)

### C. Master data module

- [ ] Produk: `GET /produk`, `GET /produk/:id`, `POST /produk`, `PUT /produk/:id`, `PATCH /produk/:id`
- [ ] Customer: `GET /customer`, `GET /customer/:id`, `POST /customer`, `PUT /customer/:id`
- [ ] Users: `GET /users`, `GET /users/:id`, `POST /users`, `PUT /users/:id`, `PATCH /users/:id`

### D. Acceptance criteria Sprint 1

- [ ] Admin login, ambil profile (`/auth/me`), dan logout berhasil.
- [ ] Viewer tidak bisa `POST/PUT/PATCH/DELETE`.
- [ ] CRUD produk & customer endpoint minimal happy-path berjalan.
- [ ] Migration + seed dapat dijalankan tanpa error.
- [ ] Test tidak panic untuk skenario auth dasar.

---

## Sprint 2 — Transaksi Penjualan

**Goal:** siklus order dari create sampai paid/closed tersedia via API.

- [ ] Order: list, detail, create, confirm, cancel
- [ ] Payment: create, verify
- [ ] Shipment: create, update status
- [ ] Validasi state transition order (pending -> confirmed -> paid -> shipped -> closed)
- [ ] Atomic transaction saat create order + detail

Acceptance:
- [ ] Tidak ada data partial saat create order multi-item.
- [ ] Cancel menyimpan alasan.
- [ ] Verify payment mengubah status order sesuai rule.

---

## Sprint 3 — Reports + AI + Telegram Settings

**Goal:** dashboard backend siap dengan data agregasi dan insight AI.

- [ ] `GET /reports` (8 tipe report MVP)
- [ ] Integrasi webhook n8n (dashboard insight)
- [ ] `GET/PUT /settings/telegram`
- [ ] Endpoint internal callback AI (`/internal/ai-result`) atau flow sinkron final
- [ ] Fallback response saat n8n timeout/error

Acceptance:
- [ ] Report valid untuk filter tanggal + sales.
- [ ] Response berisi data chart + summary/recommendation.
- [ ] Telegram settings dapat disimpan dan tervalidasi.

---

## Sprint 4 — Frontend Foundation + Public Store + Auth

**Goal:** frontend siap dipakai user untuk login dan melihat katalog.

- [ ] Setup Next.js + data fetching + auth interceptor
- [ ] Halaman login + protected route
- [ ] Halaman `/` dan `/produk/:id`
- [ ] UI base components + loading/error/empty state
- [ ] Integrasi endpoint auth + produk

---

## Sprint 5 — Frontend Admin/Sales + Dashboard AI

**Goal:** operasional harian admin/sales dan dashboard insight selesai.

- [ ] Halaman admin (produk, customer, users, telegram settings)
- [ ] Halaman sales (order list/create/detail + status stepper)
- [ ] Dashboard AI (selector, filter, chart renderer, insight card, anomaly flag)
- [ ] Chat widget SSE
- [ ] Export PDF laporan

---

## Sprint 6 — Security, QA, UAT, Release Candidate

**Goal:** siap rilis internal dengan risiko minimum.

- [ ] Rate limiting + audit authorization
- [ ] Regression test flow inti
- [ ] UAT bersama stakeholder
- [ ] Fix bug severity tinggi
- [ ] Final release checklist + runbook deploy

---

## 5) Board Eksekusi Sprint 1 (Siap Dipakai)

| ID | Task | Status | Dependency | PIC |
|---|---|---|---|---|
| S1-01 | Setup auto-migrate GORM model + tag + schema init | Done | DB ready | BE |
| S1-02 | Terapkan env policy migration (`prod/staging` run, `dev` skip) | Done | S1-01 | BE |
| S1-03 | Uji startup di profile staging (auto-migrate harus jalan) | Todo | S1-02 | BE |
| S1-04 | Seed data minimum (admin/manager/sales/customer) | Todo | S1-03 | BE |
| S1-05 | Implement `POST /auth/logout` | Todo | login existing | BE |
| S1-06 | Implement `GET /auth/me` | Todo | auth middleware | BE |
| S1-07 | Implement `POST /auth/register` | Todo | S1-02 | BE |
| S1-08 | Implement CRUD produk | Todo | S1-02 | BE |
| S1-09 | Implement CRUD customer | Todo | S1-02 | BE |
| S1-10 | Implement CRUD users | Todo | S1-01 | BE |
| S1-11 | Stabilkan auth tests (no panic if DB unavailable) | Done | test strategy | BE/QA |
| S1-12 | Update Swagger untuk endpoint Sprint 1 | Todo | S1-05..S1-10 | BE |

---

## 6) Definition of Ready / Done

### Definition of Ready

- [ ] Scope task jelas.
- [ ] Acceptance criteria jelas.
- [ ] Dependency tercatat.
- [ ] Ada estimasi effort.
- [ ] PIC sudah ditetapkan.

### Definition of Done

- [ ] Code review selesai.
- [ ] Build lulus.
- [ ] Test sesuai scope lulus.
- [ ] API docs diupdate.
- [ ] Tidak ada blocker kritikal terbuka.

---

## 7) Format Tracking Mingguan

| Sprint | Planned | Done | Progress | Status | Blocker Utama |
|---|---:|---:|---:|---|---|
| Sprint 1 | 12 | 0 | 0% | Not Started | - |
| Sprint 2 | 10 | 0 | 0% | Not Started | - |
| Sprint 3 | 8 | 0 | 0% | Not Started | - |
| Sprint 4 | 8 | 0 | 0% | Not Started | - |
| Sprint 5 | 10 | 0 | 0% | Not Started | - |
| Sprint 6 | 8 | 0 | 0% | Not Started | - |

Status:
- `Not Started`
- `On Track`
- `At Risk`
- `Blocked`
- `Done`

---

## 8) Template Review Sprint

```md
## Sprint X Review (Tanggal)

### Progress
- Planned:
- Done:
- Progress:

### Selesai
- ...

### Blocker
- ...

### Carry Over ke Sprint Berikutnya
- ...

### Keputusan
- ...
```

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
