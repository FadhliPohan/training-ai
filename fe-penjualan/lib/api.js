const BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3032";

// ─── Core fetch wrapper ────────────────────────────────────────────────────
async function apiFetch(path, options = {}) {
  const token = typeof window !== "undefined" ? localStorage.getItem("token") : null;

  const res = await fetch(`${BASE}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options.headers,
    },
    ...options,
  });

  const json = await res.json().catch(() => ({}));
  if (!res.ok) {
    const msg = json?.message || `Error ${res.status}`;
    throw new Error(msg);
  }
  return json;
}

// ─── Auth ──────────────────────────────────────────────────────────────────
export const authAPI = {
  login: (email, password) =>
    apiFetch("/api/v1/auth/login", {
      method: "POST",
      body: JSON.stringify({ email, password }),
    }),
  logout: () =>
    apiFetch("/api/v1/auth/logout", { method: "POST" }),
  me: () => apiFetch("/api/v1/auth/me"),
};

// ─── Reports ───────────────────────────────────────────────────────────────
export const reportAPI = {
  get: ({ type = "daily-sales", from, to, salesId, mode = "ai" } = {}) => {
    const qs = new URLSearchParams({ type, mode });
    if (from) qs.set("from", from);
    if (to) qs.set("to", to);
    if (salesId) qs.set("sales_id", salesId);
    return apiFetch(`/api/v1/reports?${qs}`);
  },
};

// ─── Users ─────────────────────────────────────────────────────────────────
export const usersAPI = {
  list: () => apiFetch("/api/v1/users"),
  create: (data) =>
    apiFetch("/api/v1/users", { method: "POST", body: JSON.stringify(data) }),
  update: (id, data) =>
    apiFetch(`/api/v1/users/${id}`, { method: "PUT", body: JSON.stringify(data) }),
  deactivate: (id) =>
    apiFetch(`/api/v1/users/${id}`, { method: "PATCH" }), // BUG-006 fix: PATCH not DELETE
};

// ─── Products ──────────────────────────────────────────────────────────────
export const produkAPI = {
  list: () => apiFetch("/api/v1/produk"),
  getById: (id) => apiFetch(`/api/v1/produk/${id}`),
  create: (data) =>
    apiFetch("/api/v1/produk", { method: "POST", body: JSON.stringify(data) }),
  update: (id, data) =>
    apiFetch(`/api/v1/produk/${id}`, { method: "PUT", body: JSON.stringify(data) }),
  deactivate: (id) =>
    apiFetch(`/api/v1/produk/${id}`, { method: "PATCH" }), // BUG-007 fix: PATCH for toggle aktif
};

// ─── Customers ─────────────────────────────────────────────────────────────
export const customerAPI = {
  list: () => apiFetch("/api/v1/customer"),
  getById: (id) => apiFetch(`/api/v1/customer/${id}`),
  create: (data) =>
    apiFetch("/api/v1/customer", { method: "POST", body: JSON.stringify(data) }),
  update: (id, data) =>
    apiFetch(`/api/v1/customer/${id}`, { method: "PUT", body: JSON.stringify(data) }),
};

// ─── Settings ──────────────────────────────────────────────────────────────
export const settingsAPI = {
  getTelegram: () => apiFetch("/api/v1/settings/telegram"),
  updateTelegram: (data) =>
    apiFetch("/api/v1/settings/telegram", {
      method: "PUT",
      body: JSON.stringify(data),
    }),
};

// ─── Helpers ───────────────────────────────────────────────────────────────
export function formatRupiah(n) {
  return "Rp " + Number(n).toLocaleString("id-ID");
}

export function fromPeriod(period) {
  const now = new Date();
  const to = now.toISOString().split("T")[0];
  const d = new Date(now);
  if (period === "7d") d.setDate(d.getDate() - 7);
  if (period === "30d") d.setDate(d.getDate() - 30);
  if (period === "90d") d.setDate(d.getDate() - 90);
  const from = d.toISOString().split("T")[0];
  return { from, to };
}
