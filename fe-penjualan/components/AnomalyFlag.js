"use client";

import { AlertTriangle, AlertCircle, ChevronRight } from "lucide-react";

const severityConfig = {
  high: {
    icon: AlertTriangle,
    bg:   "bg-rose-500/10 border-rose-500/30",
    icon_class: "text-rose-400",
    badge: "bg-rose-500/20 text-rose-400",
    badge_label: "Kritis",
    bar: "bg-rose-500",
  },
  medium: {
    icon: AlertCircle,
    bg:   "bg-amber-500/10 border-amber-500/30",
    icon_class: "text-amber-400",
    badge: "bg-amber-500/20 text-amber-400",
    badge_label: "Perlu Perhatian",
    bar: "bg-amber-500",
  },
};

export default function AnomalyFlag({ anomali }) {
  const cfg = severityConfig[anomali.severity] || severityConfig.medium;
  const Icon = cfg.icon;

  return (
    <div className={`relative rounded-xl border p-4 ${cfg.bg} transition-all hover:scale-[1.01] duration-200`}>
      {/* Left accent bar */}
      <div className={`absolute left-0 top-3 bottom-3 w-1 rounded-r-full ${cfg.bar}`} />

      <div className="pl-3">
        {/* Header */}
        <div className="flex items-start justify-between gap-3">
          <div className="flex items-center gap-2">
            <Icon size={16} className={cfg.icon_class} />
            <span className="text-sm font-semibold text-white">{anomali.metrik}</span>
          </div>
          <span className={`text-[10px] font-semibold px-2 py-0.5 rounded-full flex-shrink-0 ${cfg.badge}`}>
            {cfg.badge_label}
          </span>
        </div>

        {/* Data */}
        <div className="mt-2 grid grid-cols-2 gap-3">
          <div>
            <p className="text-[10px] text-slate-500 uppercase tracking-wider">Nilai Aktual</p>
            <p className="text-sm font-bold text-white mt-0.5">{anomali.nilai_aktual}</p>
          </div>
          <div>
            <p className="text-[10px] text-slate-500 uppercase tracking-wider">Ekspektasi Normal</p>
            <p className="text-sm font-bold text-slate-300 mt-0.5">{anomali.nilai_normal}</p>
          </div>
        </div>

        {/* Deviation */}
        <div className="mt-2 flex items-center gap-2">
          <span className={`text-xs font-bold ${anomali.persen < 0 ? "text-rose-400" : "text-emerald-400"}`}>
            {anomali.persen > 0 ? "+" : ""}{anomali.persen}%
          </span>
          <span className="text-[10px] text-slate-500">deviasi pada {anomali.tanggal}</span>
        </div>

        {/* Recommendation */}
        <div className="mt-3 p-3 bg-slate-800/60 rounded-lg border border-slate-700/40">
          <p className="text-[10px] font-semibold text-slate-400 uppercase tracking-wider mb-1">💡 Rekomendasi AI</p>
          <p className="text-xs text-slate-300 leading-relaxed">{anomali.rekomendasi}</p>
        </div>
      </div>
    </div>
  );
}
