# QA Report — InsightFlow Frontend
**Date:** 27 April 2026  
**Tester:** Senior QA  
**Backend:** `http://localhost:3032` (running)  
**Frontend:** `http://localhost:3000` (Next.js 14)  
**Scope:** All admin pages + settings + auth

---

## Summary

| Severity | Count |
|----------|-------|
| 🔴 Critical | 4 |
| 🟡 Medium   | 3 |
| 🟢 Low      | 2 |
| **Total**   | **9** |

---

## Bug List

---

### BUG-001 — 🔴 Critical
**Page:** `/admin/produk`  
**Feature:** Tampil data & Edit produk  
**Description:** Kolom "Kategori" di tabel selalu tampil kosong (`—`). Field yang digunakan di frontend adalah `row.kategori`, sedangkan API mengembalikan `kategori_pakaian`.  
**Root Cause:** Field name mismatch — FE uses `kategori`, API returns `kategori_pakaian`.  
**Evidence:**
```
API response: { "kategori_pakaian": "Kemeja", ... }
FE column key: "kategori"          ← WRONG
FE openEdit:   row.kategori        ← WRONG (always undefined → form.kategori = "")
FE payload:    { kategori: "..." } ← WRONG (backend ignores unknown field)
```
**Impact:** Kategori tidak tampil di tabel. Saat edit, field kategori kosong. Saat simpan, `kategori_pakaian` tidak terkirim → backend menyimpan string kosong.  
**Fix Required:** Ganti semua `kategori` → `kategori_pakaian` di `admin/produk/page.js`. Update `KATEGORI_OPTIONS` agar nilainya sesuai enum backend (`atasan`, `bawahan`, `dress`, `outerwear`, `aksesoris`). Update `searchKeys` dari `"kategori"` → `"kategori_pakaian"`.

---

### BUG-002 — 🔴 Critical
**Page:** `/admin/produk`  
**Feature:** Create & Edit produk  
**Description:** Payload create/update tidak menyertakan `kode_produk`. Backend handler mewajibkan field ini (dipakai sebagai unique key). Tanpa `kode_produk`, create akan gagal atau menyimpan string kosong yang menyebabkan constraint error.  
**Root Cause:** `EMPTY_FORM` tidak memiliki field `kode_produk`. Form tidak memiliki input untuk `kode_produk`.  
**Evidence:**
```js
const EMPTY_FORM = {
  nama: "", kategori: "", ukuran: "", warna: "",
  bahan: "", harga: "", stok: "", aktif: true
  // kode_produk MISSING
};
// Backend CreateRequest requires: kode_produk, nama, kategori_pakaian, ...
```
**Impact:** POST `/api/v1/produk` akan mengembalikan error atau conflict karena `kode_produk` kosong.  
**Fix Required:** Tambah field `kode_produk` ke `EMPTY_FORM` dan tambah input di form modal.

---

### BUG-003 — 🔴 Critical
**Page:** `/admin/customer`  
**Feature:** Tampil data & Edit customer  
**Description:** Kolom "Kota" di tabel selalu tampil `—`. Field `kota` tidak ada di API response. API mengembalikan `alamat` (full address string), tidak ada field `kota` terpisah. Form juga memiliki field `kota` yang tidak pernah terisi dari API dan tidak dikirim ke backend.  
**Root Cause:** Field `kota` tidak ada di domain `Customer`. API response fields: `id, kode_cust, nama, email, telepon, alamat, aktif, created_at`.  
**Evidence:**
```
API response keys: id, kode_cust, nama, email, telepon, alamat, aktif, created_at
FE column key: "kota"        ← field tidak ada di API
FE form field: kota          ← tidak dikirim ke backend (ignored)
FE searchKeys: ["nama","email","kota","telepon"] ← "kota" tidak berguna
```
**Impact:** Kolom Kota selalu kosong. Data kota yang diisi user tidak tersimpan.  
**Fix Required:** Hapus field `kota` dari form dan tabel. Ganti kolom "Kota" dengan kolom "Alamat" menggunakan `row.alamat`. Tambah `kode_cust` ke form (required oleh backend). Update `searchKeys`.

---

### BUG-004 — 🔴 Critical
**Page:** `/admin/users`  
**Feature:** Tampil data & Edit user  
**Description:** Kolom "Nama" di tabel selalu tampil kosong. Field yang digunakan adalah `row.name`, sedangkan API mengembalikan `nama`. Form juga menggunakan `form.name` saat create/edit, sehingga field `nama` tidak pernah terkirim ke backend.  
**Root Cause:** Field name mismatch — FE uses `name`, API returns `nama`.  
**Evidence:**
```
API response: { "nama": "Administrator", ... }
FE column key: "name"          ← WRONG
FE EMPTY_FORM: { name: "" }    ← WRONG
FE openEdit:   row.name        ← always undefined
FE payload:    { name: "..." } ← backend ignores, expects "nama"
```
**Impact:** Nama user tidak tampil di tabel. Create/edit user tidak menyimpan nama.  
**Fix Required:** Ganti semua `name` → `nama` di `admin/users/page.js` (EMPTY_FORM, openEdit, form binding, payload).

---

### BUG-005 — 🟡 Medium
**Page:** `/settings/telegram`  
**Feature:** Load & Save pengaturan  
**Description:** Form menggunakan field names yang tidak cocok dengan API. API mengembalikan `{ nama_grup, chat_id, aktif, jam_summary, threshold_pct }`, sedangkan form menggunakan `{ bot_token, chat_id, daily_summary_time, anomaly_threshold, enabled }`.  
**Root Cause:** Field name mismatch antara form state dan API contract.  
**Evidence:**
```
API GET response:  { nama_grup, chat_id, aktif, jam_summary, threshold_pct }
FE form fields:    { bot_token, chat_id, daily_summary_time, anomaly_threshold, enabled }

Mapping errors:
  - "bot_token"         → tidak ada di API (field tidak exist di TelegramConfig)
  - "daily_summary_time"→ API field adalah "jam_summary"
  - "anomaly_threshold" → API field adalah "threshold_pct"
  - "enabled"           → API field adalah "aktif"
  - "nama_grup"         → tidak ada di form (required oleh API PUT)

API PUT expects: { nama_grup, chat_id, jam_summary, threshold_pct, aktif }
```
**Impact:** Saat load, semua field tampil kosong/default. Saat save, request body salah → backend mengembalikan error atau tidak menyimpan data yang benar.  
**Fix Required:** Selaraskan form state dengan API fields: `nama_grup`, `chat_id`, `jam_summary`, `threshold_pct`, `aktif`. Hapus `bot_token` (tidak ada di API). Ganti label "Bot Token" dengan "Nama Grup".

---

### BUG-006 — 🟡 Medium
**Page:** `/admin/users`  
**Feature:** Create user  
**Description:** `usersAPI.deactivate` menggunakan method `DELETE`, tetapi backend endpoint untuk deactivate adalah `PATCH /api/v1/users/:id`. Method `DELETE` tidak terdaftar di router → akan mengembalikan 404 atau 405.  
**Root Cause:** `lib/api.js` — `deactivate` menggunakan `method: "DELETE"` yang salah.  
**Evidence:**
```js
// lib/api.js
deactivate: (id) =>
  apiFetch(`/api/v1/users/${id}`, { method: "DELETE" }), // ← WRONG

// Backend router: PATCH /api/v1/users/:id → Deactivate
```
**Impact:** Fungsi deactivate user tidak berfungsi (404/405 error).  
**Fix Required:** Ganti `method: "DELETE"` → `method: "PATCH"` di `usersAPI.deactivate`.

---

### BUG-007 — 🟡 Medium
**Page:** `/admin/produk`  
**Feature:** Toggle aktif/nonaktif  
**Description:** `handleToggleAktif` memanggil `produkAPI.update(row.id, { aktif: !row.aktif })` dengan payload yang hanya berisi `aktif`. Backend `PUT /api/v1/produk/:id` adalah full-replace — field yang tidak dikirim akan menjadi zero-value (string kosong, 0, false). Seharusnya menggunakan `PATCH /api/v1/produk/:id` yang sudah tersedia di backend.  
**Root Cause:** Menggunakan PUT (full update) untuk partial update.  
**Evidence:**
```js
// FE: handleToggleAktif
await produkAPI.update(row.id, { aktif: !row.aktif });
// Sends: PUT /api/v1/produk/:id { aktif: true }
// Backend overwrites: nama="", kode_produk="", harga=0, stok=0, ...
```
**Impact:** Toggle aktif akan menghapus semua data produk lainnya (nama, harga, stok menjadi kosong/0).  
**Fix Required:** Tambah `produkAPI.deactivate` di `lib/api.js` menggunakan `PATCH`, dan gunakan itu di `handleToggleAktif`.

---

### BUG-008 — 🟢 Low
**Page:** `/admin/customer`  
**Feature:** Create customer  
**Description:** Form tidak memiliki field `kode_cust`, yang merupakan required field di backend (`CreateRequest.KodeCust`). Tanpa `kode_cust`, create akan gagal atau menyimpan string kosong yang melanggar unique constraint.  
**Root Cause:** `EMPTY_FORM` tidak memiliki `kode_cust`. Form tidak memiliki input untuk field ini.  
**Fix Required:** Tambah field `kode_cust` ke `EMPTY_FORM` dan tambah input di form modal.

---

### BUG-009 — 🟢 Low
**Page:** `/admin/users`  
**Feature:** Role badge untuk "viewer"  
**Description:** `ROLE_OPTIONS` hanya mendefinisikan `admin`, `manager`, `sales`. Role `viewer` (yang ada di seed data) tidak memiliki konfigurasi warna/label, sehingga `RoleBadge` fallback ke `ROLE_OPTIONS[2]` (Sales) dan menampilkan badge "Sales" untuk user dengan role "viewer".  
**Root Cause:** `ROLE_OPTIONS` tidak lengkap — backend mendukung 4 role: `admin`, `manager`, `sales`, `viewer`.  
**Fix Required:** Tambah entry `viewer` ke `ROLE_OPTIONS`.

---

## Test Results per Endpoint

| Endpoint | Method | Backend | FE Integration | Status |
|----------|--------|---------|----------------|--------|
| `/api/v1/auth/login` | POST | ✅ OK | ✅ OK | ✅ Pass |
| `/api/v1/auth/logout` | POST | ✅ OK | ✅ OK | ✅ Pass |
| `/api/v1/auth/me` | GET | ✅ OK | ✅ OK | ✅ Pass |
| `/api/v1/produk` | GET | ✅ OK | 🔴 BUG-001 field mismatch | ❌ Fail |
| `/api/v1/produk` | POST | ✅ OK | 🔴 BUG-001, BUG-002 | ❌ Fail |
| `/api/v1/produk/:id` | PUT | ✅ OK | 🔴 BUG-001, BUG-002 | ❌ Fail |
| `/api/v1/produk/:id` | PATCH | ✅ OK | 🟡 BUG-007 (PUT used instead) | ❌ Fail |
| `/api/v1/customer` | GET | ✅ OK | 🔴 BUG-003 field mismatch | ❌ Fail |
| `/api/v1/customer` | POST | ✅ OK | 🔴 BUG-003, 🟢 BUG-008 | ❌ Fail |
| `/api/v1/customer/:id` | PUT | ✅ OK | 🔴 BUG-003 | ❌ Fail |
| `/api/v1/users` | GET | ✅ OK | 🔴 BUG-004 field mismatch | ❌ Fail |
| `/api/v1/users` | POST | ✅ OK | 🔴 BUG-004 | ❌ Fail |
| `/api/v1/users/:id` | PUT | ✅ OK | 🔴 BUG-004 | ❌ Fail |
| `/api/v1/users/:id` | PATCH | ✅ OK | 🟡 BUG-006 (DELETE used) | ❌ Fail |
| `/api/v1/settings/telegram` | GET | ✅ OK | 🟡 BUG-005 field mismatch | ❌ Fail |
| `/api/v1/settings/telegram` | PUT | ✅ OK | 🟡 BUG-005 field mismatch | ❌ Fail |

---

## Fixes Applied by Senior FE Developer

All 9 bugs fixed on 27 April 2026.

| Bug | File | Fix |
|-----|------|-----|
| BUG-001 | `admin/produk/page.js` | `kategori` → `kategori_pakaian` in column key, openEdit, searchKeys |
| BUG-002 | `admin/produk/page.js` | Added `kode_produk` to EMPTY_FORM and form modal |
| BUG-003 | `admin/customer/page.js` | Removed `kota` field; replaced with `alamat` column |
| BUG-004 | `admin/users/page.js` | `name` → `nama` in EMPTY_FORM, openEdit, column key, searchKeys, payload |
| BUG-005 | `settings/telegram/page.js` | Aligned all form fields: `nama_grup`, `chat_id`, `jam_summary`, `threshold_pct`, `aktif` |
| BUG-006 | `lib/api.js` | `usersAPI.deactivate`: `DELETE` → `PATCH` |
| BUG-007 | `lib/api.js` + `admin/produk/page.js` | Added `produkAPI.deactivate` (PATCH); `handleToggleAktif` now calls it |
| BUG-008 | `admin/customer/page.js` | Added `kode_cust` to EMPTY_FORM and form modal |
| BUG-009 | `admin/users/page.js` | Added `viewer` role to `ROLE_OPTIONS`; role grid changed to 2×2 |

**Status: ✅ All bugs resolved — ready for re-test**
