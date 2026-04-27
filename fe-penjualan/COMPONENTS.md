# Component Documentation

> **InsightFlow Frontend Components**  
> Reusable React components for the dashboard

---

## 🧩 Layout Components

### `<Sidebar>`
**Location:** `components/Sidebar.js`

Collapsible navigation sidebar with user profile and logout.

**Features:**
- Auto-active state based on current route
- Collapsible on desktop (toggle button)
- Mobile overlay with backdrop
- User avatar with initials
- Logout functionality

**Usage:**
```jsx
import Sidebar from "@/components/Sidebar";

<Sidebar />
```

---

### `<Topbar>`
**Location:** `components/Topbar.js`

Top navigation bar with title, date, search, and actions.

**Props:**
- `title` (string): Page title
- `onRefresh` (function): Callback for refresh button

**Usage:**
```jsx
import Topbar from "@/components/Topbar";

<Topbar title="Dashboard Analitik" onRefresh={fetchData} />
```

---

## 🔐 Auth Components

### `<ProtectedRoute>`
**Location:** `components/ProtectedRoute.js`

HOC that wraps pages requiring authentication. Redirects to `/login` if no token found.

**Usage:**
```jsx
import ProtectedRoute from "@/components/ProtectedRoute";

export default function AdminPage() {
  return (
    <ProtectedRoute>
      <YourPageContent />
    </ProtectedRoute>
  );
}
```

---

## 📊 Dashboard Components

### `<KPICard>`
**Location:** `components/KPICard.js`

Animated KPI metric card with trend indicator and mini sparkline.

**Props:**
- `kpi` (object):
  - `id` (string)
  - `label` (string)
  - `value` (number)
  - `valueFormatted` (string)
  - `change` (number): Percentage change
  - `changeLabel` (string)
  - `icon` (string): "revenue" | "orders" | "customers" | "avgorder"
  - `color` (string): "indigo" | "violet" | "purple" | "sky"
  - `trend` (string): "up" | "down"
- `index` (number): For staggered animation delay

**Usage:**
```jsx
<KPICard
  kpi={{
    id: "revenue",
    label: "Total Omzet",
    value: 187500000,
    valueFormatted: "Rp 187,5 Jt",
    change: 12.4,
    changeLabel: "vs bulan lalu",
    icon: "revenue",
    color: "indigo",
    trend: "up",
  }}
  index={0}
/>
```

---

### `<ChartRenderer>`
**Location:** `components/ChartRenderer.js`

Dynamic chart renderer using Recharts. Supports Line, Bar, Pie, and Funnel charts.

**Props:**
- `reportId` (string): "penjualan-harian" | "top-produk" | "distribusi-kategori" | "performa-sales" | "funnel-order"

**Usage:**
```jsx
<ChartRenderer reportId="penjualan-harian" />
```

**Chart Types:**
- **Line:** Dual Y-axis (omzet + order count)
- **Bar:** Horizontal or vertical bars
- **Pie:** Donut chart with percentage labels
- **Funnel:** Order conversion funnel

---

### `<AIInsightCard>`
**Location:** `components/AIInsightCard.js`

Display AI-generated insights with gradient background and glow effect.

**Props:**
- `insight` (object):
  - `summary` (string): AI-generated text
  - `chart_type` (string): Recommended chart type
  - `anomaly` (boolean): Whether anomaly detected
- `loading` (boolean): Show skeleton loader

**Usage:**
```jsx
<AIInsightCard
  insight={{
    summary: "Penjualan bulan April menunjukkan tren positif...",
    chart_type: "line",
    anomaly: true,
  }}
/>
```

---

### `<AnomalyFlag>`
**Location:** `components/AnomalyFlag.js`

Anomaly alert card with severity indicator and AI recommendation.

**Props:**
- `anomali` (object):
  - `id` (string)
  - `severity` (string): "high" | "medium"
  - `metrik` (string): Metric name
  - `tanggal` (string): Date
  - `nilai_aktual` (string): Actual value
  - `nilai_normal` (string): Expected value
  - `persen` (number): Deviation percentage
  - `rekomendasi` (string): AI recommendation

**Usage:**
```jsx
<AnomalyFlag
  anomali={{
    id: "a1",
    severity: "high",
    metrik: "Omzet Harian",
    tanggal: "13 Apr 2026",
    nilai_aktual: "Rp 4,1 Jt",
    nilai_normal: "Rp 8,2 Jt",
    persen: -50,
    rekomendasi: "Periksa apakah ada gangguan sistem...",
  }}
/>
```

---

### `<ReportSelector>`
**Location:** `components/ReportSelector.js`

Dropdown selector for report types with descriptions.

**Props:**
- `value` (string): Selected report ID
- `onChange` (function): Callback with new report ID

**Usage:**
```jsx
<ReportSelector
  value={selectedReport}
  onChange={(id) => setSelectedReport(id)}
/>
```

---

### `<FilterBar>`
**Location:** `components/FilterBar.js`

Date range filter with quick period buttons.

**Props:**
- `period` (string): Selected period ("7d" | "30d" | "3m" | "ytd")
- `onPeriodChange` (function): Callback with new period

**Usage:**
```jsx
<FilterBar
  period={period}
  onPeriodChange={(p) => setPeriod(p)}
/>
```

---

### `<TopProdukTable>`
**Location:** `components/TopProdukTable.js`

Table displaying top products with stock indicators and progress bars.

**Usage:**
```jsx
<TopProdukTable />
```

**Features:**
- Stock level badges (Kritis / Rendah / Aman)
- Omzet progress bars
- Hover animations
- Responsive design

---

## 🗂️ Admin Components

### `<DataTable>`
**Location:** `components/DataTable.js`

Generic data table with search, pagination, and custom column renderers.

**Props:**
- `columns` (array): Column definitions
  - `key` (string): Data key
  - `label` (string): Column header
  - `render` (function): Custom cell renderer (optional)
  - `align` (string): "left" | "center" | "right" (optional)
- `data` (array): Row data
- `searchKeys` (array): Keys to search against
- `emptyMessage` (string): Message when no data

**Usage:**
```jsx
<DataTable
  columns={[
    {
      key: "nama",
      label: "Nama Produk",
      render: (v) => <span className="font-medium">{v}</span>,
    },
    {
      key: "harga",
      label: "Harga",
      align: "right",
      render: (v) => formatRupiah(v),
    },
  ]}
  data={products}
  searchKeys={["nama", "kategori"]}
  emptyMessage="Belum ada produk."
/>
```

**Features:**
- Built-in search (case-insensitive)
- Pagination (10 items per page)
- Responsive design
- Hover states

---

### `<Modal>`
**Location:** `components/Modal.js`

Reusable modal dialog with backdrop and ESC key support.

**Props:**
- `open` (boolean): Visibility state
- `onClose` (function): Close callback
- `title` (string): Modal title
- `children` (ReactNode): Modal content
- `size` (string): "sm" | "md" | "lg" (default: "md")

**Usage:**
```jsx
<Modal
  open={isOpen}
  onClose={() => setIsOpen(false)}
  title="Edit Produk"
  size="lg"
>
  <form onSubmit={handleSave}>
    {/* Form fields */}
  </form>
</Modal>
```

**Features:**
- Backdrop click to close
- ESC key to close
- Smooth fade-in animation
- Responsive sizing

---

## 🎨 Styling Conventions

### Color Classes
```jsx
// Primary actions
className="bg-indigo-600 hover:bg-indigo-500 text-white"

// Secondary actions
className="border border-slate-700 text-slate-400 hover:text-white"

// Success states
className="text-emerald-400 bg-emerald-500/10 border-emerald-500/20"

// Warning states
className="text-amber-400 bg-amber-500/10 border-amber-500/20"

// Error states
className="text-rose-400 bg-rose-500/10 border-rose-500/30"
```

### Border Radius
- **Small:** `rounded-lg` (8px)
- **Medium:** `rounded-xl` (12px)
- **Large:** `rounded-2xl` (16px)
- **Full:** `rounded-full`

### Spacing
- **Tight:** `gap-1` `p-1` (4px)
- **Normal:** `gap-3` `p-3` (12px)
- **Loose:** `gap-6` `p-6` (24px)

### Typography
- **Heading:** `text-base font-semibold text-white`
- **Body:** `text-sm text-slate-300`
- **Caption:** `text-xs text-slate-500`
- **Label:** `text-[11px] font-semibold text-slate-400 uppercase tracking-wider`

---

## 🔧 Utility Functions

### `formatRupiah(n)`
**Location:** `lib/api.js`

Format number as Indonesian Rupiah.

```javascript
formatRupiah(187500000) // "Rp 187.500.000"
```

### `fromPeriod(period)`
**Location:** `lib/api.js`

Convert period string to date range.

```javascript
fromPeriod("30d") // { from: "2026-03-28", to: "2026-04-27" }
```

### `getUser()`
**Location:** `lib/auth.js`

Get current user from localStorage.

```javascript
const user = getUser();
console.log(user.name, user.role);
```

### `isAuthenticated()`
**Location:** `lib/auth.js`

Check if user is logged in.

```javascript
if (!isAuthenticated()) {
  router.push("/login");
}
```

---

## 📦 Icon Library

Using **Lucide React** for all icons:

```jsx
import {
  Home, Package, Users, Settings, LogOut,
  Plus, Pencil, Trash2, Search, Filter,
  TrendingUp, TrendingDown, AlertCircle,
  CheckCircle, Loader2, X, ChevronRight,
} from "lucide-react";

<TrendingUp size={16} className="text-emerald-400" />
```

**Common sizes:**
- `size={12}` — Small icons (badges, inline)
- `size={16}` — Default icons (buttons, cards)
- `size={20}` — Medium icons (headers)
- `size={28}` — Large icons (loaders, empty states)

---

## 🎯 Best Practices

1. **Always wrap admin pages with `<ProtectedRoute>`**
2. **Use `<DataTable>` for list views** (consistent UX)
3. **Use `<Modal>` for forms** (better focus management)
4. **Show loading states** (Loader2 icon + text)
5. **Show error states** (AlertCircle + message)
6. **Use semantic HTML** (button, form, label)
7. **Add aria-labels** for accessibility
8. **Keep components small** (<300 lines)
9. **Extract repeated logic** into custom hooks
10. **Use TypeScript** (future improvement)

---

**Last updated:** April 27, 2026
