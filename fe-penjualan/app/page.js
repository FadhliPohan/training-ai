"use client";

import { useState, useCallback } from "react";
import { Download, Play, Loader2 } from "lucide-react";

import ProtectedRoute from "@/components/ProtectedRoute";
import PageShell from "@/components/PageShell";
import KPICard from "@/components/KPICard";
import ReportSelector from "@/components/ReportSelector";
import FilterBar from "@/components/FilterBar";
import ChartRenderer from "@/components/ChartRenderer";
import AIInsightCard from "@/components/AIInsightCard";
import AnomalyFlag from "@/components/AnomalyFlag";
import TopProdukTable from "@/components/TopProdukTable";
import ChatBot from "@/components/ChatBot";

import { kpiData, anomaliData, aiInsights, salesPerformData } from "@/lib/dummyData";

export default function DashboardPage() {
  return (
    <ProtectedRoute>
      <DashboardContent />
    </ProtectedRoute>
  );
}

function DashboardContent() {
  const [selectedReport, setSelectedReport] = useState("penjualan-harian");
  const [period, setPeriod] = useState("30d");
  const [loadingReport, setLoadingReport] = useState(false);
  const [reportActive, setReportActive] = useState(true);

  const handleTampilkan = useCallback(() => {
    setLoadingReport(true);
    setReportActive(false);
    setTimeout(() => {
      setLoadingReport(false);
      setReportActive(true);
    }, 1400);
  }, []);

  const handleReportChange = (id) => {
    setSelectedReport(id);
    setReportActive(false);
  };

  const currentInsight = aiInsights[selectedReport];

  return (
    <>
      <PageShell title="Dashboard Analitik">

      {/* KPI Cards */}
      <section>
        <div className="flex items-center justify-between mb-4">
          <div>
            <h2 className="text-base font-semibold text-white">Ringkasan Performa</h2>
            <p className="text-xs text-slate-500 mt-0.5">Data dummy — April 2026</p>
          </div>
          <span className="text-[10px] text-indigo-300 bg-indigo-500/10 border border-indigo-500/20 px-2.5 py-1 rounded-full font-medium">
            🔴 LIVE (Dummy)
          </span>
        </div>
        <div className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4">
          {kpiData.map((kpi, i) => (
            <KPICard key={kpi.id} kpi={kpi} index={i} />
          ))}
        </div>
      </section>

      {/* Self-Service Report */}
      <section className="rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 p-5 space-y-5 overflow-visible">
        <div>
          <h2 className="text-base font-semibold text-white">Laporan Self-Service</h2>
          <p className="text-xs text-slate-500 mt-0.5">Pilih jenis laporan, atur filter, klik Tampilkan</p>
        </div>

        {/* Controls — stacked grid to prevent overflow */}
        <div className="grid grid-cols-1 gap-4">

          {/* Row 1: Jenis Laporan — full width */}
          <div>
            <label className="text-[10px] font-semibold text-slate-500 uppercase tracking-wider block mb-2">
              Jenis Laporan
            </label>
            <ReportSelector value={selectedReport} onChange={handleReportChange} />
          </div>

          {/* Row 2: Periode + action buttons */}
          <div className="flex flex-wrap items-end gap-3">
            <div className="flex-1 min-w-0">
              <label className="text-[10px] font-semibold text-slate-500 uppercase tracking-wider block mb-2">
                Periode
              </label>
              <FilterBar period={period} onPeriodChange={setPeriod} />
            </div>
            <div className="flex gap-2 flex-shrink-0">
              <button
                onClick={handleTampilkan}
                disabled={loadingReport}
                className="flex items-center gap-2 px-5 py-2.5 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-700 text-white shadow-lg shadow-indigo-600/30 transition-all duration-200 disabled:opacity-60 disabled:cursor-wait"
              >
                {loadingReport
                  ? <><Loader2 size={15} className="animate-spin" /> Memproses AI...</>
                  : <><Play size={15} /> Tampilkan</>
                }
              </button>
              <button className="flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-medium border border-slate-700 text-slate-400 hover:text-white hover:border-slate-600 hover:bg-slate-700/40 transition-all">
                <Download size={15} /> Export
              </button>
            </div>
          </div>

        </div>

        <div className="rounded-xl border border-slate-700/40 bg-slate-900/40 p-5">
          {loadingReport ? (
            <div className="flex flex-col items-center justify-center h-[300px] gap-3">
              <div className="relative">
                <div className="w-12 h-12 rounded-full border-2 border-indigo-500/30 border-t-indigo-500 animate-spin" />
                <div className="absolute inset-0 flex items-center justify-center">
                  <span className="text-lg">✨</span>
                </div>
              </div>
              <p className="text-sm text-slate-400">AI sedang memproses data...</p>
              <p className="text-xs text-slate-600">Mengirim ke n8n → LLM → Analisis</p>
            </div>
          ) : reportActive ? (
            <ChartRenderer reportId={selectedReport} />
          ) : (
            <div className="flex flex-col items-center justify-center h-[300px] gap-3 text-slate-500">
              <span className="text-4xl">📊</span>
              <p className="text-sm">Klik "Tampilkan" untuk melihat laporan</p>
            </div>
          )}
        </div>

        {reportActive && !loadingReport && <AIInsightCard insight={currentInsight} />}
      </section>

      {/* Bottom Grid */}
      <div className="grid grid-cols-1 xl:grid-cols-3 gap-6">
        {/* Top Produk */}
        <div className="xl:col-span-2 rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 overflow-hidden">
          <div className="px-5 py-4 border-b border-slate-700/40 flex items-center justify-between">
            <div>
              <h3 className="text-sm font-semibold text-white">🏆 Top Produk Bulan Ini</h3>
              <p className="text-xs text-slate-500 mt-0.5">Berdasarkan volume terjual</p>
            </div>
            <span className="text-[10px] text-slate-500">April 2026</span>
          </div>
          <TopProdukTable />
        </div>

        {/* Right column */}
        <div className="space-y-5">
          {/* Anomali */}
          <div className="rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 overflow-hidden">
            <div className="px-5 py-4 border-b border-slate-700/40 flex items-center gap-2">
              <span className="text-base">⚠️</span>
              <div>
                <h3 className="text-sm font-semibold text-white">Anomali Terdeteksi</h3>
                <p className="text-xs text-slate-500">{anomaliData.length} isu aktif</p>
              </div>
            </div>
            <div className="p-4 space-y-3">
              {anomaliData.map((a) => <AnomalyFlag key={a.id} anomali={a} />)}
            </div>
          </div>

          {/* Performa Sales */}
          <div className="rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 overflow-hidden">
            <div className="px-5 py-4 border-b border-slate-700/40">
              <h3 className="text-sm font-semibold text-white">👥 Performa Sales</h3>
              <p className="text-xs text-slate-500 mt-0.5">vs Target April 2026</p>
            </div>
            <div className="p-4 space-y-3">
              {salesPerformData.map((s, i) => {
                const pct = Math.min((s.omzet / s.target) * 100, 100);
                const isOver = s.omzet >= s.target;
                const isLow = pct < 50;
                return (
                  <div key={i} className="space-y-1.5">
                    <div className="flex items-center justify-between">
                      <span className="text-xs font-medium text-slate-200">{s.nama}</span>
                      <span className={`text-[11px] font-bold ${isOver ? "text-emerald-400" : isLow ? "text-rose-400" : "text-amber-400"}`}>
                        {pct.toFixed(0)}%
                      </span>
                    </div>
                    <div className="h-1.5 rounded-full bg-slate-800">
                      <div
                        className={`h-full rounded-full transition-all duration-700 ${isOver ? "bg-emerald-500" : isLow ? "bg-rose-500" : "bg-amber-500"}`}
                        style={{ width: `${pct}%` }}
                      />
                    </div>
                    <p className="text-[10px] text-slate-500">
                      Rp {(s.omzet / 1_000_000).toFixed(1)} Jt / Target Rp {(s.target / 1_000_000).toFixed(0)} Jt
                    </p>
                  </div>
                );
              })}
            </div>
          </div>
        </div>
      </div>

      <footer className="text-center text-xs text-slate-600 pb-4 pt-2">
        InsightFlow v1.0 · Dashboard Dummy Data · Powered by Next.js + Recharts
      </footer>

    </PageShell>

    {/* AI Chatbot — floating widget, rendered outside PageShell scroll container */}
    <ChatBot />
    </>
  );
}
