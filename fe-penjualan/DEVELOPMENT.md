# Development Guide

> **InsightFlow Frontend Development**  
> Guidelines for contributing to the codebase

---

## 🚀 Getting Started

### Prerequisites
- Node.js 18+ and npm
- Backend API running on `http://localhost:8080`
- Code editor (VS Code recommended)

### Initial Setup
```bash
# Clone and navigate
cd fe-penjualan

# Install dependencies
npm install

# Create environment file
cp .env.example .env.local

# Edit .env.local
NEXT_PUBLIC_API_URL=http://localhost:8080

# Start dev server
npm run dev
```

Visit `http://localhost:3000`

---

## 📁 File Organization

### Adding a New Page

1. **Create page file:**
   ```bash
   fe-penjualan/app/admin/orders/page.js
   ```

2. **Use the template:**
   ```jsx
   "use client";
   
   import ProtectedRoute from "@/components/ProtectedRoute";
   import Sidebar from "@/components/Sidebar";
   import Topbar from "@/components/Topbar";
   
   export default function OrdersPage() {
     return (
       <ProtectedRoute>
         <OrdersContent />
       </ProtectedRoute>
     );
   }
   
   function OrdersContent() {
     return (
       <div className="flex min-h-screen bg-slate-950">
         <Sidebar />
         <div className="flex-1 flex flex-col min-w-0 lg:ml-[72px]">
           <Topbar title="Manajemen Orders" />
           <main className="flex-1 p-6">
             {/* Your content */}
           </main>
         </div>
       </div>
     );
   }
   ```

3. **Add to Sidebar navigation:**
   ```jsx
   // components/Sidebar.js
   const navItems = [
     // ...existing items
     { icon: ShoppingCart, label: "Orders", href: "/admin/orders" },
   ];
   ```

---

### Adding a New Component

1. **Create component file:**
   ```bash
   fe-penjualan/components/OrderCard.js
   ```

2. **Component structure:**
   ```jsx
   "use client";
   
   /**
    * OrderCard — Display order summary
    * 
    * Props:
    *  - order (object): Order data
    *  - onEdit (function): Edit callback
    */
   export default function OrderCard({ order, onEdit }) {
     return (
       <div className="rounded-xl border border-slate-700/50 bg-[#1e293b]/60 p-4">
         {/* Component content */}
       </div>
     );
   }
   ```

3. **Export from index (optional):**
   ```jsx
   // components/index.js
   export { default as OrderCard } from "./OrderCard";
   ```

---

### Adding a New API Endpoint

1. **Add to `lib/api.js`:**
   ```javascript
   export const ordersAPI = {
     list: () => apiFetch("/api/v1/orders"),
     getById: (id) => apiFetch(`/api/v1/orders/${id}`),
     create: (data) =>
       apiFetch("/api/v1/orders", { method: "POST", body: JSON.stringify(data) }),
     update: (id, data) =>
       apiFetch(`/api/v1/orders/${id}`, { method: "PUT", body: JSON.stringify(data) }),
     confirm: (id) =>
       apiFetch(`/api/v1/orders/${id}/confirm`, { method: "POST" }),
   };
   ```

2. **Use in component:**
   ```jsx
   import { ordersAPI } from "@/lib/api";
   
   const fetchOrders = async () => {
     try {
       const res = await ordersAPI.list();
       setOrders(res.data ?? res);
     } catch (e) {
       setError(e.message);
     }
   };
   ```

---

## 🎨 Styling Guidelines

### Tailwind Classes Order
Follow this order for consistency:

1. **Layout:** `flex`, `grid`, `block`, `inline-flex`
2. **Positioning:** `relative`, `absolute`, `fixed`, `sticky`
3. **Sizing:** `w-*`, `h-*`, `min-w-*`, `max-w-*`
4. **Spacing:** `p-*`, `m-*`, `gap-*`
5. **Typography:** `text-*`, `font-*`, `leading-*`, `tracking-*`
6. **Colors:** `bg-*`, `text-*`, `border-*`
7. **Borders:** `border`, `border-*`, `rounded-*`
8. **Effects:** `shadow-*`, `opacity-*`, `blur-*`
9. **Transitions:** `transition-*`, `duration-*`, `ease-*`
10. **States:** `hover:*`, `focus:*`, `active:*`, `disabled:*`

**Example:**
```jsx
<button className="
  flex items-center gap-2
  px-4 py-2.5
  text-sm font-semibold
  bg-indigo-600 text-white
  border border-transparent rounded-xl
  shadow-lg shadow-indigo-600/30
  transition-all duration-200
  hover:bg-indigo-500
  active:bg-indigo-700
  disabled:opacity-60
">
  Save
</button>
```

---

### Color Palette

```javascript
// Primary
indigo-400  #818cf8
indigo-500  #6366f1
indigo-600  #4f46e5

// Background
slate-950   #0f172a  (body)
slate-900   #1e293b  (cards)
slate-800   #334155  (inputs)
slate-700   #475569  (borders)

// Text
slate-100   #f1f5f9  (primary)
slate-300   #cbd5e1  (secondary)
slate-500   #64748b  (muted)

// Status
emerald-400 #34d399  (success)
amber-400   #fbbf24  (warning)
rose-400    #fb7185  (error)
```

---

## 🧪 Testing Checklist

Before committing, verify:

- [ ] **Responsive:** Test on mobile (375px), tablet (768px), desktop (1440px)
- [ ] **Dark theme:** All colors readable on dark background
- [ ] **Loading states:** Show Loader2 during async operations
- [ ] **Error states:** Display error messages clearly
- [ ] **Empty states:** Show helpful message when no data
- [ ] **Form validation:** Required fields marked, validation errors shown
- [ ] **Accessibility:** Keyboard navigation works, labels present
- [ ] **Performance:** No unnecessary re-renders, debounce search inputs

---

## 🔍 Debugging Tips

### API Errors
```javascript
// Add detailed logging
try {
  const res = await produkAPI.list();
  console.log("✅ API Response:", res);
} catch (e) {
  console.error("❌ API Error:", e.message);
  console.error("Stack:", e.stack);
}
```

### State Issues
```javascript
// Use React DevTools
// Install: https://react.dev/learn/react-developer-tools

// Add debug logs
useEffect(() => {
  console.log("State changed:", { data, loading, error });
}, [data, loading, error]);
```

### Styling Issues
```javascript
// Use Tailwind Play
// https://play.tailwindcss.com/

// Check computed styles in DevTools
// Right-click element → Inspect → Computed tab
```

---

## 📦 Build & Deploy

### Production Build
```bash
# Build
npm run build

# Test production build locally
npm start

# Check bundle size
npm run build -- --analyze
```

### Environment Variables
```bash
# .env.local (development)
NEXT_PUBLIC_API_URL=http://localhost:8080

# .env.production (production)
NEXT_PUBLIC_API_URL=https://api.insightflow.id
```

### Docker Deployment
```dockerfile
# Dockerfile
FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build
EXPOSE 3000
CMD ["npm", "start"]
```

```bash
# Build and run
docker build -t insightflow-fe .
docker run -p 3000:3000 -e NEXT_PUBLIC_API_URL=http://api:8080 insightflow-fe
```

---

## 🐛 Common Issues

### Issue: "Module not found"
**Solution:** Clear cache and reinstall
```bash
rm -rf .next node_modules package-lock.json
npm install
```

### Issue: "Hydration mismatch"
**Solution:** Ensure server and client render the same HTML
```jsx
// ❌ Bad: Using Date.now() directly
<p>{Date.now()}</p>

// ✅ Good: Use useEffect for client-only values
const [time, setTime] = useState(null);
useEffect(() => setTime(Date.now()), []);
return <p>{time ?? "Loading..."}</p>;
```

### Issue: "localStorage is not defined"
**Solution:** Check if running on client
```jsx
// ❌ Bad: Direct access
const token = localStorage.getItem("token");

// ✅ Good: Check environment
const token = typeof window !== "undefined"
  ? localStorage.getItem("token")
  : null;
```

### Issue: API CORS errors
**Solution:** Configure backend CORS
```go
// be-penjualan/cmd/main.go
app.Use(cors.New(cors.Config{
  AllowOrigins: "http://localhost:3000",
  AllowHeaders: "Origin, Content-Type, Accept, Authorization",
}))
```

---

## 📚 Resources

- **Next.js Docs:** https://nextjs.org/docs
- **Tailwind CSS:** https://tailwindcss.com/docs
- **Recharts:** https://recharts.org/en-US/
- **Lucide Icons:** https://lucide.dev/icons/
- **React Patterns:** https://react.dev/learn

---

## 🤝 Contributing

1. **Create feature branch:** `git checkout -b feature/order-management`
2. **Make changes:** Follow guidelines above
3. **Test thoroughly:** Check all scenarios
4. **Commit:** `git commit -m "feat: add order management page"`
5. **Push:** `git push origin feature/order-management`
6. **Create PR:** Describe changes and attach screenshots

### Commit Message Format
```
<type>: <description>

[optional body]
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting, no code change
- `refactor`: Code restructuring
- `test`: Adding tests
- `chore`: Maintenance

**Examples:**
```
feat: add order management page
fix: resolve pagination bug in DataTable
docs: update component documentation
style: format code with prettier
refactor: extract form validation logic
```

---

**Happy coding! 🚀**
