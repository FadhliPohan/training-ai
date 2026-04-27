# n8n Workflows — InsightFlow

3 workflow siap import ke n8n. Ikuti langkah setup di bawah.

---

## Cara Import

1. Buka n8n UI: `http://localhost:5678`
2. Klik **Add workflow** → **Import from file**
3. Upload salah satu file dari folder ini
4. Ulangi untuk semua 3 workflow

---

## Setup Credentials (lakukan sekali)

### 1. OpenAI
> n8n → Settings → Credentials → Add → **OpenAI**
- **Name:** `OpenAI account`
- **API Key:** `sk-...` (dari platform.openai.com)

### 2. Telegram Bot
> n8n → Settings → Credentials → Add → **Telegram**
- **Name:** `Telegram account`
- **Bot Token:** token dari @BotFather

---

## Konfigurasi Per-Workflow

### Workflow 01 — Daily Summary

| Yang Perlu Diubah | Node | Field | Nilai |
|---|---|---|---|
| Backend URL | `Get Summary Data` | URL | `http://localhost:8080` → sesuaikan production URL |
| Internal API Key | `Get Summary Data` | Header `X-Internal-Key` | Isi dengan nilai `INTERNAL_API_KEY` dari `.env` |
| Telegram Chat ID | `Send to Telegram` | Chat ID | Isi dengan grup/user chat_id Telegram tujuan |
| Telegram Chat ID | `Send Error Alert` | Chat ID | Sama dengan di atas |

**Cara dapat Chat ID:**
- Tambahkan bot ke grup Telegram
- Kirim pesan di grup
- Buka `https://api.telegram.org/bot<TOKEN>/getUpdates`
- Cari `"chat": {"id": ...}` → itu adalah chat_id

### Workflow 02 — Telegram Q&A

| Yang Perlu Diubah | Node | Field | Nilai |
|---|---|---|---|
| Backend URL | `Get Report Data` | URL (di-generate otomatis) | Ubah `localhost:8080` di node **Build Endpoint URL** baris `const BASE = '...'` |
| Internal API Key | `Build Endpoint URL` | baris `const KEY = '...'` | Isi dengan nilai `INTERNAL_API_KEY` dari `.env` |
| Webhook URL | `Webhook from Golang` | — | Setelah aktif, copy webhook URL dan update `N8N_TELEGRAM_WEBHOOK_PATH` di `.env` |

**Update `.env` backend setelah dapat URL:**
```
N8N_BASE_URL=http://localhost:5678
N8N_TELEGRAM_WEBHOOK_PATH=/webhook/telegram-qa
```

### Workflow 03 — Dashboard AI

| Yang Perlu Diubah | Node | Field | Nilai |
|---|---|---|---|
| Webhook path | `Webhook from Backend` | Path | Pastikan `dashboard-ai` — harus cocok dengan `N8N_DASHBOARD_WEBHOOK_PATH` |

---

## Cara Aktifkan Workflow

1. Buka workflow
2. Klik toggle **Inactive → Active** (kanan atas)
3. Workflow mulai berjalan

> ⚠️ Aktifkan **03-dashboard-ai** terlebih dulu karena langsung dipakai oleh `GET /api/v1/reports`.

---

## Test Manual

### Test Workflow 01 (Daily Summary)
```
Di n8n: buka workflow → klik "Test workflow" → lihat output di setiap node
```

### Test Workflow 02 (Telegram Q&A)
```bash
# Simulate payload dari Golang:
curl -X POST http://localhost:5678/webhook/telegram-qa \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "uuid-sales-1",
    "user_nama": "Citra Sales",
    "role": "sales",
    "telegram_user_id": 123456789,
    "chat_id": 123456789,
    "message_text": "penjualan hari ini berapa?",
    "message_id": 1
  }'
```

### Test Workflow 03 (Dashboard AI)
```bash
curl -X POST http://localhost:5678/webhook/dashboard-ai \
  -H "Content-Type: application/json" \
  -d '{
    "report_type": "daily-sales",
    "data": [{"date": "2026-04-27", "value": 4750000}],
    "metrics": {"total_revenue": 4750000, "total_orders": 23},
    "filters": {"from": "2026-04-21", "to": "2026-04-27"}
  }'
```

---

## Checklist Go-Live

- [ ] n8n berjalan via Docker (`docker compose up -d n8n`)
- [ ] Credentials OpenAI & Telegram terkonfigurasi
- [ ] Workflow 03 aktif (Dashboard AI)
- [ ] Workflow 01 aktif (Daily Summary) + chat_id terisi
- [ ] Workflow 02 aktif (Telegram Q&A) + webhook URL di `.env` backend
- [ ] Backend restart setelah update `.env`
- [ ] Test kirim pesan Telegram ke bot
