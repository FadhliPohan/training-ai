DELETE FROM bisnis.tbl_produk WHERE kode_produk LIKE 'KAO-%' OR kode_produk LIKE 'CEM-%' OR kode_produk LIKE 'JKT-%' OR kode_produk LIKE 'DRS-%' OR kode_produk LIKE 'TAS-%';
DELETE FROM app.users WHERE email = 'admin@insightflow.id';
DELETE FROM app.anomaly_config;
