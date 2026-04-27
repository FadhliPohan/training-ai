"use client";

import { TrendingUp, TrendingDown, DollarSign, ShoppingBag, Users, BarChart3 } from "lucide-react";

const iconMap = {
  revenue:   DollarSign,
  orders:    ShoppingBag,
  customers: Users,
  avgorder:  BarChart3,
};

const colorMap = {
  indigo: {
    bg:    "from-indigo-500/20 to-indigo-600/5",
    icon:  "bg-indigo-500/20 text-indigo-300",
    glow:  "shadow-indigo-500/10",
    badge: "text-indigo-300",
  },
  violet: {
    bg:    "from-violet-500/20 to-violet-600/5",
    icon:  "bg-violet-500/20 text-violet-300",
    glow:  "shadow-violet-500/10",
    badge: "text-violet-300",
  },
  purple: {
    bg:    "from-purple-500/20 to-purple-600/5",
    icon:  "bg-purple-500/20 text-purple-300",
    glow:  "shadow-purple-500/10",
    badge: "text-purple-300",
  },
  sky: {
    bg:    "from-sky-500/20 to-sky-600/5",
    icon:  "bg-sky-500/20 text-sky-300",
    glow:  "shadow-sky-500/10",
    badge: "text-sky-300",
  },
};

export default function KPICard({ kpi, index = 0 }) {
  const Icon = iconMap[kpi.icon] || BarChart3;
  const colors = colorMap[kpi.color] || colorMap.indigo;
  const isPositive = kpi.trend === "up";

  return (
    <div
      className={`
        relative overflow-hidden rounded-2xl p-5 border border-slate-700/50
        bg-gradient-to-br ${colors.bg} bg-[#1e293b]
        shadow-lg ${colors.glow}
        hover:border-indigo-500/30 hover:shadow-xl transition-all duration-300
        animate-slide-up
      `}
      style={{ animationDelay: `${index * 80}ms` }}
    >
      {/* Background decoration */}
      <div className="absolute -top-4 -right-4 w-20 h-20 rounded-full bg-gradient-to-br from-white/5 to-transparent" />

      <div className="relative flex items-start justify-between">
        {/* Icon */}
        <div className={`p-2.5 rounded-xl ${colors.icon} backdrop-blur-sm`}>
          <Icon size={20} />
        </div>

        {/* Change badge */}
        <div className={`
          flex items-center gap-1 px-2 py-1 rounded-lg text-xs font-semibold
          ${isPositive ? "bg-emerald-500/15 text-emerald-400" : "bg-rose-500/15 text-rose-400"}
        `}>
          {isPositive ? <TrendingUp size={11} /> : <TrendingDown size={11} />}
          {Math.abs(kpi.change)}%
        </div>
      </div>

      {/* Value */}
      <div className="mt-4">
        <p className="text-2xl font-bold text-white tracking-tight">{kpi.valueFormatted}</p>
        <p className="text-sm text-slate-400 mt-1">{kpi.label}</p>
        <p className="text-[11px] text-slate-500 mt-1">{kpi.changeLabel}</p>
      </div>

      {/* Mini sparkline bar (decorative) */}
      <div className="mt-4 flex items-end gap-0.5 h-6">
        {[40, 65, 55, 80, 70, 90, 75, 85, 95, 100].map((h, i) => (
          <div
            key={i}
            className={`flex-1 rounded-sm ${isPositive ? "bg-indigo-500/30" : "bg-rose-500/30"}`}
            style={{ height: `${h}%` }}
          />
        ))}
      </div>
    </div>
  );
}
