"use client";

import Sidebar from "@/components/Sidebar";
import Topbar from "@/components/Topbar";
import { useSidebar } from "@/lib/sidebar-context";

/**
 * PageShell — shared layout wrapper for every protected page.
 *
 * Reads sidebar collapsed state from context and applies the correct
 * left margin so content is never covered by the fixed sidebar.
 *
 * Props:
 *  - title (string): page title shown in Topbar
 *  - onRefresh (function): optional refresh callback for Topbar
 *  - children: page content
 */
export default function PageShell({ title, onRefresh, children }) {
    const { collapsed } = useSidebar();

    return (
        <div className="flex min-h-screen bg-slate-950">
            <Sidebar />

            {/*
        Content area shifts right to clear the fixed sidebar.
        - Desktop expanded  (collapsed=false): sidebar = 240px → ml-[240px]
        - Desktop collapsed (collapsed=true):  sidebar =  72px → ml-[72px]
        - Mobile: sidebar overlays (no margin needed, handled by backdrop)
        transition-all keeps the shift smooth when toggling.
      */}
            <div
                className={`
          flex-1 flex flex-col min-w-0
          transition-all duration-300 ease-in-out
          ${collapsed ? "lg:ml-[72px]" : "lg:ml-[240px]"}
        `}
            >
                <Topbar title={title} onRefresh={onRefresh} />
                <main className="flex-1 p-6 space-y-6 overflow-y-auto">
                    {children}
                </main>
            </div>
        </div>
    );
}
