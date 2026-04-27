# InsightFlow — Frontend Dashboard

> **AI-Powered Sales Analytics Dashboard**  
> Built with Next.js 14, Tailwind CSS, and Recharts

---

## 🚀 Quick Start

```bash
# Install dependencies
npm install

# Development server
npm run dev

# Production build
npm run build
npm start
```

The app will be available at `http://localhost:3000`

---

## 📁 Project Structure

```
fe-penjualan/
├── app/
│   ├── page.js                    # Dashboard utama (protected)
│   ├── login/page.js              # Login page
│   ├── admin/
│   │   ├── produk/page.js         # CRUD Produk
│   │   ├── customer/page.js       # CRUD Customer
│   │   └── users/page.js          # CRUD Users
│   └── settings/
│       └── telegram/page.js       # Telegram bot settings
├── components/
│   ├── ProtectedRoute.js          # Auth guard HOC
│   ├── Sidebar.js                 # Navigation sidebar
│   ├── Topbar.js                  # Top navigation bar
│   ├── DataTable.js               # Reusable table with search & pagination
│   ├── Modal.js                   # Reusable modal dialog
│   ├── KPICard.js                 # KPI metric card
│   ├── ChartRenderer.js           # Dynamic chart renderer (Recharts)
│   ├── AIInsightCard.js           # AI-generated insight display
│   ├── AnomalyFlag.js             # Anomaly alert card
│   ├── FilterBar.js               # Date range filter
│   ├── ReportSelector.js          # Report type dropdown
│   └── TopProdukTable.js          # Top products table
└── lib/
    ├── api.js                     # API client (fetch wrapper)
    ├── auth.js                    # Auth helpers (token, user)
    └── dummyData.js               # Dummy data for prototype
```

---

## 🎨 Pages Overview

### 1. **Dashboard** (`/`)
- **Protected:** ✅ Yes
- **Features:**
  - 4 KPI cards (Omzet, Order, Customer, Avg Order)
  - Self-service report selector (8 report types)
  - Dynamic chart rendering (Line, Bar, Pie, Funnel)
  - AI-generated insights via n8n + LLM
  - Anomaly detection alerts
  - Top products table
  - Sales performance tracker

### 2. **Login** (`/login`)
- **Protected:** ❌ No
- **Features:**
  - Email + password authentication
  - Demo credentials quick-fill
  - JWT token storage in localStorage
  - Auto-redirect to dashboard on success

### 3. **Admin — Produk** (`/admin/produk`)
- **Protected:** ✅ Yes
- **Features:**
  - DataTable with search & pagination
  - Create / Edit produk (modal form)
  - Toggle aktif/nonaktif status
  - Stock level indicators (Kritis / Rendah / Aman)
  - Real-time API integration

### 4. **Admin — Customer** (`/admin/customer`)
- **Protected:** ✅ Yes
- **Features:**
  - DataTable with search & pagination
  - Create / Edit customer (modal form)
  - Display total orders per customer
  - Contact info (phone, email, address)

### 5. **Admin — Users** (`/admin/users`)
- **Protected:** ✅ Yes
- **Features:**
  - DataTable with search & pagination
  - Create / Edit users (modal form)
  - Role management (Admin, Manager, Sales)
  - Telegram User ID linking (for Q&A bot)
  - Password update (optional on edit)

### 6. **Settings — Telegram** (`/settings/telegram`)
- **Protected:** ✅ Yes
- **Features:**
  - Bot token configuration
  - Chat ID setup
  - Daily summary schedule (time picker)
  - Anomaly threshold setting (%)
  - Enable/disable bot toggle

---

## 🔐 Authentication

- **Method:** JWT Bearer Token
- **Storage:** `localStorage` (token + user object)
- **Protected Routes:** Wrapped with `<ProtectedRoute>` HOC
- **Auto-redirect:** Unauthenticated users → `/login`
- **Logout:** Clears token + redirects to login

---

## 🎯 API Integration

All API calls go through `lib/api.js`:

```javascript
import { authAPI, produkAPI, customerAPI, usersAPI, settingsAPI } from "@/lib/api";

// Example: Login
const res = await authAPI.login("admin@insightflow.id", "password123");
localStorage.setItem("token", res.token);

// Example: Fetch products
const products = await produkAPI.list();
```

**Base URL:** Set via `NEXT_PUBLIC_API_URL` env var (defaults to `http://localhost:3032`)

---

## 🎨 Design System

### Colors
- **Primary:** Indigo (`#6366f1`)
- **Background:** Slate 950 (`#0f172a`)
- **Card:** Slate 800 (`#1e293b`)
- **Border:** Slate 700 (`#334155`)
- **Text:** Slate 100 (`#f1f5f9`)

### Typography
- **Font:** Inter (Google Fonts)
- **Weights:** 300, 400, 500, 600, 700, 800

### Components
- **Buttons:** Rounded-xl, shadow-lg, hover states
- **Cards:** Rounded-2xl, border, gradient backgrounds
- **Inputs:** Rounded-xl, focus ring, indigo accent
- **Tables:** Hover states, zebra striping, responsive

---

## 📊 Chart Types

Powered by **Recharts**:

1. **Line Chart** — Penjualan Harian (dual Y-axis: omzet + order)
2. **Bar Chart** — Top Produk, Performa Sales
3. **Pie Chart** — Distribusi Kategori (donut style)
4. **Funnel Chart** — Order funnel (masuk → selesai)

All charts support:
- Responsive container
- Custom tooltips (dark theme)
- Smooth animations
- Accessible color palette

---

## 🧩 Reusable Components

### `<DataTable>`
Generic table with built-in search, pagination, and custom column renderers.

```jsx
<DataTable
  columns={[
    { key: "nama", label: "Nama", render: (v) => <strong>{v}</strong> },
    { key: "harga", label: "Harga", align: "right" },
  ]}
  data={products}
  searchKeys={["nama", "kategori"]}
  emptyMessage="Tidak ada data."
/>
```

### `<Modal>`
Overlay modal with backdrop, ESC to close, and size variants.

```jsx
<Modal open={isOpen} onClose={() => setIsOpen(false)} title="Edit Produk" size="lg">
  <form>...</form>
</Modal>
```

### `<ProtectedRoute>`
Auth guard that redirects to `/login` if no token found.

```jsx
export default function AdminPage() {
  return (
    <ProtectedRoute>
      <AdminContent />
    </ProtectedRoute>
  );
}
```

---

## 🔧 Environment Variables

Create `.env.local`:

```bash
NEXT_PUBLIC_API_URL=http://localhost:8080
```

---

## 🚦 Development Workflow

1. **Start backend:** `cd be-penjualan && make run`
2. **Start frontend:** `npm run dev`
3. **Login:** Use demo credentials from login page
4. **Test CRUD:** Navigate to `/admin/produk`, `/admin/customer`, `/admin/users`
5. **Configure Telegram:** Go to `/settings/telegram`

---

## 📦 Dependencies

```json
{
  "next": "14.2.35",
  "react": "^18",
  "react-dom": "^18",
  "recharts": "^3.8.1",
  "lucide-react": "^1.11.0",
  "tailwindcss": "^3.4.1"
}
```

---

## 🎯 Next Steps (Post-MVP)

- [ ] Real-time data refresh (polling / SSE)
- [ ] Export PDF reports
- [ ] AI Chat widget (SSE streaming)
- [ ] Public storefront pages
- [ ] Mobile responsive optimization
- [ ] Dark/light theme toggle
- [ ] Multi-language support (i18n)

---

## 📝 Notes

- **Dummy Data:** Currently using `lib/dummyData.js` for charts. Will be replaced with real API calls in Sprint 4.
- **AI Insights:** Generated by n8n workflow → LLM → returned to frontend.
- **Telegram Integration:** Configured via Settings page, executed by n8n workflows.

---

**Built with ❤️ for InsightFlow · April 2026**
