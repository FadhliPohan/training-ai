# Area 6 — n8n Automation & AI Workflows

> **Tujuan:** Membangun seluruh workflow AI yang menjadi "otak" InsightFlow — dari analisis dashboard, AI Chat Assistant, hingga notifikasi Telegram otomatis.

---

## 6.1 Setup n8n

- [ ] Deploy n8n via **Docker Compose** (self-hosted)
- [ ] Hanya bisa diakses via **internal network** (tidak terekspos ke publik)
- [ ] Credentials yang perlu dikonfigurasi di n8n:
  - OpenAI API Key / Gemini API Key
  - Telegram Bot Token
  - PostgreSQL connection
- [ ] Environment variables: backend base URL, webhook secret

> [!IMPORTANT]
> n8n adalah dependency kritis yang menghubungkan Backend, LLM, dan Telegram. Deploy dengan `restart: always` di Docker Compose agar otomatis restart jika crash.

---

## 6.2 Workflow 1 — Dashboard AI Analysis

**Trigger:** HTTP Webhook (dipanggil oleh Golang backend saat user request laporan)

### Input / Output

| Field | Detail |
|---|---|
| **Input** | `report_type`, `data[]` (data aggregat dari PostgreSQL), `filters` |
| **Output** | `chart_type`, `summary`, `anomalies[]`, `recommendation` |

### Steps

- [ ] **HTTP Webhook node** — terima data aggregat dari Golang
- [ ] **Code node** — format data aggregat menjadi prompt yang jelas
- [ ] **OpenAI/Gemini node** — kirim prompt ke LLM
- [ ] **Code node** — parse & validasi format JSON dari respons LLM
- [ ] **Respond to Webhook** — kembalikan JSON ke Golang

### System Prompt

```
Kamu adalah analis data penjualan pakaian. Berdasarkan data berikut: [DATA]
Tentukan: (1) jenis chart terbaik [line|bar|pie|funnel],
(2) ringkasan 2-3 kalimat Bahasa Indonesia,
(3) anomali jika varians > threshold%,
(4) 1 rekomendasi tindakan.
Kembalikan dalam format JSON.
```

### Format JSON Respons yang Diharapkan

```json
{
  "chart_type": "line",
  "summary": "Penjualan minggu ini mencapai Rp 45 juta...",
  "anomalies": [
    {
      "metric": "daily_revenue",
      "actual": 2000000,
      "expected": 5000000,
      "variance_pct": -60,
      "description": "Pendapatan turun signifikan pada Senin"
    }
  ],
  "recommendation": "Cek stok produk kategori atasan yang mungkin habis"
}
```

---

## 6.3 Workflow 2 — AI Chat Assistant (Customer)

**Trigger:** HTTP Webhook (dari Golang SSE handler)

### Input / Output

| Field | Detail |
|---|---|
| **Input** | `message` (pertanyaan customer), `context[]` (data produk relevan dari DB) |
| **Output** | `answer` (teks jawaban, di-stream ke SSE) |

### Steps

- [ ] **HTTP Webhook node** — terima pesan customer + context produk dari Golang
- [ ] **Code node** — susun system prompt + gabungkan context data produk
- [ ] **OpenAI node** — generate jawaban menggunakan data context
- [ ] **Respond to Webhook** — kembalikan jawaban ke Golang untuk di-stream via SSE

### System Prompt

```
Kamu adalah asisten toko pakaian InsightFlow. Jawab HANYA tentang produk
yang tersedia di toko. Data produk: [CONTEXT].
Jika pertanyaan di luar topik, tolak dengan sopan dalam Bahasa Indonesia.
```

---

## 6.4 Workflow 3 — Telegram Daily Summary

**Trigger:** Schedule — `cron: 0 0 * * *` (00:00 UTC = 07:00 WIB)

### Steps

- [ ] **Schedule Trigger** — `0 0 * * *`
- [ ] **PostgreSQL node** — aggregasi penjualan 24 jam terakhir
- [ ] **Code node** — format data + susun prompt
- [ ] **OpenAI/Gemini node** — generate ringkasan harian Bahasa Indonesia
- [ ] **Code node** — format pesan menjadi format Telegram-friendly

### Format Pesan Telegram

```
📊 *Ringkasan Penjualan — [TANGGAL]*

💰 Total Omzet: Rp X.XXX.XXX
📦 Total Order: XX
✅ Selesai: XX | ⏳ Pending: XX

🏆 Produk Terlaris: [Nama Produk]
[Ringkasan AI 2-3 kalimat]
```

- [ ] **Telegram node** — `sendMessage` ke semua `chat_id` aktif di `app.telegram_config`

---

## 6.5 Workflow 4 — Telegram Anomaly Alert

**Trigger:** Schedule setiap **15 menit**

### Steps

- [ ] **Schedule Trigger** — `*/15 * * * *`
- [ ] **PostgreSQL node** — ambil data terbaru vs rata-rata 7 hari terakhir
- [ ] **Code node** — hitung varians vs threshold dari `app.telegram_config`
- [ ] **IF node** — varians > threshold? → Lanjut : Stop (tidak kirim pesan)
- [ ] **OpenAI/Gemini node** — formulasikan pesan alert yang informatif
- [ ] **Telegram node** — `sendMessage` ke grup yang dikonfigurasi

### Format Alert Telegram

```
⚠️ *ALERT ANOMALI*
Metrik: [Nama Metrik] | Aktual: [X] | Ekspektasi: [Y]
Selisih: [Delta%]
💡 Rekomendasi: [Teks AI]
```

> [!NOTE]
> Threshold default adalah **10%** (dapat dikonfigurasi per grup via tabel `app.telegram_config`). Workflow harus membaca threshold dari DB setiap kali running, bukan hardcode.

---

## 6.6 Workflow 5 — Telegram Q&A Per-Role

**Trigger:** Telegram Trigger node (webhook dari Telegram API)

### Steps

- [ ] **Telegram Trigger** — terima pesan dari user di Telegram
- [ ] **PostgreSQL node** — query `app.users WHERE telegram_user_id = ?` untuk identifikasi user
- [ ] **IF node** — user terdaftar? → Tidak: balas "Akun tidak terdaftar"
- [ ] **Switch node** — route berdasarkan `role`:
  - `sales` → query data **milik sales itu saja** (`WHERE sales_id = user_id`)
  - `manager` → query **semua data tim**
- [ ] **Code node** — susun prompt + gabungkan data hasil query
- [ ] **OpenAI/Gemini node** — generate jawaban yang singkat dan kontekstual
- [ ] **Telegram node** — `sendMessage` ke `chat_id` pengirim

### Aturan Scope Data

| Role | Cakupan Data |
|---|---|
| `sales` | Hanya data order dengan `sales_id = user_id` |
| `manager` | Semua data tim tanpa filter sales |
| User tidak terdaftar | Tolak dengan pesan: "Maaf, akun Anda belum terdaftar di sistem." |

---

## 6.7 Error Handling (Semua Workflow)

> [!IMPORTANT]
> Setiap workflow **wajib** memiliki error handling. Workflow yang crash tanpa pesan yang jelas akan membuat pengguna bingung.

- [ ] **Error Trigger node** di setiap workflow
- [ ] LLM gagal → kembalikan **fallback message** (bukan error kosong atau pesan teknis)
- [ ] Telegram gagal → **retry 2x** dengan delay 30 detik
- [ ] Log semua error ke n8n execution history

### Fallback Message (LLM Gagal)

```
Maaf, saya sedang tidak bisa memproses permintaan Anda saat ini.
Silakan coba lagi dalam beberapa menit atau hubungi admin sistem.
```

---

## Ringkasan 5 Workflow

| # | Nama | Trigger | Fungsi |
|---|---|---|---|
| 1 | Dashboard AI Analysis | HTTP Webhook (dari Golang) | Analisis data & tentukan chart type + insight |
| 2 | AI Chat Assistant | HTTP Webhook (dari Golang SSE) | Jawab pertanyaan customer tentang produk |
| 3 | Telegram Daily Summary | Cron 07:00 WIB | Kirim ringkasan penjualan harian |
| 4 | Telegram Anomaly Alert | Cron setiap 15 menit | Kirim alert jika ada anomali data |
| 5 | Telegram Q&A Per-Role | Telegram Trigger | Jawab pertanyaan data per role sales/manager |

---

## Dependency Map

```
Setup n8n (6.1) harus selesai sebelum semua workflow lain bisa dibuat.

Workflow 1 (Dashboard AI) → Integrasi dengan Backend /reports
Workflow 2 (AI Chat) → Integrasi dengan Backend /chat/stream
Workflow 3, 4 (Telegram) → Independent, bisa paralel setelah DB migration selesai
Workflow 5 (Telegram Q&A) → Membutuhkan Telegram Bot Token + DB migration selesai
```

> [!TIP]
> Mulai dari Workflow 1 (Dashboard AI Analysis) karena ini yang paling sering digunakan dan paling complex. Workflow Telegram bisa dikerjakan paralel setelah DB migration selesai.
