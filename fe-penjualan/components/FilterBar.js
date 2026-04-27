"use client";

import { useState } from "react";
import { Calendar, ChevronDown } from "lucide-react";

const PERIODS = [
  { id: "7d",  label: "7 Hari Terakhir" },
  { id: "30d", label: "30 Hari Terakhir" },
  { id: "3m",  label: "3 Bulan Terakhir" },
  { id: "ytd", label: "Tahun Ini (YTD)" },
];

export default function FilterBar({ period, onPeriodChange }) {
  const [open, setOpen] = useState(false);
  const selected = PERIODS.find((p) => p.id === period) || PERIODS[1];

  return (
    <div className="flex flex-wrap items-center gap-3">
      {/* Period selector */}
      <div className="relative">
        <button
          onClick={() => setOpen(!open)}
          className="flex items-center gap-2 bg-slate-800/60 border border-slate-700 hover:border-indigo-500/40 rounded-xl px-3 py-2 text-sm text-slate-300 transition-all"
        >
          <Calendar size={14} className="text-indigo-400" />
          {selected.label}
          <ChevronDown size={14} className={`text-slate-500 transition-transform ${open ? "rotate-180" : ""}`} />
        </button>

        {open && (
          <div className="absolute top-full left-0 mt-2 w-52 z-50 bg-[#1e293b] border border-slate-700 rounded-xl shadow-xl overflow-hidden animate-fade-in">
            {PERIODS.map((p) => (
              <button
                key={p.id}
                onClick={() => { onPeriodChange(p.id); setOpen(false); }}
                className={`
                  w-full text-left px-4 py-2.5 text-sm border-b border-slate-800 last:border-0 transition-colors
                  ${p.id === period
                    ? "bg-indigo-600/15 text-indigo-300"
                    : "text-slate-300 hover:bg-slate-700/50"
                  }
                `}
              >
                {p.label}
              </button>
            ))}
          </div>
        )}
      </div>

      {/* Quick period pills */}
      <div className="flex items-center gap-1.5">
        {PERIODS.map((p) => (
          <button
            key={p.id}
            onClick={() => onPeriodChange(p.id)}
            className={`
              px-3 py-1.5 rounded-lg text-xs font-medium transition-all
              ${p.id === period
                ? "bg-indigo-600 text-white shadow-sm shadow-indigo-500/30"
                : "bg-slate-800/60 text-slate-400 hover:bg-slate-700/60 hover:text-slate-200 border border-slate-700"
              }
            `}
          >
            {p.id.toUpperCase()}
          </button>
        ))}
      </div>
    </div>
  );
}
