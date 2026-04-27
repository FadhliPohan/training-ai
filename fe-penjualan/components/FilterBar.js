"use client";

import { Calendar } from "lucide-react";

const PERIODS = [
  { id: "7d", label: "7H", full: "7 Hari Terakhir" },
  { id: "30d", label: "30H", full: "30 Hari Terakhir" },
  { id: "3m", label: "3B", full: "3 Bulan Terakhir" },
  { id: "ytd", label: "YTD", full: "Tahun Ini (YTD)" },
];

export default function FilterBar({ period, onPeriodChange }) {
  return (
    <div className="flex items-center gap-2">
      {/* Calendar icon — decorative label */}
      <Calendar size={14} className="text-indigo-400 flex-shrink-0" />

      {/* Segmented pill group — compact, no overflow */}
      <div className="flex items-center bg-slate-800/60 border border-slate-700 rounded-xl p-1 gap-0.5">
        {PERIODS.map((p) => (
          <button
            key={p.id}
            onClick={() => onPeriodChange(p.id)}
            title={p.full}
            className={`
              px-3 py-1.5 rounded-lg text-xs font-semibold transition-all whitespace-nowrap
              ${p.id === period
                ? "bg-indigo-600 text-white shadow-sm shadow-indigo-500/30"
                : "text-slate-400 hover:text-slate-200 hover:bg-slate-700/60"
              }
            `}
          >
            {p.label}
          </button>
        ))}
      </div>
    </div>
  );
}
