"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { isAuthenticated } from "@/lib/auth";
import { Loader2 } from "lucide-react";

/**
 * Wrap any page that requires authentication.
 * Redirects to /login if no token is found.
 */
export default function ProtectedRoute({ children }) {
    const router = useRouter();
    const [checking, setChecking] = useState(true);

    useEffect(() => {
        if (!isAuthenticated()) {
            router.replace("/login");
        } else {
            setChecking(false);
        }
    }, [router]);

    if (checking) {
        return (
            <div className="min-h-screen bg-slate-950 flex items-center justify-center">
                <Loader2 size={28} className="animate-spin text-indigo-400" />
            </div>
        );
    }

    return children;
}
