"use client";

import { useEffect } from "react";
import { AlertTriangle, RefreshCw, Home } from "lucide-react";
import Link from "next/link";

export default function Error({ error, reset }) {
    useEffect(() => {
        console.error("Error boundary caught:", error);
    }, [error]);

    return (
        <div className="min-h-screen bg-slate-950 flex items-center justify-center p-4">
            <div className="text-center space-y-6 max-w-md">
                {/* Icon */}
                <div className="inline-flex items-center justify-center w-20 h-20 rounded-2xl bg-rose-500/10 border border-rose-500/30">
                    <AlertTriangle size={40} className="text-rose-400" />
                </div>

                {/* Message */}
                <div className="space-y-2">
                    <h2 className="text-2xl font-bold text-white">Terjadi Kesalahan</h2>
                    <p className="text-slate-400 text-sm">
                        Maaf, terjadi kesalahan saat memuat halaman ini.
                    </p>
                    {error?.message && (
                        <div className="mt-4 p-3 rounded-xl bg-slate-800/60 border border-slate-700">
                            <p className="text-xs text-slate-500 font-mono break-all">
                                {error.message}
                            </p>
                        </div>
                    )}
                </div>

                {/* Actions */}
                <div className="flex flex-col sm:flex-row gap-3 justify-center">
                    <button
                        onClick={reset}
                        className="inline-flex items-center justify-center gap-2 px-5 py-3 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all"
                    >
                        <RefreshCw size={16} />
                        Coba Lagi
                    </button>
                    <Link
                        href="/"
                        className="inline-flex items-center justify-center gap-2 px-5 py-3 rounded-xl text-sm font-medium border border-slate-700 text-slate-400 hover:text-white hover:bg-slate-700/40 transition-all"
                    >
                        <Home size={16} />
                        Kembali ke Dashboard
                    </Link>
                </div>

                {/* Footer */}
                <p className="text-xs text-slate-600 pt-6">
                    Jika masalah berlanjut, hubungi administrator sistem.
                </p>
            </div>
        </div>
    );
}
