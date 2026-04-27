"use client";

import { topProdukData } from "@/lib/dummyData";
import { TrendingUp, AlertTriangle } from "lucide-react";

function StokBadge({ stok }) {
  if (stok <= 5) return (
    <span className="inline-flex items-center gap-1 text-[10px] font-semibold text-rose-400 bg-rose-500/10 px-2 py-0.5 rounded-full border border-rose-500/20">
      <AlertTriangle size={9} /> Kritis ({stok})
    </span>
  );
  if (stok <= 15) return (
    <span className="text-[10px] font-semibold text-amber-400 bg-amber-500/10 px-2 py-0.5 rounded-full border border-amber-500/20">
      Rendah ({stok})
    </span>
  );
  return (
    <span className="text-[10px] font-semibold text-emerald-400 bg-emerald-500/10 px-2 py-0.5 rounded-full border border-emerald-500/20">
      Aman ({stok})
    </span>
  );
}

export default function TopProdukTable() {
  const maxOmzet = Math.max(...topProdukData.map((p) => p.omzet));

  return (
    <div className="overflow-x-auto">
      <table className="w-full text-sm">
        <thead>
          <tr className="border-b border-slate-700/50">
            <th className="text-left py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider">#</th>
            <th className="text-left py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider">Produk</th>
            <th className="text-left py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider">Kategori</th>
            <th className="text-right py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider">Terjual</th>
            <th className="text-left py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider min-w-[140px]">Omzet</th>
            <th className="text-center py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider">Stok</th>
          </tr>
        </thead>
        <tbody>
          {topProdukData.map((p, i) => {
            const barWidth = (p.omzet / maxOmzet) * 100;
            return (
              <tr
                key={i}
                className="border-b border-slate-800/60 hover:bg-slate-700/20 transition-colors group"
                style={{ animationDelay: `${i * 50}ms` }}
              >
                <td className="py-3 px-4">
                  <span className={`
                    w-6 h-6 rounded-full flex items-center justify-center text-[11px] font-bold
                    ${i === 0 ? "bg-amber-500/20 text-amber-400"
                      : i === 1 ? "bg-slate-400/20 text-slate-400"
                      : i === 2 ? "bg-orange-700/20 text-orange-600"
                      : "text-slate-600"
                    }
                  `}>
                    {i + 1}
                  </span>
                </td>
                <td className="py-3 px-4">
                  <span className="font-medium text-slate-200 group-hover:text-white transition-colors">
                    {p.nama}
                  </span>
                </td>
                <td className="py-3 px-4">
                  <span className="text-[11px] text-slate-400 bg-slate-700/50 px-2 py-0.5 rounded-full">
                    {p.kategori}
                  </span>
                </td>
                <td className="py-3 px-4 text-right">
                  <div className="flex items-center justify-end gap-1.5">
                    <TrendingUp size={12} className="text-indigo-400" />
                    <span className="font-semibold text-white">{p.terjual}</span>
                    <span className="text-slate-500 text-xs">unit</span>
                  </div>
                </td>
                <td className="py-3 px-4">
                  <div className="flex flex-col gap-1">
                    <span className="text-xs font-semibold text-slate-200">
                      Rp {(p.omzet / 1_000_000).toFixed(1)} Jt
                    </span>
                    <div className="h-1.5 rounded-full bg-slate-800 overflow-hidden">
                      <div
                        className="h-full rounded-full bg-gradient-to-r from-indigo-500 to-purple-500 transition-all duration-700"
                        style={{ width: `${barWidth}%` }}
                      />
                    </div>
                  </div>
                </td>
                <td className="py-3 px-4 text-center">
                  <StokBadge stok={p.stok} />
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
