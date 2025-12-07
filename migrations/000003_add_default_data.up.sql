
INSERT INTO "category" ("name") VALUES
('shirt'),
('pants'),
('shoes');

INSERT INTO "role" ("role_name") VALUES
('employee'),
('manager'),
('admin'),
('super admin');

INSERT INTO "warehouse" ("warehouse_name", "warehouse_code") VALUES
('WH-01', 'WH-01'),
('WH-02', 'WH-02'),
('WH-03', 'WH-03');

INSERT INTO "size" ("name") VALUES
('S'),
('M'),
('L'),
('XL'),
('2XL');

INSERT INTO "status" ("name") VALUES
('Failed'),
('Pending'),
('Completed');


INSERT INTO "employee" ("user_id", "employee_name", "password", "employee_code", "id_role", "warehouse_code") VALUES
(
    'resdox-uid', -- Contoh User ID unik
    'resdox',
    'ganti_dengan_password_hash', -- Ganti dengan password yang sudah di-hash
    'SA-001', -- Contoh Kode Employee unik
    (SELECT "id" FROM "role" WHERE "role_name" = 'super admin'),
    'WH-01'
);