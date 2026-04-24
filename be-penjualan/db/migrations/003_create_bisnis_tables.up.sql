-- Migration: 003_create_bisnis_tables
-- Creates all transactional tables for the sales business domain.

-- Master product catalogue (clothing items)
-- Decision: one product = one row with ukuran+warna as fields (not separate SKUs per combination)
CREATE TABLE bisnis.tbl_produk (
    id               SERIAL PRIMARY KEY,
    kode_produk      VARCHAR(50)    UNIQUE NOT NULL,
    nama             VARCHAR(200)   NOT NULL,
    kategori_pakaian VARCHAR(100)   NOT NULL,   -- atasan | bawahan | dress | outerwear | aksesoris
    ukuran           VARCHAR(20)    NOT NULL,   -- S | M | L | XL | XXL | All Size
    warna            VARCHAR(50)    NOT NULL,
    bahan            VARCHAR(100),              -- Katun | Polyester | Denim | dll
    harga            DECIMAL(15,2)  NOT NULL CHECK (harga >= 0),
    stok             INTEGER        NOT NULL DEFAULT 0 CHECK (stok >= 0),
    aktif            BOOLEAN        NOT NULL DEFAULT TRUE
);

-- Master customer (customers can self-register — answered in clarification #2)
-- NOTE: customer auth is separate from app.users; customers are external buyers
CREATE TABLE bisnis.tbl_customer (
    id         SERIAL PRIMARY KEY,
    kode_cust  VARCHAR(50)  UNIQUE NOT NULL,
    nama       VARCHAR(200) NOT NULL,
    email      VARCHAR(150) UNIQUE,
    telepon    VARCHAR(20),
    alamat     TEXT,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- Order header
CREATE TABLE bisnis.tbl_order (
    id          SERIAL PRIMARY KEY,
    no_order    VARCHAR(50)   UNIQUE NOT NULL,
    customer_id INTEGER       NOT NULL REFERENCES bisnis.tbl_customer(id),
    sales_id    UUID          NOT NULL REFERENCES app.users(id),
    tanggal     DATE          NOT NULL DEFAULT CURRENT_DATE,
    status      VARCHAR(30)   NOT NULL DEFAULT 'pending'
                              CHECK (status IN ('pending','confirmed','paid','shipped','closed','cancelled')),
    alasan_batal TEXT,                            -- populated when status = cancelled
    total       DECIMAL(15,2) NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

-- Order line items
CREATE TABLE bisnis.tbl_order_detail (
    id         SERIAL PRIMARY KEY,
    order_id   INTEGER        NOT NULL REFERENCES bisnis.tbl_order(id) ON DELETE CASCADE,
    produk_id  INTEGER        NOT NULL REFERENCES bisnis.tbl_produk(id),
    qty        INTEGER        NOT NULL CHECK (qty > 0),
    harga_saat DECIMAL(15,2)  NOT NULL CHECK (harga_saat >= 0),  -- price snapshot at order time
    subtotal   DECIMAL(15,2)  NOT NULL CHECK (subtotal >= 0)
);

-- Payment records
CREATE TABLE bisnis.tbl_pembayaran (
    id       SERIAL PRIMARY KEY,
    order_id INTEGER        NOT NULL REFERENCES bisnis.tbl_order(id),
    jumlah   DECIMAL(15,2)  NOT NULL CHECK (jumlah > 0),
    metode   VARCHAR(50)    NOT NULL CHECK (metode IN ('transfer','tunai','kartu')),
    status   VARCHAR(30)    NOT NULL DEFAULT 'pending'
             CHECK (status IN ('pending','verified','rejected')),
    tanggal  TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

-- Shipment records
CREATE TABLE bisnis.tbl_pengiriman (
    id       SERIAL PRIMARY KEY,
    order_id INTEGER        NOT NULL REFERENCES bisnis.tbl_order(id),
    kurir    VARCHAR(100)   NOT NULL,
    no_resi  VARCHAR(100),
    status   VARCHAR(30)    NOT NULL DEFAULT 'proses'
             CHECK (status IN ('proses','dikirim','diterima')),
    tanggal  TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

-- ============================================================
-- INDEXES
-- ============================================================

-- Orders: most queries filter by date range, sales, or status
CREATE INDEX idx_order_tanggal    ON bisnis.tbl_order (tanggal);
CREATE INDEX idx_order_sales_id   ON bisnis.tbl_order (sales_id);
CREATE INDEX idx_order_status     ON bisnis.tbl_order (status);
CREATE INDEX idx_order_customer   ON bisnis.tbl_order (customer_id);

-- Products: dashboard reports often filter by category
CREATE INDEX idx_produk_kategori  ON bisnis.tbl_produk (kategori_pakaian);
CREATE INDEX idx_produk_aktif     ON bisnis.tbl_produk (aktif);

-- Order details: JOIN lookups
CREATE INDEX idx_order_detail_order   ON bisnis.tbl_order_detail (order_id);
CREATE INDEX idx_order_detail_produk  ON bisnis.tbl_order_detail (produk_id);

-- Payments: frequently queried by order
CREATE INDEX idx_pembayaran_order  ON bisnis.tbl_pembayaran (order_id);
