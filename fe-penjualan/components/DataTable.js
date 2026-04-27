"use client";

import { useState } from "react";
import { Search, ChevronLeft, ChevronRight } from "lucide-react";

const PAGE_SIZE = 10;

/**
 * Generic DataTable with search + pagination.
 *
 * Props:
 *  - columns: [{ key, label, render?, align? }]
 *  - data: array of row objects
 *  - searchKeys: array of keys to search against
 *  - emptyMessage: string
 */
export default function DataTable({ columns, data = [], searchKeys = [], emptyMessage = "Tidak ada data." }) {
    const [query, setQuery] = useState("");
    const [page, setPage] = useState(1);

    const filtered = query.trim()
        ? data.filter((row) =>
            searchKeys.some((k) =>
                String(row[k] ?? "").toLowerCase().includes(query.toLowerCase())
            )
        )
        : data;

    const totalPages = Math.max(1, Math.ceil(filtered.length / PAGE_SIZE));
    const safePage = Math.min(page, totalPages);
    const rows = filtered.slice((safePage - 1) * PAGE_SIZE, safePage * PAGE_SIZE);

    return (
        <div className="space-y-4">
            {/* Search */}
            <div className="flex items-center gap-2 bg-slate-800/60 border border-slate-700 rounded-xl px-3 py-2 w-full max-w-xs">
                <Search size={14} className="text-slate-500 flex-shrink-0" />
                <input
                    type="text"
                    value={query}
                    onChange={(e) => { setQuery(e.target.value); setPage(1); }}
                    placeholder="Cari..."
                    className="bg-transparent text-sm text-slate-300 placeholder-slate-600 outline-none w-full"
                />
            </div>

            {/* Table */}
            <div className="overflow-x-auto rounded-xl border border-slate-700/50">
                <table className="w-full text-sm">
                    <thead>
                        <tr className="border-b border-slate-700/50 bg-slate-800/40">
                            {columns.map((col) => (
                                <th
                                    key={col.key}
                                    className={`py-3 px-4 text-[11px] font-semibold text-slate-500 uppercase tracking-wider whitespace-nowrap ${col.align === "right" ? "text-right" : col.align === "center" ? "text-center" : "text-left"
                                        }`}
                                >
                                    {col.label}
                                </th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {rows.length === 0 ? (
                            <tr>
                                <td colSpan={columns.length} className="py-12 text-center text-slate-500 text-sm">
                                    {emptyMessage}
                                </td>
                            </tr>
                        ) : (
                            rows.map((row, i) => (
                                <tr
                                    key={row.id ?? i}
                                    className="border-b border-slate-800/60 hover:bg-slate-700/20 transition-colors"
                                >
                                    {columns.map((col) => (
                                        <td
                                            key={col.key}
                                            className={`py-3 px-4 ${col.align === "right" ? "text-right" : col.align === "center" ? "text-center" : ""
                                                }`}
                                        >
                                            {col.render ? col.render(row[col.key], row) : (
                                                <span className="text-slate-300">{row[col.key] ?? "—"}</span>
                                            )}
                                        </td>
                                    ))}
                                </tr>
                            ))
                        )}
                    </tbody>
                </table>
            </div>

            {/* Pagination */}
            {totalPages > 1 && (
                <div className="flex items-center justify-between text-xs text-slate-500">
                    <span>
                        {filtered.length} data · halaman {safePage} dari {totalPages}
                    </span>
                    <div className="flex items-center gap-1">
                        <button
                            onClick={() => setPage((p) => Math.max(1, p - 1))}
                            disabled={safePage === 1}
                            className="p-1.5 rounded-lg border border-slate-700 hover:bg-slate-700/50 disabled:opacity-40 disabled:cursor-not-allowed transition-all"
                        >
                            <ChevronLeft size={14} />
                        </button>
                        {Array.from({ length: totalPages }, (_, i) => i + 1)
                            .filter((p) => p === 1 || p === totalPages || Math.abs(p - safePage) <= 1)
                            .reduce((acc, p, idx, arr) => {
                                if (idx > 0 && p - arr[idx - 1] > 1) acc.push("...");
                                acc.push(p);
                                return acc;
                            }, [])
                            .map((p, i) =>
                                p === "..." ? (
                                    <span key={`ellipsis-${i}`} className="px-1">…</span>
                                ) : (
                                    <button
                                        key={p}
                                        onClick={() => setPage(p)}
                                        className={`w-7 h-7 rounded-lg text-xs font-medium transition-all ${p === safePage
                                                ? "bg-indigo-600 text-white"
                                                : "border border-slate-700 hover:bg-slate-700/50 text-slate-400"
                                            }`}
                                    >
                                        {p}
                                    </button>
                                )
                            )}
                        <button
                            onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
                            disabled={safePage === totalPages}
                            className="p-1.5 rounded-lg border border-slate-700 hover:bg-slate-700/50 disabled:opacity-40 disabled:cursor-not-allowed transition-all"
                        >
                            <ChevronRight size={14} />
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
}
