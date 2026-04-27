import { Loader2 } from "lucide-react";

export default function Loading() {
    return (
        <div className="min-h-screen bg-slate-950 flex items-center justify-center">
            <div className="text-center space-y-4">
                <div className="relative inline-block">
                    <Loader2 size={40} className="animate-spin text-indigo-400" />
                    <div className="absolute inset-0 blur-xl opacity-50 bg-indigo-600" />
                </div>
                <p className="text-sm text-slate-400">Memuat...</p>
            </div>
        </div>
    );
}
