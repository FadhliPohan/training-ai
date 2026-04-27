"use client";

import { useState, useEffect, useCallback } from "react";
import { Plus, Pencil, Loader2, AlertCircle, ShieldCheck, User } from "lucide-react";

import ProtectedRoute from "@/components/ProtectedRoute";
import PageShell from "@/components/PageShell";
import DataTable from "@/components/DataTable";
import Modal from "@/components/Modal";
import { usersAPI } from "@/lib/api";

const EMPTY_FORM = {
    nama: "",
    email: "",
    password: "",
    role: "sales",
    telegram_user_id: "",
};

const ROLE_OPTIONS = [
    { value: "admin", label: "Admin", color: "text-violet-400 bg-violet-500/10 border-violet-500/20" },
    { value: "manager", label: "Manager", color: "text-indigo-400 bg-indigo-500/10 border-indigo-500/20" },
    { value: "sales", label: "Sales", color: "text-sky-400 bg-sky-500/10 border-sky-500/20" },
    { value: "viewer", label: "Viewer", color: "text-slate-400 bg-slate-500/10 border-slate-500/20" },
];

function RoleBadge({ role }) {
    const cfg = ROLE_OPTIONS.find((r) => r.value === role) ?? ROLE_OPTIONS[2];
    return (
        <span className={`inline-flex items-center gap-1 text-[11px] font-semibold px-2 py-0.5 rounded-full border ${cfg.color}`}>
            {role === "admin" ? <ShieldCheck size={10} /> : <User size={10} />}
            {cfg.label}
        </span>
    );
}

export default function UsersPage() {
    return (
        <ProtectedRoute>
            <UsersContent />
        </ProtectedRoute>
    );
}

function UsersContent() {
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
            const res = await usersAPI.list();
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
            nama: row.nama ?? "",
            email: row.email ?? "",
            password: "",
            role: row.role ?? "sales",
            telegram_user_id: row.telegram_user_id != null ? String(row.telegram_user_id) : "",
        });
        setFormError("");
        setModalOpen(true);
    };

    const handleSave = async (e) => {
        e.preventDefault();
        setFormError("");
        setSaving(true);
        try {
            const payload = {
                nama: form.nama,
                email: form.email,
                role: form.role,
                telegram_user_id: form.telegram_user_id ? Number(form.telegram_user_id) : null,
            };
            if (form.password) payload.password = form.password;

            if (editTarget) {
                await usersAPI.update(editTarget.id, payload);
            } else {
                await usersAPI.create(payload);
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
            render: (v) => <span className="text-slate-400 text-xs">{v}</span>,
        },
        {
            key: "role",
            label: "Role",
            render: (v) => <RoleBadge role={v} />,
        },
        {
            key: "telegram_user_id",
            label: "Telegram ID",
            render: (v) => v != null ? (
                <span className="text-xs font-mono text-emerald-400 bg-emerald-500/10 px-2 py-0.5 rounded border border-emerald-500/20">{v}</span>
            ) : (
                <span className="text-xs text-slate-600 italic">Belum terhubung</span>
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
        <PageShell title="Manajemen Users" onRefresh={fetchData}>
            <div className="flex items-center justify-between">
                <div>
                    <h2 className="text-base font-semibold text-white">Daftar Users</h2>
                    <p className="text-xs text-slate-500 mt-0.5">{data.length} user terdaftar</p>
                </div>
                <button onClick={openCreate} className="flex items-center gap-2 px-4 py-2.5 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all">
                    <Plus size={16} /> Tambah User
                </button>
            </div>

            {error && (
                <div className="flex items-center gap-2 rounded-xl bg-rose-500/10 border border-rose-500/30 px-4 py-3 text-sm text-rose-400">
                    <AlertCircle size={16} /> {error}
                </div>
            )}

            <div className="flex flex-wrap gap-2">
                {ROLE_OPTIONS.map((r) => (
                    <span key={r.value} className={`inline-flex items-center gap-1 text-[11px] font-medium px-2.5 py-1 rounded-full border ${r.color}`}>
                        {r.value === "admin" ? <ShieldCheck size={10} /> : <User size={10} />}
                        {r.label}
                    </span>
                ))}
                <span className="text-[11px] text-slate-500 self-center ml-1">· Telegram ID diperlukan untuk fitur Q&A Telegram</span>
            </div>

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
                        searchKeys={["nama", "email", "role"]}
                        emptyMessage="Belum ada user. Klik 'Tambah User' untuk memulai."
                    />
                )}
            </div>

            <Modal open={modalOpen} onClose={() => setModalOpen(false)} title={editTarget ? "Edit User" : "Tambah User Baru"}>
                <form onSubmit={handleSave} className="space-y-4">
                    {formError && (
                        <div className="rounded-xl bg-rose-500/10 border border-rose-500/30 px-4 py-3 text-sm text-rose-400">{formError}</div>
                    )}
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Nama Lengkap <span className="text-rose-400">*</span></label>
                        <input required value={form.nama} onChange={(e) => setForm({ ...form, nama: e.target.value })} placeholder="Budi Santoso" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Email <span className="text-rose-400">*</span></label>
                        <input required type="email" value={form.email} onChange={(e) => setForm({ ...form, email: e.target.value })} placeholder="budi@insightflow.id" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">
                            Password {!editTarget && <span className="text-rose-400">*</span>}
                            {editTarget && <span className="text-slate-600 normal-case font-normal ml-1">(kosongkan jika tidak diubah)</span>}
                        </label>
                        <input required={!editTarget} type="password" value={form.password} onChange={(e) => setForm({ ...form, password: e.target.value })} placeholder={editTarget ? "••••••••" : "Min. 6 karakter"} className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all" />
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Role <span className="text-rose-400">*</span></label>
                        <div className="grid grid-cols-2 gap-2">
                            {ROLE_OPTIONS.map((r) => (
                                <button key={r.value} type="button" onClick={() => setForm({ ...form, role: r.value })} className={`py-2 rounded-xl text-xs font-semibold border transition-all ${form.role === r.value ? r.color : "border-slate-700 text-slate-500 hover:border-slate-600 hover:text-slate-300"}`}>
                                    {r.label}
                                </button>
                            ))}
                        </div>
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1.5">Telegram User ID</label>
                        <input value={form.telegram_user_id} onChange={(e) => setForm({ ...form, telegram_user_id: e.target.value })} placeholder="123456789" className="w-full rounded-xl bg-slate-800/60 border border-slate-700 text-white placeholder-slate-500 px-4 py-2.5 text-sm outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500/20 transition-all font-mono" />
                        <p className="text-[11px] text-slate-600 mt-1">Diperlukan agar user dapat menggunakan fitur Q&A via Telegram.</p>
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
