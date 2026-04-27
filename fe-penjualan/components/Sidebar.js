"use client";

import { useState } from "react";
import {
  TrendingUp, TrendingDown, ShoppingBag, Users,
  DollarSign, BarChart3, Menu, X, Bell,
  LayoutDashboard, Package, Settings, LogOut,
  ChevronRight, Zap
} from "lucide-react";

const navItems = [
  { icon: LayoutDashboard, label: "Dashboard",    href: "#dashboard",  active: true },
  { icon: BarChart3,       label: "Laporan",      href: "#laporan",    active: false },
  { icon: Package,         label: "Produk",       href: "#produk",     active: false },
  { icon: Users,           label: "Customer",     href: "#customer",   active: false },
  { icon: Settings,        label: "Pengaturan",   href: "#settings",   active: false },
];

export default function Sidebar() {
  const [collapsed, setCollapsed] = useState(false);
  const [activeItem, setActiveItem] = useState("Dashboard");

  return (
    <>
      {/* Mobile overlay */}
      {!collapsed && (
        <div
          className="fixed inset-0 bg-black/50 z-20 lg:hidden"
          onClick={() => setCollapsed(true)}
        />
      )}

      {/* Sidebar */}
      <aside
        className={`
          fixed top-0 left-0 h-full z-30 flex flex-col
          bg-[#1e293b] border-r border-slate-700/50
          transition-all duration-300 ease-in-out
          ${collapsed ? "w-0 lg:w-[72px] overflow-hidden lg:overflow-visible" : "w-[240px]"}
        `}
      >
        {/* Logo */}
        <div className="flex items-center gap-3 px-4 py-5 border-b border-slate-700/50 min-h-[64px]">
          <div className="flex-shrink-0 w-8 h-8 rounded-lg bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center shadow-lg shadow-indigo-500/30">
            <Zap size={16} className="text-white" />
          </div>
          {!collapsed && (
            <div className="overflow-hidden">
              <span className="font-bold text-sm tracking-tight text-white block">InsightFlow</span>
              <span className="text-[10px] text-slate-400 block">AI Sales Dashboard</span>
            </div>
          )}
        </div>

        {/* Nav */}
        <nav className="flex-1 py-4 overflow-y-auto">
          <div className={`px-3 mb-2 ${collapsed ? "hidden lg:block" : ""}`}>
            {!collapsed && (
              <span className="text-[10px] font-semibold text-slate-500 uppercase tracking-widest px-2">
                Menu Utama
              </span>
            )}
          </div>
          {navItems.map(({ icon: Icon, label, href }) => {
            const isActive = activeItem === label;
            return (
              <a
                key={label}
                href={href}
                onClick={() => setActiveItem(label)}
                className={`
                  flex items-center gap-3 mx-2 mb-1 px-3 py-2.5 rounded-xl
                  text-sm font-medium transition-all duration-200 group relative
                  ${isActive
                    ? "bg-indigo-600/20 text-indigo-300 shadow-sm shadow-indigo-500/10"
                    : "text-slate-400 hover:bg-slate-700/40 hover:text-slate-200"
                  }
                `}
              >
                {isActive && (
                  <span className="absolute left-0 top-1/2 -translate-y-1/2 w-1 h-6 bg-indigo-500 rounded-r-full" />
                )}
                <Icon size={18} className={isActive ? "text-indigo-400" : "text-slate-500 group-hover:text-slate-300"} />
                {!collapsed && (
                  <span className="truncate">{label}</span>
                )}
                {!collapsed && isActive && (
                  <ChevronRight size={14} className="ml-auto text-indigo-400" />
                )}

                {/* Tooltip for collapsed */}
                {collapsed && (
                  <span className="
                    absolute left-full ml-3 px-2 py-1 text-xs bg-slate-800 text-white
                    rounded-md opacity-0 group-hover:opacity-100 pointer-events-none
                    whitespace-nowrap transition-opacity z-50 border border-slate-700
                    hidden lg:block
                  ">
                    {label}
                  </span>
                )}
              </a>
            );
          })}
        </nav>

        {/* User profile */}
        <div className="p-3 border-t border-slate-700/50">
          <div className={`flex items-center gap-3 px-2 py-2 rounded-xl hover:bg-slate-700/40 cursor-pointer transition-colors`}>
            <div className="flex-shrink-0 w-8 h-8 rounded-full bg-gradient-to-br from-indigo-400 to-purple-500 flex items-center justify-center text-xs font-bold text-white">
              AD
            </div>
            {!collapsed && (
              <div className="overflow-hidden flex-1">
                <span className="text-xs font-semibold text-white block truncate">Administrator</span>
                <span className="text-[10px] text-slate-400 block truncate">admin@insightflow.id</span>
              </div>
            )}
            {!collapsed && (
              <LogOut size={14} className="text-slate-500 flex-shrink-0" />
            )}
          </div>
        </div>
      </aside>

      {/* Toggle button (visible on lg) */}
      <button
        onClick={() => setCollapsed(!collapsed)}
        className={`
          fixed z-40 top-4 bg-slate-800 border border-slate-700 rounded-lg p-1.5
          text-slate-400 hover:text-white transition-all duration-300 hidden lg:flex
          ${collapsed ? "left-[60px]" : "left-[228px]"}
        `}
      >
        {collapsed ? <Menu size={14} /> : <X size={14} />}
      </button>

      {/* Mobile hamburger */}
      <button
        onClick={() => setCollapsed(!collapsed)}
        className="fixed z-40 top-4 left-4 bg-slate-800 border border-slate-700 rounded-lg p-2 text-slate-400 hover:text-white lg:hidden"
      >
        <Menu size={18} />
      </button>
    </>
  );
}
