"use client";

import { useState, useEffect, useCallback } from "react";
import { Plus, Pencil, ToggleLeft, ToggleRight, Loader2, AlertCircle } from "lucide-react";

import ProtectedRoute from "@/components/ProtectedRoute";
import PageShell from "@/components/PageShell";
import DataTable from "@/components/DataTable";
import Modal from "@/components/Modal";
import { produkAPI, formatRupiah } from "@/lib/api";

const EMPTY_FORM = {
    kode_produk: "",
    nama: "",
    kategori_pakaian: "",
    ukuran: "",
    warna: "",
    bahan: "",
    harga: "",
    stok: "",
    aktif: true,
};

const KATEGORI_OPTIONS = [
    { value: "atasan", label: "Atasan" },
    { value: "bawahan", label: "Bawahan" },
    { value: "dress", label: "Dress" },
    { value: "outerwear", label: "Outerwear" },
    { value: "aksesoris", label: "Aksesoris" },
];

export default function ProdukPage() {
    return (
        <ProtectedRoute>
            <ProdukContent />
        </ProtectedRoute>
    );
}

function ProdukContent() {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const [modalOpen, setModalOpen] = useState(false);
    const [editTarget, setEditTarget] = useState(null);
    const [form, setForm] = useState(EMPTY_FORM);
    const [saving, setSaving] = useState(false);
    const [formError, setFormError] = useState("");

    const fetchData = useCallback(async () => {
        setLoading(true);
        setError("");
        try {
            const res = await produkAPI.list();
            setData(res.data ?? res ?? []);
        } catch (e) {
            setError(e.message);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => { fetchData(); }, [fetchData]);

    const openCreate = () => {
        setEditTarget(null);
        setForm(EMPTY_FORM);
        setFormError("");
        setModalOpen(true);
    };

    const openEdit = (row) => {
        setEditTarget(row);
        setForm({
            kode_produk: row.kode_produk ?? "",
            nama: row.nama ?? "",
            kategori_pakaian: row.kategori_pakaian ?? "",
            ukuran: row.ukuran ?? "",
            warna: row.warna ?? "",
            bahan: row.bahan ?? "",
            harga: String(row.harga ?? ""),
            stok: String(row.stok ?? ""),
            aktif: row.aktif ?? true,
        });
        setFormError("");
        setModalOpen(true);
    };

    const handleSave = async (e) => {
        e.preventDefault();
        setFormError("");
        setSaving(true);
        try {
            const payload = { ...form, harga: Number(form.harga), stok: Number(form.stok) };
            if (editTarget) {
                await produkAPI.update(editTarget.id, payload);
            } else {
                await produkAPI.create(payload);
            }
            setModalOpen(false);
            fetchData();
        } catch (e) {
            setFormError(e.message);
        } finally {
            setSaving(false);
        }
    };

    const handleToggleAktif = async (row) => {
        try {
            await produkAPI.deactivate(row.id);
            fetchData();
        } catch (e) {
            alert(e.message);
        }
    };

    const columns = [
        {
            key: "nama",
            label: "Nama Produk",
            render: (v) => <span className="font-medium text-slate-200">{v}</span>,
        },
        {
            key: "kategori_pakaian",
            label: "Kategori",
            render: (v) => (
                <span className="text-[11px] text-slate-400 bg-slate-700/50 px-2 py-0.5 rounded-full capitalize">{v}</span>
            ),
        },
        { key: "ukuran", label: "Ukuran" },
        { key: "warna", label: "Warna" },
        {
            key: "harga",
            label: "Harga",
            align: "right",
            render: (v) => <span className="font-semibold text-slate-200">{formatRupiah(v)}</span>,
        },
        {
            key: "stok",
            label: "Stok",
            align: "center",
            render: (v) => {
                const n = Number(v);
                const cls = n <= 5
                    ? "text-rose-400 bg-rose-500/10 border-rose-500/20"
                    : n <= 15
                        ? "text-amber-400 bg-amber-500/10 border-amber-500/20"
                        : "text-emerald-400 bg-emerald-500/10 border-emerald-500/20";
                return <span className={`text-[11px] font-semibold px-2 py-0.5 rounded-full border ${cls}`}>{n}</span>;
            },
        },
        {
            key: "aktif",
            label: "Status",
            align: "center",
            render: (v, row) => (
                <button onClick={() => handleToggleAktif(row)} className="inline-flex items-center gap-1.5 text-xs font-medium transition-colors">
                    {v ? (
                        <><ToggleRight size={18} className="text-emerald-400" /><span className="text-emerald-400">Aktif</span></>
                    ) : (
                        <><ToggleLeft size={18} className="text-slate-500" /><span className="text-slate-500">Nonaktif</span></>
                    )}
                </button>
            ),
        },
        {
            key: "id",
            label: "Aksi",
            align: "center",
            render: (_, row) => (
                <button onClick={() => openEdit(row)} className="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium bg-indigo-600/15 text-indigo-300 hover:bg-indigo-600/30 border border-indigo-500/20 transition-all">
                    <Pencil size={12} /> Edit
                </button>
            ),
        },
    ];

    return (
        <PageShell title="Manajemen Produk" onRefresh={fetchData}>
            <div className="flex items-center justify-between">
                <div>
                    <h2 className="text-base font-semibold text-white">Daftar Produk</h2>
                    <p className="text-xs text-slate-500 mt-0.5">{data.length} produk terdaftar</p>
                </div>
                <button onClick={openCreate} className="flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all">
                    <Plus size={16} /> Tambah Produk
                </button>
            </div>

            {error && (
                <div className="flex items-center gap-2 rounded-xl bg-rose-500/10 border border-rose-500/30 px-4 py-3 text-sm text-rose-400">
                    <AlertCircle size={16} /> {error}
                </div>
            )}

            <div className="rounded-2xl border border-slate-700/50 bg-[#1e293b]/60 p-5">
                {loading ? (
                    <div className="flex items-center justify-center py-16 gap-3 text-slate-400">
                        <Loader2 size={20} className="animate-spin" />
                        <span className="text-sm">Memuat data...</span>
                    </div>
                ) : (
                    <DataTable
                        columns={columns}
                        data={data}
                        searchKeys={["nama", "kategori_pakaian", "warna"]}
                        emptyMessage="Belum ada produk. Klik 'Tambah Produk' untuk memulai."
                    />
                )}
            </div>

            <Modal open={modalOpen} onClose={() => setModalOpen(false)} title={editTarget ? "Edit Produk" : "Tambah Produk Baru"} size="lg">
                <form onSubmit={handleSave} className="space-y-4">
                    {formError && (
                        <div className="rounded-xl bg-rose-500/10 border border-rose-500/30 px-4 py-3 text-sm text-rose-400">{formError}</div>
                    )}
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <div className="sm:col-span-2">
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Kode Produk <span className="text-rose-400">*</span></label>
                            <input required value={form.kode_produk} onChange={(e) => setForm({ ...form, kode_produk: e.target.value })} placeholder="KAO-001-M-HTM" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all font-mono" />
                        </div>
                        <div className="sm:col-span-2">
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Nama Produk <span className="text-rose-400">*</span></label>
                            <input required value={form.nama} onChange={(e) => setForm({ ...form, nama: e.target.value })} placeholder="Kemeja Batik Lengan Panjang" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Kategori <span className="text-rose-400">*</span></label>
                            <select required value={form.kategori_pakaian} onChange={(e) => setForm({ ...form, kategori_pakaian: e.target.value })} className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all">
                                <option value="">Pilih kategori...</option>
                                {KATEGORI_OPTIONS.map((k) => <option key={k.value} value={k.value}>{k.label}</option>)}
                            </select>
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Ukuran</label>
                            <input value={form.ukuran} onChange={(e) => setForm({ ...form, ukuran: e.target.value })} placeholder="S, M, L, XL, XXL" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Warna</label>
                            <input value={form.warna} onChange={(e) => setForm({ ...form, warna: e.target.value })} placeholder="Hitam, Putih, Navy..." className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Bahan</label>
                            <input value={form.bahan} onChange={(e) => setForm({ ...form, bahan: e.target.value })} placeholder="Cotton, Polyester, Linen..." className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Harga (Rp) <span className="text-rose-400">*</span></label>
                            <input required type="number" min="0" value={form.harga} onChange={(e) => setForm({ ...form, harga: e.target.value })} placeholder="185000" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                        <div>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Stok <span className="text-rose-400">*</span></label>
                            <input required type="number" min="0" value={form.stok} onChange={(e) => setForm({ ...form, stok: e.target.value })} placeholder="50" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                        {editTarget && (
                            <div className="sm:col-span-2 flex items-center gap-3">
                                <label className="text-xs font-semibold text-slate-400 uppercase tracking-wider">Status Produk</label>
                                <button type="button" onClick={() => setForm({ ...form, aktif: !form.aktif })} className={`flex items-center gap-2 px-3 py-1.5 rounded-lg text-xs font-medium border transition-all ${form.aktif ? "bg-emerald-500/15 text-emerald-400 border-emerald-500/30" : "bg-slate-700/50 text-slate-400 border-slate-600"}`}>
                                    {form.aktif ? "✓ Aktif" : "✗ Nonaktif"}
                                </button>
                            </div>
                        )}
                    </div>
                    <div className="flex justify-end gap-3 pt-2 border-t border-slate-700/50">
                        <button type="button" onClick={() => setModalOpen(false)} className="px-4 py-2.5 rounded-xl text-sm font-medium border border-slate-700 text-slate-400 hover:text-white hover:bg-slate-700/40 transition-all">Batal</button>
                        <button type="submit" disabled={saving} className="flex items-center gap-2 px-5 py-2.5 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all disabled:opacity-60">
                            {saving ? <><Loader2 size={14} className="animate-spin" /> Menyimpan...</> : "Simpan"}
                        </button>
                    </div>
                </form>
            </Modal>
        </PageShell>
    );
}
