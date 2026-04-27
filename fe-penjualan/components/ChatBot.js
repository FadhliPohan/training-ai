"use client";

import { useState, useRef, useEffect, useCallback } from "react";
import { chatAPI } from "@/lib/api";

// ─── Suggested Prompts ─────────────────────────────────────────────────────
const SUGGESTED_PROMPTS = [
  { icon: "📈", text: "Bagaimana tren penjualan bulan ini?" },
  { icon: "⚠️", text: "Anomali apa yang perlu diperhatikan?" },
  { icon: "🏆", text: "Produk mana yang paling laris?" },
  { icon: "👥", text: "Bagaimana performa tim sales saya?" },
  { icon: "📊", text: "Buat ringkasan laporan penjualan" },
  { icon: "💡", text: "Rekomendasi untuk meningkatkan omzet" },
];

// ─── Message Bubble ────────────────────────────────────────────────────────
function MessageBubble({ message }) {
  const isUser = message.role === "user";
  const isTyping = message.typing;

  return (
    <div
      className={`flex gap-2.5 items-end ${isUser ? "flex-row-reverse" : "flex-row"}`}
    >
      {/* Avatar */}
      {!isUser && (
        <div className="flex-shrink-0 w-7 h-7 rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center text-xs shadow-lg shadow-indigo-500/30">
          🤖
        </div>
      )}

      {/* Bubble */}
      <div
        className={`
          max-w-[85%] rounded-2xl px-4 py-2.5 text-sm leading-relaxed shadow-sm
          ${isUser
            ? "bg-indigo-600 text-white rounded-br-sm shadow-indigo-600/20"
            : "bg-slate-800 text-slate-100 border border-slate-700/50 rounded-bl-sm"
          }
        `}
      >
        {isTyping ? (
          <div className="flex items-center gap-1 py-1">
            <span className="w-1.5 h-1.5 rounded-full bg-slate-400 animate-bounce" style={{ animationDelay: "0ms" }} />
            <span className="w-1.5 h-1.5 rounded-full bg-slate-400 animate-bounce" style={{ animationDelay: "150ms" }} />
            <span className="w-1.5 h-1.5 rounded-full bg-slate-400 animate-bounce" style={{ animationDelay: "300ms" }} />
          </div>
        ) : (
          <span className="whitespace-pre-wrap">{message.content}</span>
        )}
      </div>

      {/* User avatar */}
      {isUser && (
        <div className="flex-shrink-0 w-7 h-7 rounded-full bg-slate-700 border border-slate-600 flex items-center justify-center text-xs">
          👤
        </div>
      )}
    </div>
  );
}

// ─── Main ChatBot Component ────────────────────────────────────────────────
export default function ChatBot() {
  const [isOpen, setIsOpen] = useState(false);
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [hasUnread, setHasUnread] = useState(false);
  const [showSuggestions, setShowSuggestions] = useState(true);
  const [error, setError] = useState(null);

  const messagesEndRef = useRef(null);
  const inputRef = useRef(null);
  const chatWindowRef = useRef(null);

  // Auto-scroll to bottom when new messages arrive
  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);

  // Focus input when chat opens
  useEffect(() => {
    if (isOpen && inputRef.current) {
      setTimeout(() => inputRef.current?.focus(), 150);
    }
  }, [isOpen]);

  // Unread badge when chat is closed and AI replies
  useEffect(() => {
    if (!isOpen && messages.length > 0) {
      const lastMsg = messages[messages.length - 1];
      if (lastMsg.role === "assistant" && !lastMsg.typing) {
        setHasUnread(true);
      }
    }
  }, [messages, isOpen]);

  // Clear unread when opened
  const handleOpen = () => {
    setIsOpen(true);
    setHasUnread(false);
  };

  // Send a message
  const sendMessage = useCallback(async (text) => {
    const trimmed = (text || input).trim();
    if (!trimmed || isLoading) return;

    setError(null);
    setShowSuggestions(false);
    setInput("");

    // Append user message
    const userMsg = { role: "user", content: trimmed };
    setMessages((prev) => [...prev, userMsg]);

    // Show typing indicator
    const typingId = Date.now();
    setMessages((prev) => [
      ...prev,
      { role: "assistant", content: "", typing: true, id: typingId },
    ]);
    setIsLoading(true);

    try {
      // Build conversation history (excluding typing indicator)
      const history = [...messages, userMsg].map(({ role, content }) => ({
        role,
        content,
      }));

      const res = await chatAPI.chat(history);
      const reply = res?.data?.reply || "Maaf, tidak ada balasan dari AI.";

      // Replace typing indicator with real response
      setMessages((prev) =>
        prev.map((m) =>
          m.id === typingId
            ? { role: "assistant", content: reply, typing: false }
            : m
        )
      );
    } catch (err) {
      setMessages((prev) => prev.filter((m) => m.id !== typingId));
      setError(err.message || "Gagal menghubungi AI. Coba lagi.");
    } finally {
      setIsLoading(false);
    }
  }, [input, isLoading, messages]);

  // Handle Enter key
  const handleKeyDown = (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      sendMessage();
    }
  };

  // Clear conversation
  const clearChat = () => {
    setMessages([]);
    setShowSuggestions(true);
    setError(null);
    setInput("");
  };

  const hasMessages = messages.length > 0;

  return (
    <>
      {/* ── Floating Toggle Button ── */}
      <button
        id="chatbot-toggle-btn"
        onClick={isOpen ? () => setIsOpen(false) : handleOpen}
        aria-label="Buka AI Chat Assistant"
        className={`
          fixed bottom-6 right-6 z-50
          w-14 h-14 rounded-full
          flex items-center justify-center
          shadow-xl shadow-indigo-600/40
          transition-all duration-300 ease-out
          ${isOpen
            ? "bg-slate-700 hover:bg-slate-600 rotate-0 scale-95"
            : "bg-gradient-to-br from-indigo-500 to-purple-600 hover:from-indigo-400 hover:to-purple-500 hover:scale-110"
          }
        `}
      >
        <span
          className={`text-xl transition-all duration-300 ${isOpen ? "opacity-0 scale-0 absolute" : "opacity-100 scale-100"}`}
        >
          💬
        </span>
        <span
          className={`text-xl transition-all duration-300 ${isOpen ? "opacity-100 scale-100" : "opacity-0 scale-0 absolute"}`}
        >
          ✕
        </span>

        {/* Unread badge */}
        {hasUnread && !isOpen && (
          <span className="absolute -top-1 -right-1 w-4 h-4 rounded-full bg-rose-500 text-[9px] text-white font-bold flex items-center justify-center animate-pulse">
            !
          </span>
        )}

        {/* Pulse ring when closed */}
        {!isOpen && !hasUnread && (
          <span className="absolute inset-0 rounded-full bg-indigo-500/30 animate-ping" />
        )}
      </button>

      {/* ── Chat Window ── */}
      <div
        ref={chatWindowRef}
        id="chatbot-window"
        className={`
          fixed bottom-24 right-6 z-50
          w-[380px] max-w-[calc(100vw-2rem)]
          rounded-2xl shadow-2xl shadow-black/40
          border border-slate-700/60
          overflow-hidden
          transition-all duration-300 ease-out origin-bottom-right
          ${isOpen
            ? "opacity-100 scale-100 translate-y-0 pointer-events-auto"
            : "opacity-0 scale-95 translate-y-4 pointer-events-none"
          }
        `}
        style={{ background: "linear-gradient(180deg, #1a2234 0%, #151d2e 100%)" }}
      >
        {/* ── Header ── */}
        <div className="px-4 py-3 border-b border-slate-700/50 flex items-center gap-3"
          style={{ background: "linear-gradient(135deg, rgba(99,102,241,0.12) 0%, rgba(139,92,246,0.08) 100%)" }}
        >
          {/* AI Avatar */}
          <div className="relative flex-shrink-0">
            <div className="w-9 h-9 rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center text-base shadow-lg shadow-indigo-500/40">
              🤖
            </div>
            <span className="absolute bottom-0 right-0 w-2.5 h-2.5 rounded-full bg-emerald-400 border-2 border-slate-800" />
          </div>

          {/* Title */}
          <div className="flex-1 min-w-0">
            <h3 className="text-sm font-semibold text-white leading-tight">InsightFlow AI</h3>
            <p className="text-[10px] text-emerald-400 font-medium">● Online · GPT-4o</p>
          </div>

          {/* Actions */}
          <div className="flex items-center gap-1">
            {hasMessages && (
              <button
                id="chatbot-clear-btn"
                onClick={clearChat}
                title="Hapus percakapan"
                className="w-7 h-7 rounded-lg flex items-center justify-center text-slate-500 hover:text-slate-300 hover:bg-slate-700/60 transition-all text-sm"
              >
                🗑
              </button>
            )}
            <button
              id="chatbot-close-btn"
              onClick={() => setIsOpen(false)}
              className="w-7 h-7 rounded-lg flex items-center justify-center text-slate-500 hover:text-slate-300 hover:bg-slate-700/60 transition-all text-sm"
            >
              ✕
            </button>
          </div>
        </div>

        {/* ── Messages Area ── */}
        <div className="h-[360px] overflow-y-auto p-4 space-y-3 scrollbar-thin">

          {/* Welcome / Suggestions */}
          {!hasMessages && showSuggestions && (
            <div className="space-y-4 animate-fade-in">
              {/* Welcome card */}
              <div className="rounded-xl border border-indigo-500/20 bg-indigo-500/5 p-3 text-center">
                <div className="text-2xl mb-1.5">🤖</div>
                <p className="text-sm font-semibold text-white">Halo! Saya InsightFlow AI</p>
                <p className="text-xs text-slate-400 mt-1 leading-relaxed">
                  Asisten analitik penjualan Anda. Tanya apa saja seputar data penjualan, monitoring, dan performa bisnis.
                </p>
              </div>

              {/* Suggestion chips */}
              <div>
                <p className="text-[10px] text-slate-500 uppercase tracking-wider font-semibold mb-2">
                  Pertanyaan cepat
                </p>
                <div className="grid grid-cols-2 gap-1.5">
                  {SUGGESTED_PROMPTS.map((prompt, i) => (
                    <button
                      key={i}
                      id={`chatbot-suggestion-${i}`}
                      onClick={() => sendMessage(prompt.text)}
                      className="flex items-start gap-1.5 p-2 rounded-xl text-left text-xs text-slate-300 bg-slate-800/60 border border-slate-700/40 hover:border-indigo-500/40 hover:bg-indigo-500/5 hover:text-white transition-all duration-200 leading-tight"
                    >
                      <span className="text-sm flex-shrink-0">{prompt.icon}</span>
                      <span>{prompt.text}</span>
                    </button>
                  ))}
                </div>
              </div>
            </div>
          )}

          {/* Message history */}
          {messages.map((msg, i) => (
            <MessageBubble key={i} message={msg} />
          ))}

          {/* Error state */}
          {error && (
            <div className="rounded-xl border border-rose-500/30 bg-rose-500/5 p-3 text-center animate-fade-in">
              <p className="text-xs text-rose-400">⚠️ {error}</p>
              <button
                onClick={() => setError(null)}
                className="text-[10px] text-rose-400/70 hover:text-rose-300 mt-1 underline"
              >
                Tutup
              </button>
            </div>
          )}

          {/* Scroll anchor */}
          <div ref={messagesEndRef} />
        </div>

        {/* ── Input Area ── */}
        <div className="px-3 py-3 border-t border-slate-700/50 bg-slate-900/40">
          <div className="flex items-end gap-2">
            <textarea
              ref={inputRef}
              id="chatbot-input"
              value={input}
              onChange={(e) => setInput(e.target.value)}
              onKeyDown={handleKeyDown}
              placeholder="Tanya seputar penjualan..."
              disabled={isLoading}
              rows={1}
              className="
                flex-1 resize-none rounded-xl px-3.5 py-2.5 text-sm
                bg-slate-800/80 border border-slate-700/50
                text-white placeholder-slate-500
                focus:outline-none focus:border-indigo-500/60 focus:ring-1 focus:ring-indigo-500/30
                transition-all duration-200
                disabled:opacity-50 disabled:cursor-not-allowed
                min-h-[40px] max-h-[120px] overflow-y-auto
              "
              style={{ lineHeight: "1.4" }}
            />
            <button
              id="chatbot-send-btn"
              onClick={() => sendMessage()}
              disabled={!input.trim() || isLoading}
              className="
                flex-shrink-0 w-10 h-10 rounded-xl
                flex items-center justify-center
                bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-700
                disabled:opacity-40 disabled:cursor-not-allowed
                shadow-lg shadow-indigo-600/30
                transition-all duration-200 hover:scale-105 active:scale-95
              "
            >
              {isLoading ? (
                <span className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin" />
              ) : (
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-4 h-4 text-white">
                  <path d="M3.478 2.405a.75.75 0 00-.926.94l2.432 7.905H13.5a.75.75 0 010 1.5H4.984l-2.432 7.905a.75.75 0 00.926.94 60.519 60.519 0 0018.445-8.986.75.75 0 000-1.218A60.517 60.517 0 003.478 2.405z" />
                </svg>
              )}
            </button>
          </div>

          {/* Footer hint */}
          <p className="text-[10px] text-slate-600 text-center mt-2">
            Enter untuk kirim · Shift+Enter baris baru · Powered by GPT-4o
          </p>
        </div>
      </div>
    </>
  );
}
