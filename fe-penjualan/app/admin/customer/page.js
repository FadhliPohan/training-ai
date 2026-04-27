"use client";

import { useState, useEffect, useCallback } from "react";
import { Plus, Pencil, Loader2, AlertCircle, Phone, MapPin } from "lucide-react";

import ProtectedRoute from "@/components/ProtectedRoute";
import PageShell from "@/components/PageShell";
import DataTable from "@/components/DataTable";
import Modal from "@/components/Modal";
import { customerAPI } from "@/lib/api";

const EMPTY_FORM = {
    kode_cust: "",
    nama: "",
    email: "",
    telepon: "",
    alamat: "",
};

export default function CustomerPage() {
    return (
        <ProtectedRoute>
            <CustomerContent />
        </ProtectedRoute>
    );
}

function CustomerContent() {
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
            const res = await customerAPI.list();
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
            kode_cust: row.kode_cust ?? "",
            nama: row.nama ?? "",
            email: row.email ?? "",
            telepon: row.telepon ?? "",
            alamat: row.alamat ?? "",
        });
        setFormError("");
        setModalOpen(true);
    };

    const handleSave = async (e) => {
        e.preventDefault();
        setFormError("");
        setSaving(true);
        try {
            if (editTarget) {
                await customerAPI.update(editTarget.id, { ...form, aktif: true });
            } else {
                await customerAPI.create(form);
            }
            setModalOpen(false);
            fetchData();
        } catch (e) {
            setFormError(e.message);
        } finally {
            setSaving(false);
        }
    };

    const columns = [
        {
            key: "nama",
            label: "Nama",
            render: (v) => <span className="font-medium text-slate-200">{v}</span>,
        },
        {
            key: "email",
            label: "Email",
            render: (v) => <span className="text-slate-400 text-xs">{v || "—"}</span>,
        },
        {
            key: "telepon",
            label: "Telepon",
            render: (v) => v ? (
                <span className="inline-flex items-center gap-1 text-xs text-slate-300">
                    <Phone size={11} className="text-slate-500" /> {v}
                </span>
            ) : <span className="text-slate-600">—</span>,
        },
        {
            key: "alamat",
            label: "Alamat",
            render: (v) => v ? (
                <span className="inline-flex items-center gap-1 text-xs text-slate-300 max-w-[180px] truncate">
                    <MapPin size={11} className="text-slate-500 flex-shrink-0" />
                    <span className="truncate">{v}</span>
                </span>
            ) : <span className="text-slate-600">—</span>,
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
        <PageShell title="Manajemen Customer" onRefresh={fetchData}>
            <div className="flex items-center justify-between">
                <div>
                    <h2 className="text-base font-semibold text-white">Daftar Customer</h2>
                    <p className="text-xs text-slate-500 mt-0.5">{data.length} customer terdaftar</p>
                </div>
                <button onClick={openCreate} className="flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all">
                    <Plus size={16} /> Tambah Customer
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
                        searchKeys={["nama", "email", "telepon", "alamat"]}
                        emptyMessage="Belum ada customer. Klik 'Tambah Customer' untuk memulai."
                    />
                )}
            </div>

            <Modal open={modalOpen} onClose={() => setModalOpen(false)} title={editTarget ? "Edit Customer" : "Tambah Customer Baru"}>
                <form onSubmit={handleSave} className="space-y-4">
                    {formError && (
                        <div className="rounded-xl bg-rose-500/10 border border-rose-500/30 px-4 py-3 text-sm text-rose-400">{formError}</div>
                    )}
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Kode Customer <span className="text-rose-400">*</span></label>
                        <input required value={form.kode_cust} onChange={(e) => setForm({ ...form, kode_cust: e.target.value })} placeholder="CUST-001" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all font-mono" />
                    </div>
                    {[
                        { key: "nama", label: "Nama Lengkap", required: true, placeholder: "PT Maju Bersama" },
                        { key: "email", label: "Email", required: true, placeholder: "kontak@perusahaan.com", type: "email" },
                        { key: "telepon", label: "Nomor Telepon", required: false, placeholder: "021-12345678" },
                    ].map(({ key, label, required, placeholder, type = "text" }) => (
                        <div key={key}>
                            <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">
                                {label} {required && <span className="text-rose-400">*</span>}
                            </label>
                            <input type={type} required={required} value={form[key]} onChange={(e) => setForm({ ...form, [key]: e.target.value })} placeholder={placeholder} className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                        </div>
                    ))}
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Alamat</label>
                        <textarea rows={3} value={form.alamat} onChange={(e) => setForm({ ...form, alamat: e.target.value })} placeholder="Jl. Merdeka No. 1, Jakarta Pusat" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all resize-none" />
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
