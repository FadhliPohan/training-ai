"use client";

import { useEffect } from "react";
import { X } from "lucide-react";

/**
 * Generic modal overlay.
 * Props: open, onClose, title, children, size ("sm"|"md"|"lg")
 */
export default function Modal({ open, onClose, title, children, size = "md" }) {
    // Close on Escape
    useEffect(() => {
        if (!open) return;
        const handler = (e) => { if (e.key === "Escape") onClose(); };
        window.addEventListener("keydown", handler);
        return () => window.removeEventListener("keydown", handler);
    }, [open, onClose]);

    if (!open) return null;

    const widthClass = {
        sm: "max-w-sm",
        md: "max-w-lg",
        lg: "max-w-2xl",
    }[size] ?? "max-w-lg";

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            {/* Backdrop */}
            <div
                className="absolute inset-0 bg-black/60 backdrop-blur-sm"
                onClick={onClose}
            />

            {/* Panel */}
            <div
                className={`relative w-full ${widthClass} bg-[#1e293b] border border-slate-700/60 rounded-2xl shadow-2xl shadow-black/50 animate-fade-in`}
            >
                {/* Header */}
                <div className="flex items-center justify-between px-6 py-4 border-b border-slate-700/50">
                    <h2 className="text-base font-semibold text-white">{title}</h2>
                    <button
                        onClick={onClose}
                        className="p-1.5 rounded-lg text-slate-400 hover:text-white hover:bg-slate-700/50 transition-all"
                    >
                        <X size={16} />
                    </button>
                </div>

                {/* Body */}
                <div className="px-6 py-5">{children}</div>
            </div>
        </div>
    );
}
