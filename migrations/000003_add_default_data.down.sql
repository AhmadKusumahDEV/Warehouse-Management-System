-- File: xxxxxxxxxx_seed_initial_data.down.sql

-- Hapus data dari tabel employee terlebih dahulu
DELETE FROM "employee" WHERE "employee_name" = 'resdox';

-- Hapus data dari tabel lainnya
DELETE FROM "status" WHERE "name" IN ('Failed', 'Pending', 'Completed');
DELETE FROM "size" WHERE "name" IN ('S', 'M', 'L', 'XL', '2XL');
DELETE FROM "warehouse" WHERE "warehouse_code" IN ('WH-01', 'WH-02', 'WH-03');
DELETE FROM "role" WHERE "role_name" IN ('employee', 'manager', 'admin', 'super admin');
DELETE FROM "category" WHERE "name" IN ('shirt', 'pants', 'shoes');