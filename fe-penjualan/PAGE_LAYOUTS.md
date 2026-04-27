# Page Layouts — Visual Guide

> **InsightFlow Frontend Pages**  
> Visual reference for all page layouts

---

## 🏠 Dashboard (`/`)

```
┌─────────────────────────────────────────────────────────────┐
│ [☰] InsightFlow                    [Search] [🔄] [🔔] [AD] │ ← Topbar
├──────┬──────────────────────────────────────────────────────┤
│      │  📊 Dashboard Analitik                               │
│ [📊] │  Senin, 27 April 2026                                │
│ Dash │                                                       │
│      │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐   │
│ [📦] │  │ 💰 Omzet│ │ 📦 Order│ │ 👥 Cust │ │ 📊 Avg  │   │
│ Prod │  │ 187.5Jt │ │   342   │ │   128   │ │  548Rb  │   │
│      │  │ ↑ 12.4% │ │ ↑ 8.7%  │ │ ↓ 3.2%  │ │ ↑ 5.1%  │   │
│ [👥] │  └─────────┘ └─────────┘ └─────────┘ └─────────┘   │
│ Cust │                                                       │
│      │  ┌─────────────────────────────────────────────────┐ │
│ [👤] │  │ 📊 Laporan Self-Service                         │ │
│ User │  │                                                  │ │
│      │  │ [Pilih Laporan ▼] [30d] [Tampilkan] [Export]   │ │
│ [⚙️] │  │                                                  │ │
│ Sett │  │  ┌────────────────────────────────────────┐    │ │
│      │  │  │                                         │    │ │
│      │  │  │         📈 Chart Area                   │    │ │
│      │  │  │                                         │    │ │
│      │  │  └────────────────────────────────────────┘    │ │
│      │  │                                                  │ │
│      │  │  ✨ AI Insight: "Penjualan menunjukkan..."     │ │
│      │  └─────────────────────────────────────────────────┘ │
│      │                                                       │
│      │  ┌──────────────────────┐ ┌────────────────────┐   │
│      │  │ 🏆 Top Produk        │ │ ⚠️ Anomali         │   │
│      │  │ [Table with 7 rows]  │ │ [2 alerts]         │   │
│      │  └──────────────────────┘ │                    │   │
│      │                            │ 👥 Performa Sales  │   │
│      │                            │ [5 progress bars]  │   │
│      │                            └────────────────────┘   │
└──────┴──────────────────────────────────────────────────────┘
  ↑
Sidebar (collapsible)
```

---

## 🔐 Login (`/login`)

```
┌─────────────────────────────────────────────────────────────┐
│                                                               │
│                                                               │
│                        ⚡                                     │
│                   InsightFlow                                │
│              AI Sales Dashboard                              │
│                                                               │
│              ┌─────────────────────────┐                     │
│              │                         │                     │
│              │  Email                  │                     │
│              │  [________________]     │                     │
│              │                         │                     │
│              │  Password               │                     │
│              │  [________________] 👁  │                     │
│              │                         │                     │
│              │  [Masuk ke Dashboard]   │                     │
│              │                         │                     │
│              │  ─────────────────────  │                     │
│              │  Demo credentials       │                     │
│              │  [Admin] [Sales]        │                     │
│              └─────────────────────────┘                     │
│                                                               │
│              © 2026 InsightFlow                              │
└─────────────────────────────────────────────────────────────┘
```

---

## 📦 Admin Produk (`/admin/produk`)

```
┌─────────────────────────────────────────────────────────────┐
│ [☰] InsightFlow                    [Search] [🔄] [🔔] [AD] │
├──────┬──────────────────────────────────────────────────────┤
│      │  📦 Manajemen Produk                                 │
│ [📊] │  Senin, 27 April 2026                                │
│ Dash │                                                       │
│      │  Daftar Produk                    [+ Tambah Produk]  │
│ [📦] │  42 produk terdaftar                                 │
│ Prod │                                                       │
│      │  [Search: ___________]                               │
│ [👥] │                                                       │
│ Cust │  ┌─────────────────────────────────────────────────┐ │
│      │  │ Nama      │ Kat │ Ukuran │ Harga │ Stok │ Aksi │ │
│ [👤] │  ├─────────────────────────────────────────────────┤ │
│ User │  │ Kemeja... │ Kem │ L      │ 185K  │ 🟢13 │ ✏️   │ │
│      │  │ Kaos...   │ Kao │ M      │ 75K   │ 🟢57 │ ✏️   │ │
│ [⚙️] │  │ Celana... │ Cel │ 32     │ 220K  │ 🟡8  │ ✏️   │ │
│ Sett │  │ Jaket...  │ Jak │ XL     │ 350K  │ 🔴6  │ ✏️   │ │
│      │  │ ...                                              │ │
│      │  └─────────────────────────────────────────────────┘ │
│      │                                                       │
│      │  [◀] 1 2 3 ... 5 [▶]                                │
└──────┴──────────────────────────────────────────────────────┘

Modal (when editing):
┌─────────────────────────────────┐
│ Edit Produk                  [X]│
├─────────────────────────────────┤
│ Nama Produk *                   │
│ [_________________________]     │
│                                 │
│ Kategori *        Ukuran        │
│ [Kemeja ▼]        [L____]       │
│                                 │
│ Warna             Bahan         │
│ [Navy___]         [Cotton]      │
│                                 │
│ Harga (Rp) *      Stok *        │
│ [185000_]         [13___]       │
│                                 │
│ Status: [✓ Aktif]               │
│                                 │
│           [Batal]  [Simpan]     │
└─────────────────────────────────┘
```

---

## 👥 Admin Customer (`/admin/customer`)

```
┌─────────────────────────────────────────────────────────────┐
│ [☰] InsightFlow                    [Search] [🔄] [🔔] [AD] │
├──────┬──────────────────────────────────────────────────────┤
│      │  👥 Manajemen Customer                               │
│ [📊] │  Senin, 27 April 2026                                │
│ Dash │                                                       │
│      │  Daftar Customer                [+ Tambah Customer]  │
│ [📦] │  128 customer terdaftar                              │
│ Prod │                                                       │
│      │  [Search: ___________]                               │
│ [👥] │                                                       │
│ Cust │  ┌─────────────────────────────────────────────────┐ │
│      │  │ Nama      │ Email      │ Telepon │ Kota │ Aksi │ │
│ [👤] │  ├─────────────────────────────────────────────────┤ │
│ User │  │ Budi S.   │ budi@...   │ 0812... │ JKT  │ ✏️   │ │
│      │  │ Siti A.   │ siti@...   │ 0813... │ BDG  │ ✏️   │ │
│ [⚙️] │  │ Ahmad R.  │ ahmad@...  │ 0821... │ SBY  │ ✏️   │ │
│ Sett │  │ ...                                              │ │
│      │  └─────────────────────────────────────────────────┘ │
│      │                                                       │
│      │  [◀] 1 2 3 ... 13 [▶]                               │
└──────┴──────────────────────────────────────────────────────┘
```

---

## 👤 Admin Users (`/admin/users`)

```
┌─────────────────────────────────────────────────────────────┐
│ [☰] InsightFlow                    [Search] [🔄] [🔔] [AD] │
├──────┬──────────────────────────────────────────────────────┤
│      │  👤 Manajemen Users                                  │
│ [📊] │  Senin, 27 April 2026                                │
│ Dash │                                                       │
│      │  Daftar Users                      [+ Tambah User]   │
│ [📦] │  15 user terdaftar                                   │
│ Prod │                                                       │
│      │  [🛡️ Admin] [👔 Manager] [👤 Sales]                 │
│ [👥] │  · Telegram ID diperlukan untuk fitur Q&A            │
│ Cust │                                                       │
│      │  [Search: ___________]                               │
│ [👤] │                                                       │
│ User │  ┌─────────────────────────────────────────────────┐ │
│      │  │ Nama    │ Email     │ Role  │ Telegram │ Aksi  │ │
│ [⚙️] │  ├─────────────────────────────────────────────────┤ │
│ Sett │  │ Admin   │ admin@... │ Admin │ 12345... │ ✏️    │ │
│      │  │ Citra D.│ citra@... │ Sales │ 67890... │ ✏️    │ │
│      │  │ Budi S. │ budi@...  │ Sales │ —        │ ✏️    │ │
│      │  │ ...                                              │ │
│      │  └─────────────────────────────────────────────────┘ │
│      │                                                       │
│      │  [◀] 1 2 [▶]                                        │
└──────┴──────────────────────────────────────────────────────┘
```

---

## ⚙️ Settings Telegram (`/settings/telegram`)

```
┌─────────────────────────────────────────────────────────────┐
│ [☰] InsightFlow                    [Search] [🔄] [🔔] [AD] │
├──────┬──────────────────────────────────────────────────────┤
│      │  ⚙️ Pengaturan Telegram                             │
│ [📊] │  Senin, 27 April 2026                                │
│ Dash │                                                       │
│      │  Konfigurasi Telegram Bot                            │
│ [📦] │  Atur bot token, chat ID, dan jadwal notifikasi      │
│ Prod │                                                       │
│      │  ┌─────────────────────────────────────────────────┐ │
│ [👥] │  │ 📤 Konfigurasi Bot                              │ │
│ Cust │  │                                                  │ │
│      │  │ Bot Token *                                      │ │
│ [👤] │  │ [••••••••••••••••••••••••••••••]                │ │
│ User │  │ Dapatkan dari @BotFather                         │ │
│      │  │                                                  │ │
│ [⚙️] │  │ Chat ID *                                        │ │
│ Sett │  │ [-1001234567890_______________]                  │ │
│      │  │ Gunakan @userinfobot untuk mendapatkan ID        │ │
│      │  │                                                  │ │
│      │  │ Status Bot              [🟢 Aktif ────────]      │ │
│      │  └─────────────────────────────────────────────────┘ │
│      │                                                       │
│      │  ┌─────────────────────────────────────────────────┐ │
│      │  │ 🔔 Jadwal & Notifikasi                          │ │
│      │  │                                                  │ │
│      │  │ ⏰ Waktu Daily Summary    ⚠️ Threshold Anomali  │ │
│      │  │ [07:00]                   [10___] %             │ │
│      │  │                                                  │ │
│      │  └─────────────────────────────────────────────────┘ │
│      │                                                       │
│      │  ℹ️ Fitur Telegram Bot:                             │
│      │  • Daily Summary: Laporan otomatis setiap pagi      │
│      │  • Anomaly Alert: Notifikasi real-time              │
│      │  • On-demand Q&A: Tanya jawab via chat              │
│      │                                                       │
│      │                          [Reset] [💾 Simpan]         │
└──────┴──────────────────────────────────────────────────────┘
```

---

## 🚫 404 Page (`/not-found`)

```
┌─────────────────────────────────────────────────────────────┐
│                                                               │
│                                                               │
│                                                               │
│                         404                                   │
│                                                               │
│              Halaman Tidak Ditemukan                         │
│     Maaf, halaman yang Anda cari tidak ada atau              │
│              telah dipindahkan.                              │
│                                                               │
│         [🏠 Kembali ke Dashboard]  [← Halaman Sebelumnya]   │
│                                                               │
│                                                               │
│              InsightFlow · AI Sales Dashboard                │
└─────────────────────────────────────────────────────────────┘
```

---

## ⚠️ Error Page (`/error`)

```
┌─────────────────────────────────────────────────────────────┐
│                                                               │
│                                                               │
│                        ⚠️                                     │
│                                                               │
│                 Terjadi Kesalahan                            │
│        Maaf, terjadi kesalahan saat memuat                   │
│                  halaman ini.                                │
│                                                               │
│              ┌─────────────────────────┐                     │
│              │ Error: Network timeout  │                     │
│              └─────────────────────────┘                     │
│                                                               │
│              [🔄 Coba Lagi]  [🏠 Kembali ke Dashboard]      │
│                                                               │
│                                                               │
│      Jika masalah berlanjut, hubungi administrator.         │
└─────────────────────────────────────────────────────────────┘
```

---

## 📱 Mobile Layout (375px)

```
┌─────────────────────┐
│ [☰] InsightFlow [🔔]│ ← Topbar (compact)
├─────────────────────┤
│ 📊 Dashboard        │
│ Senin, 27 Apr 2026  │
│                     │
│ ┌─────────────────┐ │
│ │ 💰 Total Omzet  │ │
│ │ Rp 187,5 Jt     │ │
│ │ ↑ 12.4%         │ │
│ └─────────────────┘ │
│                     │
│ ┌─────────────────┐ │
│ │ 📦 Total Order  │ │
│ │ 342             │ │
│ │ ↑ 8.7%          │ │
│ └─────────────────┘ │
│                     │
│ [Stacked layout]    │
│ [continues...]      │
│                     │
└─────────────────────┘

Sidebar (overlay):
┌─────────────────────┐
│ ⚡ InsightFlow      │
│ AI Sales Dashboard  │
├─────────────────────┤
│ 📊 Dashboard        │
│ 📦 Produk           │
│ 👥 Customer         │
│ 👤 Users            │
│ ⚙️ Pengaturan       │
├─────────────────────┤
│ [AD] Administrator  │
│ admin@...           │
│ 🚪 Keluar           │
└─────────────────────┘
```

---

## 🎨 Component States

### Button States
```
Normal:    [Simpan]
Hover:     [Simpan] (brighter)
Active:    [Simpan] (darker)
Disabled:  [Simpan] (grayed out)
Loading:   [⏳ Menyimpan...]
```

### Input States
```
Normal:    [____________]
Focus:     [____________] (blue ring)
Error:     [____________] (red border)
Disabled:  [____________] (grayed out)
```

### Table States
```
Normal:    │ Row data │
Hover:     │ Row data │ (highlighted)
Selected:  │ Row data │ (blue bg)
Empty:     │ Tidak ada data. │
Loading:   │ ⏳ Memuat... │
```

---

## 🎯 Responsive Breakpoints

```
Mobile:    < 640px   (1 column, stacked)
Tablet:    640-1024px (2 columns, collapsible sidebar)
Desktop:   > 1024px   (full layout, sidebar visible)
```

---

**Last updated:** April 27, 2026
