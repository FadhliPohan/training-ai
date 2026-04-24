-- Migration: 004_seed_initial_data
-- Seeds default admin user, anomaly config defaults, and sample product data for development.

-- ============================================================
-- Default anomaly config (per-metric thresholds)
-- ============================================================
INSERT INTO app.anomaly_config (metric_key, threshold_pct, aktif) VALUES
    ('daily_revenue',  10.00, true),
    ('order_count',    15.00, true),
    ('cancelled_rate', 20.00, true),
    ('low_stock',       5.00, true)  -- alert when stock drops more than 5 units below avg
ON CONFLICT (metric_key) DO NOTHING;

-- ============================================================
-- Default admin user
-- Password: Admin@12345  (bcrypt hash generated with cost=12)
-- CHANGE THIS PASSWORD immediately after first login!
-- ============================================================
INSERT INTO app.users (id, nama, email, password, role, aktif) VALUES
    (
        gen_random_uuid(),
        'Administrator',
        'admin@insightflow.id',
        '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQyCkJ0JRcMQjf7IcjnYT5a.m',
        'admin',
        true
    )
ON CONFLICT (email) DO NOTHING;

-- ============================================================
-- Sample products (development only)
-- ============================================================
INSERT INTO bisnis.tbl_produk (kode_produk, nama, kategori_pakaian, ukuran, warna, bahan, harga, stok) VALUES
    ('KAO-001-M-HTM', 'Kaos Polos Cotton Combed 30s', 'atasan', 'M', 'Hitam', 'Katun Combed 30s', 89000, 50),
    ('KAO-001-L-HTM', 'Kaos Polos Cotton Combed 30s', 'atasan', 'L', 'Hitam', 'Katun Combed 30s', 89000, 40),
    ('KAO-001-M-PTH', 'Kaos Polos Cotton Combed 30s', 'atasan', 'M', 'Putih', 'Katun Combed 30s', 89000, 35),
    ('KAO-001-L-PTH', 'Kaos Polos Cotton Combed 30s', 'atasan', 'L', 'Putih', 'Katun Combed 30s', 89000, 30),
    ('CEM-001-M-ABU', 'Celana Chino Slim Fit', 'bawahan', 'M', 'Abu-abu', 'Cotton Twill', 175000, 25),
    ('CEM-001-L-NVY', 'Celana Chino Slim Fit', 'bawahan', 'L', 'Navy', 'Cotton Twill', 175000, 20),
    ('JKT-001-L-HTM', 'Jaket Bomber Basic', 'outerwear', 'L', 'Hitam', 'Polyester', 250000, 15),
    ('DRS-001-M-FLR', 'Dress Floral Casual', 'dress', 'M', 'Floral', 'Rayon', 195000, 12),
    ('DRS-001-S-FLR', 'Dress Floral Casual', 'dress', 'S', 'Floral', 'Rayon', 195000, 8),
    ('TAS-001-AS-HTM', 'Tote Bag Canvas', 'aksesoris', 'All Size', 'Hitam', 'Canvas', 75000, 30)
ON CONFLICT (kode_produk) DO NOTHING;
