# Area 1 — API Documentation

> **Tujuan:** Seluruh kontrak API terdefinisi sebelum implementasi agar frontend & backend bisa berjalan secara paralel.

---

## 1.1 Format & Standar

### Setup Tools

- [ ] Setup **Postman Collection** atau **OpenAPI 3.0 (Swagger UI)**
- [ ] Base URL + versioning: `/api/v1/`

### Format Respons Standar

```json
{
  "success": true,
  "message": "...",
  "data": {},
  "errors": null
}
```

### Format Error

| HTTP Status | Kode | Keterangan |
|---|---|---|
| `400` | Bad Request | Validation Error |
| `401` | Unauthorized | Unauthenticated |
| `403` | Forbidden | Tidak ada izin akses |
| `404` | Not Found | Resource tidak ditemukan |
| `409` | Conflict | Data konflik / duplikat |
| `500` | Internal Server Error | Kesalahan server |

---

## 1.2 Daftar Endpoint

| Group | Endpoint | Method | Role |
|---|---|---|---|
| Auth | `/auth/login` | POST | Public |
| Auth | `/auth/logout` | POST | Auth |
| Auth | `/auth/me` | GET | Auth |
| Produk | `/produk` | GET, POST | Auth |
| Produk | `/produk/:id` | GET, PUT, PATCH | Auth |
| Customer | `/customer` | GET, POST | Auth |
| Customer | `/customer/:id` | GET, PUT | Auth |
| Users | `/users` | GET, POST | Admin |
| Users | `/users/:id` | GET, PUT, PATCH | Admin |
| Order | `/orders` | GET, POST | Auth |
| Order | `/orders/:id` | GET | Auth |
| Order | `/orders/:id/confirm` | POST | Sales/Admin |
| Order | `/orders/:id/cancel` | POST | Sales/Admin |
| Pembayaran | `/payments` | POST | Sales/Admin |
| Pembayaran | `/payments/:id/verify` | POST | Admin |
| Pengiriman | `/shipments` | POST | Sales/Admin |
| Pengiriman | `/shipments/:id` | PUT | Sales/Admin |
| Dashboard | `/reports?type=&from=&to=` | GET | Manager/Admin |
| AI Chat | `/chat/stream?message=` | GET (SSE) | Public |
| Settings | `/settings/telegram` | GET, PUT | Admin |
| Internal n8n | `/internal/ai-result` | POST | Internal |

---

## 1.3 Deliverables

- [ ] Postman Collection `.json` di-commit ke repo
- [ ] Swagger UI dihosting di `/api/docs`
- [ ] Contoh request + response per endpoint
- [ ] Dokumentasi semua error code dan artinya

---

## Catatan Implementasi

> [!IMPORTANT]
> Area API Documentation adalah **prerequisite** sebelum tim backend dan frontend bisa mulai coding secara paralel. Selesaikan ini di Minggu 1.

> [!NOTE]
> Endpoint `/chat/stream` menggunakan **Server-Sent Events (SSE)** bukan WebSocket. Header yang wajib di-set:
> - `Content-Type: text/event-stream`
> - `Cache-Control: no-cache`
> - `Connection: keep-alive`

> [!TIP]
> Gunakan Swagger UI yang dihosting di `/api/docs` agar frontend developer bisa langsung mencoba endpoint secara interaktif tanpa perlu import Postman Collection.
