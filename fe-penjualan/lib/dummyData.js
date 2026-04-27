// ============================================================
// DUMMY DATA — InsightFlow Dashboard
// Semua data bersifat dummy untuk prototype dashboard.
// Akan diganti dengan API call ke backend Golang nanti.
// ============================================================

// ---- KPI Metrics ----
export const kpiData = [
  {
    id: "revenue",
    label: "Total Omzet",
    value: 187_500_000,
    valueFormatted: "Rp 187,5 Jt",
    change: 12.4,
    changeLabel: "vs bulan lalu",
    icon: "revenue",
    color: "indigo",
    trend: "up",
  },
  {
    id: "orders",
    label: "Total Order",
    value: 342,
    valueFormatted: "342",
    change: 8.7,
    changeLabel: "vs bulan lalu",
    icon: "orders",
    color: "violet",
    trend: "up",
  },
  {
    id: "customers",
    label: "Customer Aktif",
    value: 128,
    valueFormatted: "128",
    change: -3.2,
    changeLabel: "vs bulan lalu",
    icon: "customers",
    color: "purple",
    trend: "down",
  },
  {
    id: "avg_order",
    label: "Rata-rata Order",
    value: 548_000,
    valueFormatted: "Rp 548 Rb",
    change: 5.1,
    changeLabel: "vs bulan lalu",
    icon: "avgorder",
    color: "sky",
    trend: "up",
  },
];

// ---- Penjualan Harian (30 hari terakhir) ----
export const dailySalesData = [
  { tanggal: "28 Mar", omzet: 4_200_000, order: 8 },
  { tanggal: "29 Mar", omzet: 3_800_000, order: 7 },
  { tanggal: "30 Mar", omzet: 5_100_000, order: 10 },
  { tanggal: "31 Mar", omzet: 6_300_000, order: 12 },
  { tanggal: "01 Apr", omzet: 7_200_000, order: 14 },
  { tanggal: "02 Apr", omzet: 5_800_000, order: 11 },
  { tanggal: "03 Apr", omzet: 4_500_000, order: 9 },
  { tanggal: "04 Apr", omzet: 8_100_000, order: 16 },
  { tanggal: "05 Apr", omzet: 9_300_000, order: 18 },
  { tanggal: "06 Apr", omzet: 7_600_000, order: 15 },
  { tanggal: "07 Apr", omzet: 6_200_000, order: 12 },
  { tanggal: "08 Apr", omzet: 5_400_000, order: 10 },
  { tanggal: "09 Apr", omzet: 8_700_000, order: 17 },
  { tanggal: "10 Apr", omzet: 11_200_000, order: 22 },
  { tanggal: "11 Apr", omzet: 9_800_000, order: 19 },
  { tanggal: "12 Apr", omzet: 7_300_000, order: 14 },
  { tanggal: "13 Apr", omzet: 4_100_000, order: 8 },  // anomali turun
  { tanggal: "14 Apr", omzet: 6_500_000, order: 13 },
  { tanggal: "15 Apr", omzet: 10_200_000, order: 20 },
  { tanggal: "16 Apr", omzet: 12_400_000, order: 24 },
  { tanggal: "17 Apr", omzet: 11_100_000, order: 22 },
  { tanggal: "18 Apr", omzet: 9_600_000, order: 18 },
  { tanggal: "19 Apr", omzet: 7_800_000, order: 15 },
  { tanggal: "20 Apr", omzet: 8_200_000, order: 16 },
  { tanggal: "21 Apr", omzet: 13_500_000, order: 27 },
  { tanggal: "22 Apr", omzet: 14_100_000, order: 28 },
  { tanggal: "23 Apr", omzet: 12_700_000, order: 25 },
  { tanggal: "24 Apr", omzet: 10_400_000, order: 20 },
  { tanggal: "25 Apr", omzet: 11_800_000, order: 23 },
  { tanggal: "26 Apr", omzet: 9_200_000, order: 18 },
];

// ---- Top Produk ----
export const topProdukData = [
  { nama: "Kemeja Batik Lengan Panjang", kategori: "Kemeja",  terjual: 87, omzet: 16_095_000, stok: 13 },
  { nama: "Kaos Polos Premium",          kategori: "Kaos",    terjual: 143, omzet: 10_725_000, stok: 57 },
  { nama: "Celana Chino Slim Fit",       kategori: "Celana",  terjual: 62, omzet: 13_640_000, stok: 8 },
  { nama: "Jaket Bomber Unisex",         kategori: "Jaket",   terjual: 34, omzet: 11_900_000, stok: 6 },
  { nama: "Dress Batik Sogan",           kategori: "Dress",   terjual: 21, omzet:  9_450_000, stok: 4 },
  { nama: "Kemeja Flanel Kotak",         kategori: "Kemeja",  terjual: 55, omzet:  9_350_000, stok: 25 },
  { nama: "Hoodie Oversize Fleece",      kategori: "Jaket",   terjual: 41, omzet:  8_610_000, stok: 9 },
];

// ---- Distribusi per Kategori ----
export const kategoriData = [
  { name: "Kemeja",  value: 142, fill: "#6366f1" },
  { name: "Kaos",    value: 143, fill: "#8b5cf6" },
  { name: "Celana",  value: 62,  fill: "#a78bfa" },
  { name: "Jaket",   value: 75,  fill: "#7c3aed" },
  { name: "Dress",   value: 21,  fill: "#c084fc" },
];

// ---- Funnel Status Order ----
export const funnelOrderData = [
  { status: "Order Masuk",  jumlah: 420, fill: "#6366f1" },
  { status: "Dikonfirmasi", jumlah: 395, fill: "#7c3aed" },
  { status: "Dibayar",      jumlah: 371, fill: "#8b5cf6" },
  { status: "Dikirim",      jumlah: 358, fill: "#a78bfa" },
  { status: "Selesai",      jumlah: 342, fill: "#c084fc" },
];

// ---- Performa Per Sales ----
export const salesPerformData = [
  { nama: "Citra Dewi",    order: 98,  omzet: 54_200_000, target: 50_000_000 },
  { nama: "Budi Santoso",  order: 87,  omzet: 48_100_000, target: 50_000_000 },
  { nama: "Rina Pertiwi",  order: 73,  omzet: 39_600_000, target: 40_000_000 },
  { nama: "Denny Halim",   order: 62,  omzet: 32_400_000, target: 40_000_000 },
  { nama: "Sari Andini",   order: 22,  omzet: 12_200_000, target: 30_000_000 },  // underperform
];

// ---- Anomali Terdeteksi ----
export const anomaliData = [
  {
    id: "a1",
    severity: "high",
    metrik: "Omzet Harian",
    tanggal: "13 Apr 2026",
    nilai_aktual: "Rp 4,1 Jt",
    nilai_normal: "Rp 8,2 Jt",
    persen: -50,
    rekomendasi: "Periksa apakah ada gangguan sistem atau hari libur yang tidak tercatat. Pertimbangkan kampanye flash sale untuk recovery.",
  },
  {
    id: "a2",
    severity: "medium",
    metrik: "Stok Celana Chino",
    tanggal: "26 Apr 2026",
    nilai_aktual: "8 unit",
    nilai_normal: "30 unit (buffer minimum)",
    persen: -73,
    rekomendasi: "Segera lakukan reorder Celana Chino Slim Fit. Stok kritis akan menghambat penjualan dalam 3-5 hari ke depan.",
  },
  {
    id: "a3",
    severity: "medium",
    metrik: "Performa Sales — Sari Andini",
    tanggal: "Apr 2026",
    nilai_aktual: "Rp 12,2 Jt",
    nilai_normal: "Target Rp 30 Jt",
    persen: -59,
    rekomendasi: "Sales berada di 40% target. Rekomendasikan coaching atau review territory untuk mengidentifikasi hambatan.",
  },
];

// ---- AI Insight (per laporan) ----
export const aiInsights = {
  "penjualan-harian": {
    summary:
      "Penjualan bulan April menunjukkan tren positif dengan pertumbuhan 12,4% dibanding Maret. Peak terjadi pada tanggal 21–22 April dengan omzet Rp 14,1 Jt/hari. Namun, sempat terjadi penurunan signifikan pada 13 April (−50% dari rata-rata) yang perlu ditelusuri penyebabnya.",
    chart_type: "line",
    anomaly: true,
  },
  "top-produk": {
    summary:
      "Kaos Polos Premium menjadi produk dengan volume terjual tertinggi (143 unit), sementara Kemeja Batik Lengan Panjang memimpin dari sisi omzet (Rp 16,1 Jt). Dress Batik Sogan dan Jaket Bomber memiliki stok kritis — perlu reorder segera untuk menghindari kehilangan peluang penjualan.",
    chart_type: "bar",
    anomaly: false,
  },
  "distribusi-kategori": {
    summary:
      "Kategori Kaos dan Kemeja mendominasi volume penjualan (masing-masing 143 dan 142 unit, total 58% dari semua penjualan). Kategori Dress memiliki kontribusi terkecil (4,8%) namun dengan nilai rata-rata transaksi tertinggi — segmen premium yang perlu diperkuat stoknya.",
    chart_type: "pie",
    anomaly: false,
  },
  "performa-sales": {
    summary:
      "Citra Dewi memimpin performa dengan omzet Rp 54,2 Jt (108% target). Budi Santoso mendekati target di 96%. Sari Andini hanya mencapai 41% target — membutuhkan perhatian dan coaching segera. Secara keseluruhan, 3 dari 5 sales berada di bawah atau tepat di batas target.",
    chart_type: "bar",
    anomaly: true,
  },
  "funnel-order": {
    summary:
      "Conversion rate dari order masuk ke selesai adalah 81,4% — di atas rata-rata industri fashion B2B (70-75%). Drop terbesar terjadi antara Order Masuk → Dikonfirmasi (6%), yang mengindikasikan ada delay konfirmasi oleh tim sales. Perlu SOP untuk respons konfirmasi maksimal 2 jam.",
    chart_type: "funnel",
    anomaly: false,
  },
};

// ---- Pilihan Laporan (Report Types) ----
export const reportOptions = [
  {
    id: "penjualan-harian",
    label: "📈 Tren Penjualan Harian",
    description: "Omzet dan jumlah order per hari selama periode pilihan",
    chartType: "line",
  },
  {
    id: "top-produk",
    label: "🏆 Top Produk Terlaris",
    description: "Ranking produk berdasarkan volume terjual dan omzet",
    chartType: "bar",
  },
  {
    id: "distribusi-kategori",
    label: "🥧 Distribusi Kategori Produk",
    description: "Persentase penjualan per kategori pakaian",
    chartType: "pie",
  },
  {
    id: "performa-sales",
    label: "👥 Performa Per Sales",
    description: "Perbandingan omzet realisasi vs target per sales",
    chartType: "bar",
  },
  {
    id: "funnel-order",
    label: "🔻 Funnel Status Order",
    description: "Alur order dari masuk hingga selesai",
    chartType: "funnel",
  },
];
