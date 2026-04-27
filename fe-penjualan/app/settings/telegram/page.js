"use client";

import { useState, useEffect, useCallback } from "react";
import { Save, Loader2, AlertCircle, CheckCircle, Send, Clock, Bell } from "lucide-react";

import ProtectedRoute from "@/components/ProtectedRoute";
import PageShell from "@/components/PageShell";
import { settingsAPI } from "@/lib/api";

const EMPTY_FORM = {
    nama_grup: "",
    chat_id: "",
    jam_summary: "07:00",
    threshold_pct: 10,
    aktif: true,
};

export default function TelegramSettingsPage() {
    return (
        <ProtectedRoute>
            <TelegramSettingsContent />
        </ProtectedRoute>
    );
}

function TelegramSettingsContent() {
    const [form, setForm] = useState(EMPTY_FORM);
    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");

    const fetchData = useCallback(async () => {
        setLoading(true);
        setError("");
        try {
            const res = await settingsAPI.getTelegram();
            const d = res.data ?? res ?? {};
            setForm({
                nama_grup: d.nama_grup ?? "",
                chat_id: d.chat_id != null ? String(d.chat_id) : "",
                jam_summary: d.jam_summary ?? "07:00",
                threshold_pct: d.threshold_pct ?? 10,
                aktif: d.aktif ?? true,
            });
        } catch (e) {
            setError(e.message);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => { fetchData(); }, [fetchData]);

    const handleSave = async (e) => {
        e.preventDefault();
        setError("");
        setSuccess("");
        setSaving(true);
        try {
            await settingsAPI.updateTelegram({
                nama_grup: form.nama_grup,
                chat_id: Number(form.chat_id),
                jam_summary: form.jam_summary,
                threshold_pct: Number(form.threshold_pct),
                aktif: form.aktif,
            });
            setSuccess("Pengaturan Telegram berhasil disimpan.");
            setTimeout(() => setSuccess(""), 4000);
        } catch (e) {
            setError(e.message);
        } finally {
            setSaving(false);
        }
    };

    if (loading) {
        return (
            <PageShell title="Pengaturan Telegram">
                <div className="flex items-center justify-center py-32">
                    <Loader2 size={28} className="animate-spin text-indigo-400" />
                </div>
            </PageShell>
        );
    }

    return (
        <PageShell title="Pengaturan Telegram" onRefresh={fetchData}>
            <div>
                <h2 className="text-base font-semibold text-white">Konfigurasi Telegram Bot</h2>
                <p className="text-xs text-slate-500 mt-0.5">Atur nama grup, chat ID, dan jadwal notifikasi otomatis</p>
            </div>

            {error && (
                <div className="flex items-center gap-2 rounded-xl bg-rose-500/10 border border-rose-500/30 px-4 py-3 text-sm text-rose-400">
                    <AlertCircle size={16} /> {error}
                </div>
            )}
            {success && (
                <div className="flex items-center gap-2 rounded-xl bg-emerald-500/10 border border-emerald-500/30 px-4 py-3 text-sm text-emerald-400">
                    <CheckCircle size={16} /> {success}
                </div>
            )}

            <form onSubmit={handleSave} className="space-y-6">
                {/* Bot Configuration */}
                <div className="rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 p-6 space-y-5">
                    <div className="flex items-center gap-2 pb-3 border-b border-slate-700/40">
                        <Send size={16} className="text-indigo-400" />
                        <h3 className="text-sm font-semibold text-white">Konfigurasi Bot</h3>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-2">Nama Grup <span className="text-rose-400">*</span></label>
                        <input required value={form.nama_grup} onChange={(e) => setForm({ ...form, nama_grup: e.target.value })} placeholder="Sales Team InsightFlow" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        <p className="text-[11px] text-slate-600 mt-1.5">Nama identifikasi grup Telegram untuk konfigurasi ini.</p>
                    </div>

                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-2">Chat ID <span className="text-rose-400">*</span></label>
                        <input required value={form.chat_id} onChange={(e) => setForm({ ...form, chat_id: e.target.value })} placeholder="-1001234567890" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all font-mono" />
                        <p className="text-[11px] text-slate-600 mt-1.5">
                            ID grup atau channel Telegram. Gunakan{" "}
                            <a href="https://t.me/userinfobot" target="_blank" rel="noopener noreferrer" className="text-indigo-400 hover:underline">@userinfobot</a>{" "}
                            untuk mendapatkan ID.
                        </p>
                    </div>

                    <div className="flex items-center justify-between pt-3 border-t border-slate-700/40">
                        <div>
                            <span className="text-sm font-medium text-white">Status Bot</span>
                            <p className="text-[11px] text-slate-500 mt-0.5">
                                {form.aktif ? "Bot aktif dan akan mengirim notifikasi" : "Bot dinonaktifkan sementara"}
                            </p>
                        </div>
                        <button type="button" onClick={() => setForm({ ...form, aktif: !form.aktif })} className={`relative inline-flex h-7 w-12 items-center rounded-full transition-colors ${form.aktif ? "bg-indigo-600" : "bg-slate-700"}`}>
                            <span className={`inline-block h-5 w-5 transform rounded-full bg-white transition-transform ${form.aktif ? "translate-x-6" : "translate-x-1"}`} />
                        </button>
                    </div>
                </div>

                {/* Schedule & Notifications */}
                <div className="rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 p-6 space-y-5">
                    <div className="flex items-center gap-2 pb-3 border-b border-slate-700/40">
                        <Bell size={16} className="text-indigo-400" />
                        <h3 className="text-sm font-semibold text-white">Jadwal & Notifikasi</h3>
                    </div>
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-5">
                        <div>
                            <label className="flex items-center gap-2 text-xs font-semibold text-slate-400 uppercase tracking-wider mb-2">
                                <Clock size={12} /> Waktu Daily Summary
                            </label>
                            <input type="time" value={form.jam_summary} onChange={(e) => setForm({ ...form, jam_summary: e.target.value })} className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                            <p className="text-[11px] text-slate-600 mt-1.5">Laporan harian dikirim otomatis setiap hari pada waktu ini (WIB).</p>
                        </div>
                        <div>
                            <label className="flex items-center gap-2 text-xs font-semibold text-slate-400 uppercase tracking-wider mb-2">
                                <AlertCircle size={12} /> Threshold Anomali (%)
                            </label>
                            <input type="number" min="0" max="100" value={form.threshold_pct} onChange={(e) => setForm({ ...form, threshold_pct: e.target.value })} className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                            <p className="text-[11px] text-slate-600 mt-1.5">Notifikasi anomali dikirim jika deviasi melebihi persentase ini.</p>
                        </div>
                    </div>
                </div>

                {/* Info */}
                <div className="rounded-2xl border border-indigo-500/20 bg-gradient-to-br from-indigo-900/20 to-slate-800/40 p-5">
                    <div className="flex items-start gap-3">
                        <div className="p-2 rounded-lg bg-indigo-500/20 flex-shrink-0">
                            <Send size={16} className="text-indigo-300" />
                        </div>
                        <div>
                            <h4 className="text-sm font-semibold text-indigo-300 mb-1">Fitur Telegram Bot</h4>
                            <ul className="text-xs text-slate-400 space-y-1">
                                <li>• <strong>Daily Summary:</strong> Laporan penjualan harian otomatis setiap pagi</li>
                                <li>• <strong>Anomaly Alert:</strong> Notifikasi real-time jika terdeteksi anomali data</li>
                                <li>• <strong>On-demand Q&A:</strong> Tanya jawab data penjualan via chat Telegram</li>
                            </ul>
                        </div>
                    </div>
                </div>

                <div className="flex justify-end gap-3">
                    <button type="button" onClick={fetchData} className="px-4 py-2.5 rounded-xl text-sm font-medium border border-slate-700 text-slate-400 hover:text-white hover:bg-slate-700/40 transition-all">Reset</button>
                    <button type="submit" disabled={saving} className="flex items-center gap-2 px-6 py-2.5 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all disabled:opacity-60">
                        {saving ? <><Loader2 size={14} className="animate-spin" /> Menyimpan...</> : <><Save size={14} /> Simpan Pengaturan</>}
                    </button>
                </div>
            </form>
        </PageShell>
    );
}
