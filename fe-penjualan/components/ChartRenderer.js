"use client";

import {
  LineChart, Line, BarChart, Bar, PieChart, Pie, Cell,
  XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer,
  Legend, FunnelChart, Funnel, LabelList,
} from "recharts";

import {
  dailySalesData, topProdukData, kategoriData,
  funnelOrderData, salesPerformData,
} from "@/lib/dummyData";

// ---- Formatters ----
const formatRupiah = (v) => {
  if (v >= 1_000_000) return `${(v / 1_000_000).toFixed(1)} Jt`;
  if (v >= 1_000)     return `${(v / 1_000).toFixed(0)} Rb`;
  return v;
};

const tooltipStyle = {
  backgroundColor: "#1e293b",
  border: "1px solid #334155",
  borderRadius: "12px",
  color: "#f1f5f9",
  fontSize: "12px",
  padding: "8px 12px",
};

const COLORS = ["#6366f1", "#8b5cf6", "#a78bfa", "#7c3aed", "#c084fc"];

// ---- Chart: Line — Penjualan Harian ----
function LineChartPenjualan() {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={dailySalesData} margin={{ top: 5, right: 10, left: 0, bottom: 5 }}>
        <CartesianGrid strokeDasharray="3 3" stroke="#334155" vertical={false} />
        <XAxis
          dataKey="tanggal"
          tick={{ fill: "#64748b", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          interval={4}
        />
        <YAxis
          yAxisId="left"
          tickFormatter={formatRupiah}
          tick={{ fill: "#64748b", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          width={52}
        />
        <YAxis
          yAxisId="right"
          orientation="right"
          tick={{ fill: "#64748b", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          width={30}
        />
        <Tooltip
          contentStyle={tooltipStyle}
          formatter={(value, name) => {
            if (name === "omzet") return [`Rp ${formatRupiah(value)}`, "Omzet"];
            return [value, "Order"];
          }}
          labelStyle={{ color: "#94a3b8", marginBottom: 4 }}
        />
        <Legend
          wrapperStyle={{ fontSize: 12, color: "#94a3b8", paddingTop: 12 }}
          formatter={(v) => v === "omzet" ? "Omzet" : "Jumlah Order"}
        />
        <Line
          yAxisId="left"
          type="monotone"
          dataKey="omzet"
          stroke="#6366f1"
          strokeWidth={2.5}
          dot={false}
          activeDot={{ r: 5, fill: "#6366f1", strokeWidth: 0 }}
        />
        <Line
          yAxisId="right"
          type="monotone"
          dataKey="order"
          stroke="#a78bfa"
          strokeWidth={2}
          dot={false}
          strokeDasharray="4 2"
          activeDot={{ r: 4, fill: "#a78bfa", strokeWidth: 0 }}
        />
      </LineChart>
    </ResponsiveContainer>
  );
}

// ---- Chart: Bar — Top Produk ----
function BarChartProduk() {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart data={topProdukData} layout="vertical" margin={{ top: 5, right: 10, left: 90, bottom: 5 }}>
        <CartesianGrid strokeDasharray="3 3" stroke="#334155" horizontal={false} />
        <XAxis
          type="number"
          tickFormatter={formatRupiah}
          tick={{ fill: "#64748b", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
        />
        <YAxis
          type="category"
          dataKey="nama"
          tick={{ fill: "#94a3b8", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          width={88}
          tickFormatter={(v) => v.length > 18 ? v.slice(0, 16) + "…" : v}
        />
        <Tooltip
          contentStyle={tooltipStyle}
          formatter={(v, name) => {
            if (name === "omzet") return [`Rp ${(v / 1_000_000).toFixed(2)} Jt`, "Omzet"];
            return [v, "Terjual (unit)"];
          }}
        />
        <Bar dataKey="omzet" fill="#6366f1" radius={[0, 6, 6, 0]} />
        <Bar dataKey="terjual" fill="#a78bfa" radius={[0, 6, 6, 0]} />
      </BarChart>
    </ResponsiveContainer>
  );
}

// ---- Chart: Pie — Kategori ----
const RADIAN = Math.PI / 180;
const renderCustomLabel = ({ cx, cy, midAngle, innerRadius, outerRadius, percent }) => {
  const r = innerRadius + (outerRadius - innerRadius) * 0.5;
  const x = cx + r * Math.cos(-midAngle * RADIAN);
  const y = cy + r * Math.sin(-midAngle * RADIAN);
  if (percent < 0.06) return null;
  return (
    <text x={x} y={y} fill="white" textAnchor="middle" dominantBaseline="central" fontSize={11} fontWeight={600}>
      {`${(percent * 100).toFixed(0)}%`}
    </text>
  );
};

function PieChartKategori() {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <PieChart>
        <Pie
          data={kategoriData}
          cx="50%"
          cy="50%"
          innerRadius={70}
          outerRadius={110}
          paddingAngle={3}
          dataKey="value"
          labelLine={false}
          label={renderCustomLabel}
        >
          {kategoriData.map((entry, i) => (
            <Cell key={i} fill={entry.fill} stroke="transparent" />
          ))}
        </Pie>
        <Tooltip
          contentStyle={tooltipStyle}
          formatter={(v, name, props) => [`${v} unit`, props.payload.name]}
        />
        <Legend
          iconType="circle"
          iconSize={8}
          wrapperStyle={{ fontSize: 12, color: "#94a3b8" }}
        />
      </PieChart>
    </ResponsiveContainer>
  );
}

// ---- Chart: Bar — Performa Sales ----
function BarChartSales() {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart data={salesPerformData} margin={{ top: 5, right: 10, left: 10, bottom: 5 }}>
        <CartesianGrid strokeDasharray="3 3" stroke="#334155" vertical={false} />
        <XAxis
          dataKey="nama"
          tick={{ fill: "#94a3b8", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          tickFormatter={(v) => v.split(" ")[0]}
        />
        <YAxis
          tickFormatter={formatRupiah}
          tick={{ fill: "#64748b", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          width={52}
        />
        <Tooltip
          contentStyle={tooltipStyle}
          formatter={(v, name) => [`Rp ${(v / 1_000_000).toFixed(1)} Jt`, name === "omzet" ? "Realisasi" : "Target"]}
        />
        <Legend wrapperStyle={{ fontSize: 12, color: "#94a3b8", paddingTop: 12 }} />
        <Bar dataKey="omzet"  fill="#6366f1" radius={[6, 6, 0, 0]} name="Realisasi" />
        <Bar dataKey="target" fill="#334155" radius={[6, 6, 0, 0]} name="Target" />
      </BarChart>
    </ResponsiveContainer>
  );
}

// ---- Chart: Funnel — Order ----
function FunnelChartOrder() {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <FunnelChart margin={{ top: 5, right: 10, left: 10, bottom: 5 }}>
        <Tooltip contentStyle={tooltipStyle} formatter={(v, name, props) => [v, props.payload.status]} />
        <Funnel dataKey="jumlah" data={funnelOrderData} isAnimationActive>
          {funnelOrderData.map((entry, i) => (
            <Cell key={i} fill={entry.fill} />
          ))}
          <LabelList
            position="center"
            fill="#fff"
            stroke="none"
            fontSize={12}
            fontWeight={600}
            formatter={(v, name, props) => `${props?.payload?.status}: ${v}`}
          />
        </Funnel>
      </FunnelChart>
    </ResponsiveContainer>
  );
}

// ---- Main Chart Renderer ----
export default function ChartRenderer({ reportId }) {
  const chartMap = {
    "penjualan-harian":     <LineChartPenjualan />,
    "top-produk":           <BarChartProduk />,
    "distribusi-kategori":  <PieChartKategori />,
    "performa-sales":       <BarChartSales />,
    "funnel-order":         <FunnelChartOrder />,
  };

  return (
    <div className="w-full animate-fade-in">
      {chartMap[reportId] ?? (
        <div className="flex items-center justify-center h-[300px] text-slate-500 text-sm">
          Pilih laporan untuk menampilkan chart.
        </div>
      )}
    </div>
  );
}
