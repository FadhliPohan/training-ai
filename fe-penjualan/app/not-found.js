import Link from "next/link";
import { Home, ArrowLeft } from "lucide-react";

export default function NotFound() {
    return (
        <div className="min-h-screen bg-slate-950 flex items-center justify-center p-4">
            <div className="text-center space-y-6 max-w-md">
                {/* 404 */}
                <div className="relative">
                    <h1 className="text-9xl font-bold text-transparent bg-clip-text bg-gradient-to-br from-indigo-400 to-purple-600">
                        404
                    </h1>
                    <div className="absolute inset-0 blur-3xl opacity-30 bg-gradient-to-br from-indigo-600 to-purple-600" />
                </div>

                {/* Message */}
                <div className="space-y-2">
                    <h2 className="text-2xl font-bold text-white">Halaman Tidak Ditemukan</h2>
                    <p className="text-slate-400 text-sm">
                        Maaf, halaman yang Anda cari tidak ada atau telah dipindahkan.
                    </p>
                </div>

                {/* Actions */}
                <div className="flex flex-col sm:flex-row gap-3 justify-center">
                    <Link
                        href="/"
                        className="inline-flex items-center justify-center gap-2 px-5 py-3 rounded-xl text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all"
                    >
                        <Home size={16} />
                        Kembali ke Dashboard
                    </Link>
                    <button
                        onClick={() => window.history.back()}
                        className="inline-flex items-center justify-center gap-2 px-5 py-3 rounded-xl text-sm font-medium border border-slate-700 text-slate-400 hover:text-white hover:bg-slate-700/40 transition-all"
                    >
                        <ArrowLeft size={16} />
                        Halaman Sebelumnya
                    </button>
                </div>

                {/* Footer */}
                <p className="text-xs text-slate-600 pt-6">
                    InsightFlow · AI Sales Dashboard
                </p>
            </div>
        </div>
    );
}
