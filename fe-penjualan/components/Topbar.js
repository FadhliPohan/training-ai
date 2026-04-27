"use client";

import { Bell, Search, RefreshCw } from "lucide-react";
import { useState } from "react";

export default function Topbar({ title = "Dashboard", onRefresh }) {
  const [refreshing, setRefreshing] = useState(false);

  const handleRefresh = () => {
    setRefreshing(true);
    setTimeout(() => setRefreshing(false), 1200);
    onRefresh?.();
  };

  const now = new Date();
  const dateStr = now.toLocaleDateString("id-ID", {
    weekday: "long", day: "numeric", month: "long", year: "numeric",
  });

  return (
    <header className="flex items-center justify-between px-6 py-4 border-b border-slate-700/50 bg-[#1e293b]/70 backdrop-blur-md sticky top-0 z-10">
      {/* Left: title + date */}
      <div className="ml-8 lg:ml-0">
        <h1 className="text-lg font-bold text-white tracking-tight">{title}</h1>
        <p className="text-xs text-slate-400 mt-0.5">{dateStr}</p>
      </div>

      {/* Right: actions */}
      <div className="flex items-center gap-2">
        {/* Search (decorative for now) */}
        <div className="hidden md:flex items-center gap-2 bg-slate-800/60 border border-slate-700 rounded-xl px-3 py-2 w-52">
          <Search size={14} className="text-slate-500" />
          <input
            type="text"
            placeholder="Cari laporan..."
            className="bg-transparent text-sm text-slate-300 placeholder-slate-600 outline-none w-full"
          />
        </div>

        {/* Refresh */}
        <button
          onClick={handleRefresh}
          className="p-2 rounded-xl bg-slate-800/60 border border-slate-700 text-slate-400 hover:text-white hover:bg-slate-700/60 transition-all"
          title="Refresh data"
        >
          <RefreshCw size={16} className={refreshing ? "animate-spin" : ""} />
        </button>

        {/* Notifications */}
        <button className="relative p-2 rounded-xl bg-slate-800/60 border border-slate-700 text-slate-400 hover:text-white hover:bg-slate-700/60 transition-all">
          <Bell size={16} />
          <span className="absolute top-1.5 right-1.5 w-2 h-2 bg-indigo-500 rounded-full ring-2 ring-slate-900" />
        </button>

        {/* Avatar */}
        <div className="w-8 h-8 rounded-full bg-gradient-to-br from-indigo-400 to-purple-600 flex items-center justify-center text-xs font-bold text-white cursor-pointer ml-1 hover:ring-2 hover:ring-indigo-500/50 transition-all">
          AD
        </div>
      </div>
    </header>
  );
}
