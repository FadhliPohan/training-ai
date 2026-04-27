"use client";

import { useState } from "react";
import { ChevronDown, BarChart3 } from "lucide-react";
import { reportOptions } from "@/lib/dummyData";

export default function ReportSelector({ value, onChange }) {
  const [open, setOpen] = useState(false);
  const selected = reportOptions.find((r) => r.id === value);

  return (
    <div className="relative">
      <button
        onClick={() => setOpen(!open)}
        className="
          flex items-center gap-3 w-full md:w-auto min-w-[280px]
          bg-slate-800/60 border border-slate-700 hover:border-indigo-500/50
          rounded-xl px-4 py-2.5 text-sm text-left transition-all duration-200
          hover:bg-slate-700/60 focus:outline-none focus:ring-2 focus:ring-indigo-500/30
        "
      >
        <div className="flex-1">
          {selected ? (
            <span className="text-white font-medium">{selected.label}</span>
          ) : (
            <span className="text-slate-400">Pilih jenis laporan...</span>
          )}
          {selected && (
            <p className="text-[11px] text-slate-500 mt-0.5 truncate">{selected.description}</p>
          )}
        </div>
        <ChevronDown
          size={16}
          className={`text-slate-400 transition-transform duration-200 flex-shrink-0 ${open ? "rotate-180" : ""}`}
        />
      </button>

      {/* Dropdown */}
      {open && (
        <div className="
          absolute top-full left-0 mt-2 w-full md:w-[360px] z-50
          bg-[#1e293b] border border-slate-700 rounded-xl shadow-2xl shadow-black/40
          overflow-hidden animate-fade-in
        ">
          {reportOptions.map((opt) => {
            const isActive = opt.id === value;
            return (
              <button
                key={opt.id}
                onClick={() => { onChange(opt.id); setOpen(false); }}
                className={`
                  w-full text-left px-4 py-3 flex items-start gap-3 transition-colors
                  border-b border-slate-800 last:border-0
                  ${isActive
                    ? "bg-indigo-600/15 text-indigo-200"
                    : "text-slate-300 hover:bg-slate-700/50 hover:text-white"
                  }
                `}
              >
                <div className="mt-0.5">
                  <p className="text-sm font-medium">{opt.label}</p>
                  <p className="text-[11px] text-slate-500 mt-0.5">{opt.description}</p>
                </div>
                {isActive && (
                  <div className="ml-auto flex-shrink-0 w-2 h-2 rounded-full bg-indigo-400 mt-1.5" />
                )}
              </button>
            );
          })}
        </div>
      )}
    </div>
  );
}
