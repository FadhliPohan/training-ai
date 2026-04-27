import "./globals.css";
import { SidebarProvider } from "@/lib/sidebar-context";

export const metadata = {
  title: "InsightFlow — AI Sales Dashboard",
  description: "Self-service AI analytics dashboard untuk penjualan pakaian InsightFlow",
};

export default function RootLayout({ children }) {
  return (
    <html lang="id" className="dark">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
        <link
          href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap"
          rel="stylesheet"
        />
      </head>
      <body className="bg-slate-950 text-slate-100 min-h-screen antialiased">
        <SidebarProvider>
          {children}
        </SidebarProvider>
      </body>
    </html>
  );
}
