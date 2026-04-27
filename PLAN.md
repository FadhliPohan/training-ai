# InsightFlow — Execution Plan

> **Project:** InsightFlow Self-Service AI Dashboard Penjualan Pakaian
> **Stack:** Go Fiber · Next.js 14 · PostgreSQL · n8n · Telegram Bot API
> **Start:** 27 April 2026 · **Target MVP:** 17 Juli 2026

---

## Konsep Sistem (4 Fitur Utama)

```
WEB (2 Fitur)
├── 1. Dashboard Laporan   → User pilih laporan → chart + AI insight
└── 2. Admin Panel         → CRUD produk, customer, users, settings

TELEGRAM (2 Fitur)
├── 3. Daily Summary       → n8n scheduler 07:00 → hit endpoint → AI → Telegram
└── 4. On-demand Q&A       → User kirim pesan → n8n webhook → intent parsing
                              → hit endpoint yang sesuai → AI normalize → balas
```

---

## Arsitektur n8n-Centric

```
┌─────────────────────────────────────────────────────────────────┐
│                        n8n (Orchestrator)                        │
│                                                                  │
│  Fitur 3: Scheduler 07:00 ──→ GET /reports ──→ AI ──→ Telegram  │
│                                                                  │
│  Fitur 4: Telegram Webhook ──→ Parse Intent                      │
│                                  ├── GET /reports?type=...       │
│                                  ├── GET /produk                 │
│                                  ├── GET /orders/summary         │
│                                  └── GET /users/summary          │
│                               ──→ AI (normalize/format) ──→ Telegram│
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                      Golang API (Data Layer)                     │
│  - Menyediakan endpoint data yang bersih                         │
│  - Auth + RBAC untuk web                                         │
│  - Endpoint publik internal untuk n8n (no auth / internal key)   │
│  - Fitur 1: GET /reports + panggil n8n AI workflow               │
│  - Fitur 2: CRUD admin (produk, customer, users, settings)       │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                     Next.js (Web Frontend)                       │
│  - Fitur 1: Dashboard — pilih laporan, tampil chart + insight    │
│  - Fitur 2: Admin Panel — CRUD + settings Telegram               │
│  - Login / protected routes                                      │
└─────────────────────────────────────────────────────────────────┘
```

---

## Evaluasi Arsitektur Lama vs Baru

| Aspek                  | Arsitektur Lama                     | Arsitektur Baru                                  |
| ---------------------- | ----------------------------------- | ------------------------------------------------ |
| Telegram Daily Summary | n8n query DB langsung               | n8n hit `GET /reports` (cleaner, reuse logic)    |
| Telegram Q&A           | n8n query DB langsung per role      | n8n parse intent → hit endpoint yang sesuai → AI |
| Dashboard AI           | Golang panggil n8n, tunggu response | Tetap: Golang → n8n → LLM → balik ke Golang      |
| Anomaly                | n8n polling DB langsung             | n8n hit `GET /reports/anomaly` → AI → Telegram   |
| Role filter            | Di n8n (query SQL per role)         | Di Golang endpoint (parameter `?sales_id=`)      |

### Bottleneck & Risiko

| #   | Potensi Masalah                                 | Mitigasi                                                                 |
| --- | ----------------------------------------------- | ------------------------------------------------------------------------ |
| 1   | n8n → Golang → n8n → LLM: latency chain panjang | Timeout 15 detik di n8n; fallback message jika LLM lambat                |
| 2   | LLM gagal / rate limit                          | Error Trigger node + fallback teks statis di setiap workflow             |
| 3   | Telegram webhook flood (spam pesan)             | Rate limit per `telegram_user_id` di n8n (IF node cek last request time) |
| 4   | Intent parsing salah → endpoint salah           | Prompt NLP yang ketat + whitelist intent di Code node                    |
| 5   | Golang endpoint tidak authenticated dari n8n    | Gunakan internal API key header (`X-Internal-Key`)                       |

---

## Current Status (27 Apr 2026)

| Area                   | Status     | Keterangan                                                |
| ---------------------- | ---------- | --------------------------------------------------------- |
| Backend Foundation     | ✅ Selesai | Auth, CRUD, Reports GET, Settings                         |
| Telegram Webhook       | ✅ Selesai | `POST /api/v1/telegram/webhook` + forward ke n8n          |
| Internal Middleware    | ✅ Selesai | `InternalKeyGuard` via `X-Internal-Key`                   |
| n8n Dashboard Workflow | ✅ Aktif   | Golang → n8n → LLM berjalan                               |
| n8n Telegram Daily     | 🟡 Parsial | Workflow belum end-to-end ke Telegram                     |
| n8n Telegram Q&A       | 🔴 Belum   | n8n workflow telegram-qa (konsumsi payload dari webhook)  |
| Frontend Dashboard     | 🔴 Belum   | Sprint 4                                                  |
| Frontend Admin         | 🔴 Belum   | Sprint 5                                                  |
| Security               | 🔴 Belum   | Sprint 6                                                  |

### Endpoint Aktif Saat Ini

```
GET  /health
GET  /swagger/*
POST /api/v1/auth/login
POST /api/v1/auth/logout
POST /api/v1/auth/register
GET  /api/v1/auth/me
GET  /api/v1/produk
GET  /api/v1/produk/:id
POST /api/v1/produk          (admin)
PUT  /api/v1/produk/:id      (admin)
PATCH /api/v1/produk/:id     (admin)
GET  /api/v1/customer
POST /api/v1/customer
PUT  /api/v1/customer/:id
GET  /api/v1/users
POST /api/v1/users
PUT  /api/v1/users/:id
PATCH /api/v1/users/:id
GET  /api/v1/reports
GET  /api/v1/settings/telegram (admin)
PUT  /api/v1/settings/telegram (admin)
```

---

## FASE 1 — Backend: Internal Endpoints untuk n8n

> **Goal:** Semua endpoint yang dibutuhkan n8n tersedia dan bersih.
> **Status:** ✅ Selesai

### Chunk 1.1 — Endpoint Reports (n8n-ready)

Endpoint ini dipanggil oleh n8n (bukan hanya frontend). Harus ringan dan fast.

- [x] `GET /api/v1/reports?type=&from=&to=&sales_id=` — 8 tipe laporan
- [x] Whitelist `type` parameter (tolak type tidak dikenal)
- [x] Response: `{ data[], meta: { total, from, to } }`
- [x] `GET /api/internal/reports/summary` — endpoint khusus n8n daily summary
  - Auth: `X-Internal-Key` header
  - Return: omzet, order, top produk, stok rendah, anomali flag
- [x] `GET /api/internal/reports/anomaly` — untuk workflow anomaly check
  - Parameter: `?threshold=10` (default 10%)
  - Return: `has_anomaly`, list anomali dengan `variance_pct`, `severity`, `direction`

**Selesai jika:** n8n bisa hit kedua endpoint internal dan mendapat data terstruktur. ✅

---

### Chunk 1.2 — Endpoint Dynamic untuk Telegram Q&A

n8n akan hit endpoint berbeda berdasarkan intent user. Semua endpoint ini harus support filter via query param.

| Intent User          | Endpoint yang Di-hit n8n                                     | Contoh                  |
| -------------------- | ------------------------------------------------------------ | ----------------------- |
| "penjualan hari ini" | `GET /api/internal/reports/summary?from=today`               | Omzet + order hari ini  |
| "produk terlaris"    | `GET /api/internal/reports?type=top-products`                | Top 5 produk            |
| "stok rendah"        | `GET /api/internal/reports?type=low-stock`                   | Produk stok < threshold |
| "order pending"      | `GET /api/internal/reports?type=order-funnel&status=pending` | Jumlah order pending    |
| "performa sales X"   | `GET /api/internal/reports?type=sales-by-person&sales_id=X`  | Data sales tertentu     |

- [x] `GET /api/internal/reports?type=top-products` — Top produk
- [x] `GET /api/internal/reports?type=low-stock` — Stok rendah
- [x] `GET /api/internal/reports?type=order-funnel` — Funnel status order
- [x] `GET /api/internal/reports?type=sales-by-person&sales_id=X` — Per sales (scoped)
- [x] `GET /api/internal/users/by-telegram?telegram_user_id=X` — Resolve user role
- [x] Auth via `X-Internal-Key` header
- [x] Response `mode=raw` — array of objects, label jelas, siap dikirim ke LLM

**Selesai jika:** n8n bisa hit semua endpoint di atas dan mendapat data yang siap dikirim ke LLM. ✅

---

### Chunk 1.3 — Stabilisasi & Swagger

- [x] Auto-migrate GORM + seed data dummy
- [x] CRUD produk, customer, users lengkap
- [ ] Validasi soft-delete produk (`aktif = false`)
- [ ] `telegram_user_id` pada user dapat diisi via `PUT /users/:id`
- [ ] Swagger dokumentasi semua endpoint (termasuk `/internal/*`)
- [ ] Logging structured untuk trace latency `/reports → n8n`

---

## FASE 2 — n8n: 2 Workflow Telegram

> **Goal:** Kedua fitur Telegram berjalan end-to-end dan stabil.
> **Status:** 🟡 Parsial (paralel dengan Sprint 1 & 2)

### Chunk 2.1 — Setup n8n

- [ ] Deploy n8n Docker Compose (`restart: always`)
- [ ] Hanya bisa diakses via internal network
- [ ] Credentials di n8n:
  - [ ] OpenAI / Gemini API Key
  - [ ] Telegram Bot Token
  - [ ] `BACKEND_URL` = `http://localhost:8080`
  - [ ] `INTERNAL_API_KEY` = sama dengan `.env` backend

---

### Chunk 2.2 — Fitur 3: Daily Summary (Scheduler)

**Alur:**

```
n8n Scheduler 07:00 WIB
  └─→ GET /api/internal/reports/summary
        └─→ Code node: format data + susun prompt
              └─→ LLM: "Buat ringkasan penjualan harian yang ramah"
                    └─→ Code node: format pesan Telegram
                          └─→ sendMessage ke semua chat_id aktif
```

**Tasks:**

- [ ] Schedule Trigger — cron `0 0 * * *` (UTC = 07:00 WIB)
- [ ] HTTP Request node — `GET /api/internal/reports/summary` dengan `X-Internal-Key`
- [ ] Code node — format payload menjadi prompt:

  ```
  Data penjualan hari ini:
  - Total omzet: [X]
  - Total order: [Y] (selesai: A, pending: B)
  - Produk terlaris: [nama]
  [data anomali jika ada]

  Buat ringkasan singkat (3-4 kalimat) dalam Bahasa Indonesia
  yang mudah dibaca oleh manajer penjualan.
  ```

- [ ] OpenAI/Gemini node — generate ringkasan
- [ ] Code node — format output:

  ```
  📊 *Laporan Penjualan — [TANGGAL]*

  💰 Omzet: Rp X.XXX.XXX
  📦 Order: XX total (✅ XX selesai | ⏳ XX pending)
  🏆 Terlaris: [Nama Produk]

  [Ringkasan AI]

  _Dikirim otomatis pukul 07.00 WIB_
  ```

- [ ] Telegram node — `sendMessage` ke semua `chat_id` aktif
- [ ] Error Trigger node — jika gagal, kirim pesan fallback:
      `"⚠️ Gagal mengambil laporan pagi ini. Silakan cek dashboard manual."`
- [ ] Retry: 2x dengan delay 30 detik

**Selesai jika:** pesan terkirim otomatis 07:00 WIB setiap hari tanpa intervensi manual.

---

### Chunk 2.3 — Fitur 4: On-Demand Q&A via Telegram

**Alur:**

```
User kirim pesan di Telegram
  └─→ Telegram Webhook → n8n
        └─→ Verifikasi telegram_user_id (lookup via endpoint atau DB)
              └─→ LLM: Parse intent + tentukan endpoint + extract params
                    └─→ HTTP Request: hit endpoint yang sesuai
                          └─→ LLM: Normalize & format data menjadi respons ramah
                                └─→ sendMessage kembali ke user
```

**Tasks:**

- [ ] Telegram Trigger node — aktifkan webhook Telegram
- [ ] HTTP Request node — `GET /api/internal/users/by-telegram?telegram_user_id=X`
  - Jika tidak ditemukan: langsung balas "Maaf, akun Telegram Anda belum terdaftar."
- [ ] LLM node (Intent Parsing) — parse pesan user:

  ```
  Kamu adalah intent parser untuk sistem laporan penjualan pakaian.
  Pesan user: "[PESAN]"
  Role user: "[ROLE]" (sales atau manager)

  Tentukan:
  1. intent: salah satu dari [daily_summary, top_products, low_stock,
     order_pending, sales_performance, unknown]
  2. params: { from, to, sales_id (jika sales: otomatis isi dengan user_id) }

  Kembalikan JSON. Jika unknown, kembalikan intent: "unknown".
  ```

- [ ] Switch node — route berdasarkan `intent`:
  - `daily_summary` → `GET /api/internal/reports/summary?...`
  - `top_products` → `GET /api/internal/reports?type=top-products&...`
  - `low_stock` → `GET /api/internal/reports?type=low-stock`
  - `order_pending` → `GET /api/internal/reports?type=order-funnel&status=pending`
  - `sales_performance` → `GET /api/internal/reports?type=sales-by-person&sales_id=...`
  - `unknown` → balas "Maaf, saya tidak mengerti pertanyaan Anda. Coba: 'penjualan hari ini', 'produk terlaris', 'stok rendah'."
- [ ] HTTP Request node — hit endpoint sesuai intent (dengan `X-Internal-Key`)
- [ ] LLM node (Response Formatter) — normalize dan format data:

  ```
  Data laporan: [DATA_JSON]
  Pertanyaan asal user: "[PESAN]"

  Buat jawaban singkat (max 5 baris) dalam Bahasa Indonesia yang
  langsung menjawab pertanyaan. Gunakan emoji relevan. Format Telegram markdown.
  ```

- [ ] Telegram node — `sendMessage` ke `chat_id` pengirim
- [ ] Error handling per node:
  - Endpoint gagal → "Maaf, data sedang tidak bisa diambil. Coba lagi nanti."
  - LLM timeout → "Maaf, AI sedang sibuk. Data mentah: [raw data singkat]"

**Catatan penting (Role Scoping):**

- Role **sales**: `sales_id` di-inject otomatis dari profil user (tidak bisa query data sales lain)
- Role **manager**: bisa query semua data, `sales_id` opsional
- Enforcement dilakukan di endpoint backend (bukan di n8n)

**Selesai jika:** user Telegram bisa kirim pesan natural language dan mendapat jawaban relevan ≤ 15 detik.

---

## FASE 3 — n8n: Workflow Dashboard AI (Web Fitur 1)

> **Goal:** `GET /reports` dari web → n8n → LLM → response dengan chart + insight.
> **Status:** ✅ Aktif (Chunk 2.2 lama — tinggal hardening)

**Alur:**

```
Frontend → GET /reports → Golang aggregasi SQL
  └─→ POST /n8n/webhook/dashboard-ai (data + report_type)
        └─→ LLM: tentukan chart_type + summary + anomali + rekomendasi
              └─→ Golang gabungkan → return ke frontend
```

**Tasks:**

- [x] HTTP Webhook node — terima data dari Golang
- [x] Code node — format prompt
- [x] LLM node — generate analysis
- [x] Code node — parse JSON response
- [x] Respond to Webhook
- [ ] Error Trigger node — fallback jika LLM gagal:
  ```json
  {
    "chart_type": "bar",
    "summary": "Data tersedia namun analisis AI sedang tidak tersedia.",
    "anomalies": [],
    "recommendation": "Silakan refresh atau coba beberapa saat lagi."
  }
  ```
- [ ] Timeout di n8n: 10 detik (Golang timeout 12 detik)

**System Prompt:**

```
Kamu adalah analis data penjualan pakaian.
Tipe laporan: [REPORT_TYPE]
Data: [DATA_JSON]

Tentukan:
1. chart_type: "line" | "bar" | "pie" | "funnel"
2. summary: ringkasan 2-3 kalimat Bahasa Indonesia
3. anomalies: array [{metric, actual, expected, delta_pct}] jika ada varians > [THRESHOLD]%
4. recommendation: 1 kalimat tindakan konkret

Kembalikan JSON valid saja, tanpa teks lain.
```

---

## FASE 4 — Backend Transactions (Backlog)

> **Status:** 🔴 Sprint 2 (11–22 Mei 2026)
> CRUD transaksi untuk operasional toko — tidak blocking untuk n8n/Telegram.

### Chunk 4.1 — Orders

- [ ] `POST /orders` — validasi stok, atomic transaction
- [ ] `GET /orders` + `GET /orders/:id`
- [ ] `POST /orders/:id/confirm` → `confirmed`
- [ ] `POST /orders/:id/cancel` → `cancelled` + alasan

### Chunk 4.2 — Payments & Shipments

- [ ] `POST /payments` + `POST /payments/:id/verify`
- [ ] `POST /shipments` + `PUT /shipments/:id`

---

## FASE 5 — Frontend Web (2 Fitur)

> **Status:** 🔴 Sprint 4–5 (Jun 2026)
> **Dimulai setelah:** Fase 1 + Fase 2 + Fase 3 selesai

### Chunk 5.1 — Setup Next.js

- [ ] `npx create-next-app@latest ./` (TypeScript + Tailwind + App Router)
- [ ] Axios instance + JWT interceptor (redirect 401 → `/login`)
- [ ] React Query provider
- [ ] Design tokens: Primary `#2563eb`, dark surface `#0f172a`, font Inter

### Chunk 5.2 — Auth Minimal

- [ ] `/login` — form + validasi + error Bahasa Indonesia
- [ ] Protected route HOC
- [ ] Logout

### Chunk 5.3 — Fitur 1: Dashboard

- [ ] Layout: sidebar collapsible + top navbar
- [ ] `ReportSelector` — dropdown 8 tipe laporan
- [ ] `FilterBar` — date range, filter sales
- [ ] Tombol "Tampilkan" → `GET /api/v1/reports`
- [ ] `ChartRenderer` — render `line`/`bar`/`pie`/`funnel` sesuai `chart_type`
- [ ] `AIInsightCard` — teks ringkasan AI
- [ ] `AnomalyFlag` — banner ⚠️ + rekomendasi
- [ ] Skeleton loader + error state + retry

### Chunk 5.4 — Fitur 2: Admin Panel

- [ ] `/admin/produk` — DataTable + CRUD (form: nama, kategori, ukuran, warna, bahan, harga, stok)
- [ ] `/admin/customer` — DataTable + CRUD
- [ ] `/admin/users` — DataTable + set role + `telegram_user_id`
- [ ] `/settings/telegram` — form: `chat_id`, `jam_summary`, `threshold`

### Chunk 5.5 — Fitur Web Lanjutan (Backlog)

- [ ] Halaman toko publik (`/`, `/produk/:id`)
- [ ] AI Chat widget SSE
- [ ] Export PDF laporan

---

## FASE 6 — Security + UAT + Go-Live

> **Status:** 🔴 Sprint 6 (06–17 Jul 2026)

### Chunk 6.1 — Security

- [ ] JWT HS256, secret ≥ 32 karakter, tidak di repo
- [ ] Rate limit login: 5x/menit per IP
- [ ] `X-Internal-Key` untuk semua endpoint `/api/internal/*`
- [ ] RoleGuard: sales hanya akses data sendiri (`WHERE sales_id = user_id`)
- [ ] Telegram: reject pesan dari `telegram_user_id` tidak terdaftar
- [ ] Payload ke LLM: hanya data aggregat, tidak ada PII
- [ ] HTTPS via Nginx, security headers, CORS whitelist
- [ ] n8n tidak expose ke publik

### Chunk 6.2 — UAT & Go-Live

- [ ] Regression test: login → laporan → chart → Telegram summary terkirim → Q&A Telegram berjalan
- [ ] UAT minimal 3 pengguna non-teknis
- [ ] Fix bug High/Critical
- [ ] Docker Compose production + Nginx verified
- [ ] Runbook deploy + rollback plan

---

## Sprint Timeline

| Sprint   | Periode       | Fokus                                         | Status         |
| -------- | ------------- | --------------------------------------------- | -------------- |
| Sprint 1 | 27 Apr–08 Mei | Backend internal endpoints + n8n setup        | 🟡 In Progress |
| Sprint 2 | 11–22 Mei     | n8n Telegram workflows + Backend transactions | 🔴 Planned     |
| Sprint 3 | 25 Mei–05 Jun | Hardening n8n + Dashboard AI workflow         | 🔴 Planned     |
| Sprint 4 | 08–19 Jun     | Frontend: Auth + Dashboard + Settings         | 🔴 Planned     |
| Sprint 5 | 22 Jun–03 Jul | Frontend: Admin Panel + polish                | 🔴 Planned     |
| Sprint 6 | 06–17 Jul     | Security + UAT + Go-Live                      | 🔴 Planned     |

> **Prioritas:** Telegram (Fitur 3 & 4) dikerjakan **sebelum** frontend.
> n8n workflows bisa paralel dengan backend transactions.

---

## Quick Dev Commands

```bash
# Start services
cd /home/sandi/PUSRI/training-ai
docker compose up -d

# Backend
cd be-penjualan
make run          # development
make staging      # staging (auto-migrate)

# Test
curl http://localhost:8080/health
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@insightflow.id","password":"Admin@12345"}'

# n8n UI
open http://localhost:5678

# Swagger
open http://localhost:8080/swagger/index.html

# DB
PGPASSWORD=insightflow123 psql -h localhost -p 5433 -U insightflow -d insightflow_db
```

---

## API Contract

**Response sukses:**

```json
{ "success": true, "message": "...", "data": {} }
```

**Response error:**

```json
{
  "success": false,
  "message": "Pesan error ramah",
  "errors": { "field": "detail" }
}
```

| Code | Artinya          |
| ---- | ---------------- |
| 200  | OK               |
| 201  | Created          |
| 400  | Validation error |
| 401  | Unauthenticated  |
| 403  | Forbidden        |
| 404  | Not Found        |
| 500  | Server error     |

---

## Definition of Done

Chunk selesai jika:

- [ ] Acceptance criteria terpenuhi
- [ ] Build lulus
- [ ] Swagger/docs diupdate (backend)
- [ ] Tidak ada bug critical terbuka
- [ ] Code sudah di-commit

---

_Last updated: 27 April 2026 — Konsep: Web 2 fitur + Telegram 2 fitur, n8n sebagai orchestrator pusat._
