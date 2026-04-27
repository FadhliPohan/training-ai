"use client";

/**
 * Auth helpers — token stored in localStorage.
 * Used by the protected route HOC and API layer.
 */

export function getToken() {
    if (typeof window === "undefined") return null;
    return localStorage.getItem("token");
}

export function getUser() {
    if (typeof window === "undefined") return null;
    try {
        return JSON.parse(localStorage.getItem("user") || "null");
    } catch {
        return null;
    }
}

export function isAuthenticated() {
    return !!getToken();
}

export function logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
}

export function isAdmin() {
    const user = getUser();
    return user?.role === "admin";
}
