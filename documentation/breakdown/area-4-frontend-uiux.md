# Area 4 — Frontend & UI/UX (Next.js 14 + TypeScript)

> **Tujuan:** Membangun antarmuka yang modern, intuitif, dan bisa digunakan oleh pengguna non-teknis tanpa pelatihan lebih dari 5 menit.

---

## 4.1 Setup

- [ ] **Next.js 14** + TypeScript
- [ ] **shadcn/ui** sebagai base component library (customizable)
- [ ] **Tailwind CSS** dengan custom token (biru-putih, dark mode navy)
- [ ] **React Query** untuk data fetching + caching
- [ ] **Axios** instance dengan JWT interceptor (auto-refresh atau redirect ke login)

---

## 4.2 Design System

### Token Desain

| Token | Nilai |
|---|---|
| Primary Color | `#2563eb` (Blue 600) |
| Surface Light | `#ffffff` |
| Surface Dark | `#0f172a` (Slate 900) |
| Font | Inter (Google Fonts) |
| Border Radius | `0.5rem` |
| Shadow | `sm` |

### Base Components

> [!IMPORTANT]
> Komponen-komponen di bawah ini harus dibuat **satu kali** dan dipakai di seluruh halaman. Jangan membuat ulang komponen yang sama di setiap halaman.

- [ ] `Button` — variant: `primary`, `secondary`, `danger`, `ghost`
- [ ] `Input`, `Select`, `Textarea` — wajib punya label + error state
- [ ] `Modal` / `Dialog`
- [ ] `Badge` / `StatusBadge` — untuk status order (pending, confirmed, paid, shipped, closed, cancelled)
- [ ] `DataTable` — sortable + pagination
- [ ] `Alert` / `Toast` — feedback operasi berhasil/gagal
- [ ] `Skeleton` loader — tampil saat data sedang di-fetch
- [ ] `EmptyState` — tampil saat data kosong
- [ ] `PageHeader` — judul halaman + breadcrumb + action button
- [ ] `ChatWidget` — floating AI chat bubble (pojok kanan bawah)

---

## 4.3 Halaman Toko (Public — Tidak Perlu Login)

| Halaman | Isi & Fitur |
|---|---|
| `/` — Beranda | Hero section, grid produk, filter kategori/ukuran/warna |
| `/produk/:id` | Foto produk, detail lengkap, pilih ukuran & warna, tombol order |

### ChatWidget (AI Chat Assistant)

- [ ] Floating button → panel chat slide-up
- [ ] SSE streaming response dari backend
- [ ] Typing indicator saat AI sedang memproses
- [ ] Mobile responsive

---

## 4.4 Autentikasi

- [ ] `/login` — Form email + password + validasi inline Bahasa Indonesia
- [ ] Protected route HOC — redirect ke `/login` jika tidak ada token valid
- [ ] Redirect otomatis ke `/login` jika token expired
- [ ] Pesan error: Bahasa Indonesia, user-friendly (bukan pesan teknis)

---

## 4.5 Admin — Master Data

| Halaman | Fitur |
|---|---|
| `/admin/produk` | DataTable produk + tombol Tambah/Edit/Nonaktifkan |
| `/admin/produk/baru` | Form: nama, kategori, ukuran, warna, bahan, harga, stok |
| `/admin/customer` | DataTable customer + CRUD |
| `/admin/users` | DataTable user + form tambah + set role + `telegram_user_id` |
| `/settings/telegram` | Form konfigurasi Telegram (chat_id, jam summary, threshold) |

---

## 4.6 Sales — Transaksi

| Halaman | Fitur |
|---|---|
| `/sales/orders` | Daftar order + filter status |
| `/sales/orders/baru` | Form order: pilih customer, produk, ukuran, warna, qty |
| `/sales/orders/:id` | Detail order + stepper status + tombol aksi sesuai status |

### Stepper Status Order

```
Pending → Confirmed → Paid → Shipped → Closed
                                    ↘ Cancelled
```

- [ ] Stepper visual menampilkan tahapan order secara jelas
- [ ] Modal konfirmasi **setiap** perubahan status (tidak ada perubahan langsung tanpa konfirmasi)

---

## 4.7 Dashboard (Manager + Admin)

### Komponen Dashboard

| Komponen | Fungsi |
|---|---|
| `ReportSelector` | Dropdown pilih laporan + deskripsi singkat laporan |
| `FilterBar` | Filter: date range, sales, produk, kategori |
| `ChartRenderer` | Render Line / Bar / Pie / Funnel sesuai respons AI |
| `AIInsightCard` | Card biru berisi teks insight 2-3 kalimat dari AI |
| `AnomalyFlag` | Banner ⚠️ + penjelasan anomali + rekomendasi tindakan |

### Checklist Dashboard

- [ ] `ReportSelector` — dropdown + deskripsi laporan
- [ ] `FilterBar` — date range, sales, produk, kategori
- [ ] `ChartRenderer` — Line / Bar / Pie / Funnel sesuai respons AI
- [ ] `AIInsightCard` — card biru teks insight 2-3 kalimat
- [ ] `AnomalyFlag` — banner ⚠️ + penjelasan + rekomendasi
- [ ] Skeleton loader saat AI memproses (jangan tampilkan halaman kosong)
- [ ] Download PDF menggunakan `html2canvas` + `jsPDF`

---

## 4.8 UX Rules (Wajib Dipatuhi)

> [!IMPORTANT]
> Rules di bawah ini bersifat **wajib** dan harus dipatuhi di **setiap halaman**.

- [ ] **Setiap halaman** harus punya: loading state, empty state, error state
- [ ] **Semua form** harus punya validasi inline dalam **Bahasa Indonesia**
- [ ] Maksimal **3 klik** dari halaman utama ke informasi apapun
- [ ] Sidebar collapsible + top navbar (navigasi 2 level, tidak ada sub-sub-menu)
- [ ] Semua label, tombol, dan pesan menggunakan **Bahasa Indonesia sehari-hari** (tidak ada jargon teknis)

---

## Struktur Folder Next.js (Rekomendasi)

```
/src
  /components
    /ui          ← Base components (Button, Input, Modal, dll)
    /shared      ← Reusable components (DataTable, PageHeader, dll)
    /features    ← Feature-specific components (ChatWidget, ChartRenderer, dll)
  /pages
    /admin
    /sales
    /dashboard
  /hooks         ← Custom hooks (useAuth, useReports, dll)
  /services      ← API clients (axios instances)
  /types         ← TypeScript interfaces & types
  /lib           ← Utilities & helpers
```

> [!TIP]
> Gunakan struktur folder `ui/base → shared → features → pages` untuk memisahkan tanggung jawab komponen dengan jelas dan menghindari duplikasi kode.
