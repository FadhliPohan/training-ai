# Area 5 — Security

> **Tujuan:** Membangun sistem yang aman secara berlapis — dari autentikasi, otorisasi, validasi input, hingga keamanan data dan infrastruktur.

---

## 5.1 Auth & Session

| Aspek | Spesifikasi |
|---|---|
| Algoritma JWT | HS256 |
| Secret | Minimal 32 karakter, dari `.env` (tidak di-hardcode) |
| Token expiry | 8 jam |
| Penyimpanan token | `httpOnly cookie` atau memory — **TIDAK** di `localStorage` |
| Logout | Invalidasi token (blacklist Redis atau gunakan short-lived token) |
| Rate limit login | Maksimal 5 percobaan/menit per IP |

- [ ] JWT HS256, secret ≥ 32 karakter dari env (tidak di-hardcode)
- [ ] Token expiry 8 jam
- [ ] Token tidak disimpan di `localStorage` — gunakan `httpOnly cookie` atau memory
- [ ] Logout invalidasi token (blacklist Redis atau short-lived token)
- [ ] Rate limit login: 5 percobaan/menit per IP

> [!CAUTION]
> Menyimpan JWT di `localStorage` sangat rentan terhadap serangan XSS. Gunakan `httpOnly cookie` agar token tidak bisa diakses oleh JavaScript di browser.

---

## 5.2 Authorization

| Rule | Detail |
|---|---|
| Semua endpoint | Wajib melewati middleware `AuthRequired` |
| Endpoint sensitif | Wajib melewati middleware `RoleGuard` |
| Data sales | Query wajib filter `WHERE sales_id = user_id` (sales hanya lihat datanya sendiri) |
| Telegram Bot | Hanya merespons `telegram_user_id` yang terdaftar di `app.users` |

- [ ] Semua endpoint: `AuthRequired` middleware
- [ ] Endpoint sensitif: `RoleGuard`
- [ ] Sales hanya bisa akses order miliknya (`WHERE sales_id = user_id`)
- [ ] Telegram bot: hanya respons `telegram_user_id` yang terdaftar di `app.users`

> [!IMPORTANT]
> Authorization **tidak boleh** hanya mengandalkan tampilan UI (misalnya sembunyikan tombol untuk role tertentu). Semua pengecekan role harus ada di **backend/middleware**.

---

## 5.3 Input Validation

- [ ] Validasi di backend (bukan hanya frontend) menggunakan `go-playground/validator`
- [ ] Sanitasi field text panjang (alamat, catatan) — cegah XSS
- [ ] Whitelist nilai `type` di endpoint `/reports` (hanya nilai yang diizinkan yang diterima)
- [ ] File upload (foto produk): validasi MIME type + max ukuran **2MB**

> [!NOTE]
> Validasi frontend hanya untuk UX (feedback cepat). Validasi backend adalah **garis pertahanan utama** dan tidak bisa dilewati.

---

## 5.4 Data Protection

| Aspek | Spesifikasi |
|---|---|
| Password hashing | bcrypt dengan cost factor ≥ 12 |
| Data dikirim ke LLM | **Hanya data aggregat** — tidak ada data personal customer |
| Log aplikasi | Tidak boleh mencatat password, token, atau PII (Personally Identifiable Information) |
| Privilege DB user | Minimal — hanya `SELECT`, `INSERT`, `UPDATE` (tidak ada `DROP`, `CREATE`, `TRUNCATE`) |

- [ ] Password: bcrypt cost factor ≥ 12
- [ ] Kirim ke LLM: hanya data aggregat — **tidak ada data personal customer**
- [ ] Log: tidak mencatat password, token, atau PII
- [ ] DB user: privilege minimal (`SELECT`, `INSERT`, `UPDATE` — tidak `DROP`)

> [!CAUTION]
> Sebelum go-live, **review payload** yang dikirim ke n8n/LLM untuk memastikan tidak ada data personal customer (nama, email, telepon, alamat) yang ter-expose ke layanan eksternal.

---

## 5.5 Transport & Infrastructure

| Aspek | Spesifikasi |
|---|---|
| HTTPS | Wajib — TLS 1.2+ via Nginx |
| Security headers | `X-Frame-Options`, `X-Content-Type-Options`, `Strict-Transport-Security` |
| CORS | Whitelist domain frontend saja |
| n8n | Tidak boleh terekspos ke publik — hanya akses internal network |
| Secrets | Semua di `.env`, tidak boleh masuk ke repository Git |

- [ ] HTTPS wajib — TLS 1.2+ via Nginx
- [ ] Security headers: `X-Frame-Options`, `X-Content-Type-Options`, `Strict-Transport-Security`
- [ ] CORS: whitelist domain frontend saja
- [ ] n8n: tidak terekspos ke publik — hanya internal network
- [ ] Semua secret di `.env`, tidak di repository

> [!WARNING]
> Pastikan file `.env` masuk ke `.gitignore` sebelum commit pertama. Jika secret sudah ter-expose ke repository, semua secret harus dirotasi segera.

---

## 5.6 Monitoring & Alerting

- [ ] Log semua akses endpoint sensitif: login, CRUD user, perubahan konfigurasi Telegram
- [ ] Alert jika terjadi spike 401/403 dalam waktu singkat (indikasi brute force)
- [ ] Health check endpoint `GET /health` untuk monitoring uptime

---

## Security Checklist (Pre Go-Live)

| # | Checklist | Status |
|---|---|---|
| 1 | Semua secret ada di `.env` dan tidak di repo | ☐ |
| 2 | JWT menggunakan secret ≥ 32 karakter | ☐ |
| 3 | Rate limit login aktif | ☐ |
| 4 | Semua endpoint melewati `AuthRequired` | ☐ |
| 5 | Role guard aktif di endpoint sensitif | ☐ |
| 6 | Payload ke LLM tidak mengandung data personal | ☐ |
| 7 | bcrypt cost factor ≥ 12 | ☐ |
| 8 | HTTPS aktif dan redirect HTTP → HTTPS | ☐ |
| 9 | n8n hanya bisa diakses dari internal network | ☐ |
| 10 | CORS hanya whitelist domain frontend | ☐ |

> [!TIP]
> Jadwalkan **security review session** setelah semua fitur core selesai dan sebelum UAT dimulai. Ini adalah momen terbaik untuk menjalankan seluruh checklist di atas.
