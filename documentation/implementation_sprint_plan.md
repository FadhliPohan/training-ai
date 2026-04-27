# Implementation Sprint Plan — InsightFlow Self-Service AI Dashboard

> Dokumen ini menerjemahkan `task_plan.md` ke rencana sprint eksekusi agar progress tim bisa dilacak harian dan mingguan.
>
> **Versi:** 1.0  
> **Tanggal dibuat:** 24 April 2026  
> **Owner:** Engineering Team (Backend, Frontend, QA, DevOps)  
> **Ritme sprint:** 2 minggu

---

## 1) Tujuan Sprint Plan

- Menetapkan urutan pengerjaan yang realistis dari fondasi sampai release.
- Memastikan setiap sprint punya output terukur (bukan daftar tugas umum).
- Menyediakan format tracking progress untuk daily standup dan sprint review.

---

## 2) Aturan Main Eksekusi

### Definition of Ready (DoR)

Task baru boleh masuk sprint jika:
- Scope jelas dan ada acceptance criteria.
- Dependency sudah diidentifikasi.
- Estimasi sudah disepakati tim.

### Definition of Done (DoD)

Task dianggap selesai jika:
- Kode merge ke branch utama dan lulus review.
- Test terkait lulus (unit/integration sesuai kebutuhan).
- API/fitur terdokumentasi singkat.
- Tidak ada blocker kritikal terbuka.

### Prioritas Eksekusi

1. Backend API siap pakai dulu.
2. Frontend konsumsi API stabil.
3. n8n + AI insight disambungkan setelah data pipeline valid.
4. Security hardening dan UAT sebelum release.

---

## 3) Sprint Timeline (Target)

| Sprint | Tanggal | Fokus Utama | Target Progress Kumulatif |
|---|---|---|---|
| Sprint 1 | 27 Apr 2026 - 08 Mei 2026 | Backend Auth + Master Data inti | 35% |
| Sprint 2 | 11 Mei 2026 - 22 Mei 2026 | Backend Transaksi inti | 55% |
| Sprint 3 | 25 Mei 2026 - 05 Jun 2026 | Backend Reports + n8n fondasi | 70% |
| Sprint 4 | 08 Jun 2026 - 19 Jun 2026 | Frontend fondasi + Auth + Public Store | 80% |
| Sprint 5 | 22 Jun 2026 - 03 Jul 2026 | Frontend Admin/Sales + Dashboard AI | 92% |
| Sprint 6 | 06 Jul 2026 - 17 Jul 2026 | Security, QA, UAT, Release Candidate | 100% |

> Catatan: jika ada scope change besar dari stakeholder, lakukan rebaseline di sprint planning berikutnya.

---

## 4) Rencana Task Per Sprint

## Sprint 1 — Backend Auth + Master Data Inti

**Goal:** sistem login, role-based access, dan data master dasar siap dipakai frontend.

### Deliverables

- [ ] Endpoint `POST /auth/login`, `POST /auth/logout`, `GET /auth/me`
- [ ] JWT flow stabil + middleware `AuthRequired` dan `RoleGuard`
- [ ] CRUD Produk (`GET/POST/PUT/PATCH`) dengan validasi
- [ ] CRUD Customer (`GET/POST/PUT`)
- [ ] Seeder user manager + sales + customer dummy
- [ ] API doc awal untuk endpoint yang sudah jalan

### Acceptance Criteria

- [ ] Admin bisa login dan akses endpoint admin.
- [ ] Role `viewer` hanya read-only.
- [ ] Endpoint master data mengembalikan format response standar.
- [ ] Semua endpoint Sprint 1 teruji minimal happy path + validation case.

### Risiko Utama

- Scope auth melebar (refresh token, session revocation detail) sebelum MVP stabil.

---

## Sprint 2 — Backend Transaksi Inti

**Goal:** siklus order utama berjalan dari create sampai cancel/paid.

### Deliverables

- [ ] Endpoint Order: list, detail, create, confirm, cancel
- [ ] Endpoint Pembayaran: create, verify
- [ ] Transaction safety: create order + order detail atomic
- [ ] Validasi stok dan status order transition
- [ ] Logging transaksi dan error handling yang konsisten

### Acceptance Criteria

- [ ] Sales bisa membuat order multi-item tanpa data partial.
- [ ] Admin bisa verify pembayaran dan status order berubah sesuai rule.
- [ ] Cancel order menyimpan alasan pembatalan.

### Risiko Utama

- Bug state transition (pending/confirmed/paid/cancelled) jika rule belum dipusatkan di service.

---

## Sprint 3 — Backend Reports + n8n Fondasi

**Goal:** data agregasi dashboard tersedia dan n8n siap menerima payload.

### Deliverables

- [ ] Endpoint `GET /reports` dengan parameter `type/from/to/sales_id`
- [ ] Minimal 4 laporan prioritas aktif: `daily-sales`, `monthly-sales`, `top-products`, `sales-by-person`
- [ ] Integrasi webhook ke n8n untuk AI summary
- [ ] Endpoint internal callback hasil AI (`/internal/ai-result`) atau pola sinkron final
- [ ] Konfigurasi Telegram settings endpoint (`GET/PUT /settings/telegram`)

### Acceptance Criteria

- [ ] Report query valid untuk range tanggal dan filter sales.
- [ ] Respons dashboard memuat data chart + summary AI (fallback aman jika AI gagal).
- [ ] Konfigurasi Telegram tersimpan dan tervalidasi.

### Risiko Utama

- Latensi n8n/AI tinggi; perlu timeout dan fallback response.

---

## Sprint 4 — Frontend Fondasi + Auth + Public Store

**Goal:** aplikasi frontend usable dari sisi login dan halaman toko publik.

### Deliverables

- [ ] Setup Next.js 14, state/data fetching, API client
- [ ] Halaman login + protected route
- [ ] Halaman public `/` dan `/produk/:id`
- [ ] Komponen dasar UI (button, input, modal, table, alert, skeleton)
- [ ] Error/loading/empty states di halaman utama

### Acceptance Criteria

- [ ] User bisa login lalu mengakses halaman sesuai role.
- [ ] Halaman publik responsif desktop/mobile.
- [ ] Request error tampil jelas dalam Bahasa Indonesia.

### Risiko Utama

- Ketergantungan API belum stabil, menyebabkan banyak mocking ulang frontend.

---

## Sprint 5 — Frontend Admin/Sales + Dashboard AI

**Goal:** alur operasional harian admin/sales dan dashboard insight siap demo.

### Deliverables

- [ ] Halaman admin: produk, customer, users, settings telegram
- [ ] Halaman sales: daftar order, buat order, detail order + stepper status
- [ ] Dashboard: report selector, filter bar, chart renderer, insight card, anomaly banner
- [ ] Integrasi AI chat widget (SSE) di public/frontend utama
- [ ] Download PDF untuk dashboard (versi MVP)

### Acceptance Criteria

- [ ] Admin dapat kelola master data end-to-end.
- [ ] Sales dapat menjalankan flow order tanpa pindah tool.
- [ ] Dashboard menampilkan insight AI dan tetap usable saat AI timeout.

### Risiko Utama

- Kompleksitas UI dashboard tinggi; rawan penundaan jika tidak diprioritaskan komponen reusable.

---

## Sprint 6 — Security, QA, UAT, Release Candidate

**Goal:** sistem siap rilis internal dengan risiko produksi minimum.

### Deliverables

- [ ] Security hardening: rate limit, audit role access, input validation review
- [ ] Regression test untuk flow utama (auth, order, payment, report)
- [ ] UAT checklist bersama stakeholder
- [ ] Perbaikan bug prioritas tinggi
- [ ] Release candidate + dokumentasi deploy/runbook ringkas

### Acceptance Criteria

- [ ] Tidak ada bug severity kritikal saat UAT selesai.
- [ ] Endpoint inti memiliki test coverage fungsional yang memadai.
- [ ] Semua isu blocker sprint ditutup atau punya workaround yang disepakati.

### Risiko Utama

- Temuan UAT besar di akhir sprint; mitigasi dengan demo incremental sejak Sprint 3.

---

## 5) Sprint Progress Tracker (Isi Berkala)

Gunakan tabel ini saat standup/review. Update minimal 2x seminggu.

| Sprint | Planned Task | Done Task | Progress | Status | Blocker | PIC |
|---|---:|---:|---:|---|---|---|
| Sprint 1 | 18 | 0 | 0% | Not Started | - | BE Lead |
| Sprint 2 | 14 | 0 | 0% | Not Started | - | BE Lead |
| Sprint 3 | 12 | 0 | 0% | Not Started | - | BE + n8n |
| Sprint 4 | 16 | 0 | 0% | Not Started | - | FE Lead |
| Sprint 5 | 20 | 0 | 0% | Not Started | - | FE + BE |
| Sprint 6 | 12 | 0 | 0% | Not Started | - | QA + BE + FE |

Status yang dipakai:
- `Not Started`
- `On Track`
- `At Risk`
- `Blocked`
- `Done`

---

## 6) Template Update Mingguan

Salin template ini ke catatan sprint review:

```md
## Sprint X Review (Tanggal)

### Progress
- Planned: xx task
- Done: xx task
- Progress: xx%

### Done This Sprint
- ...

### Blockers
- ...

### Carry Over ke Sprint Berikutnya
- ...

### Keputusan Penting
- ...
```

---

## 7) Catatan Eksekusi Penting

- Jaga WIP kecil: fokus selesai per modul, jangan buka terlalu banyak task paralel.
- Kunci kualitas di API contract lebih awal agar frontend tidak sering rework.
- Demo internal di tengah sprint (bukan hanya akhir sprint) untuk deteksi scope drift.
