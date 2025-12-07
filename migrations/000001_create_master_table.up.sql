-- File: xxxxxxxxxx_create_initial_tables.up.sql

-- Tabel yang tidak memiliki ketergantungan
CREATE TABLE "category" (
	"id"   SERIAL PRIMARY KEY,
	"name" TEXT NOT NULL
);

CREATE TABLE "role" (
	"id"        SERIAL PRIMARY KEY,
	"role_name" TEXT NOT NULL
);

CREATE TABLE "size" (
	"id"   SERIAL PRIMARY KEY,
	"name" TEXT NOT NULL
);

CREATE TABLE "status" (
	"id"   SERIAL PRIMARY KEY,
	"name" TEXT NOT NULL
);

CREATE TABLE "warehouse" (
	"id"                   SERIAL PRIMARY KEY,
	"warehouse_name"       TEXT NOT NULL UNIQUE,
	"warehouse_code"       TEXT NOT NULL UNIQUE,
	"location_description" TEXT
);

-- Tabel "employee" dibuat terakhir karena memiliki foreign key ke "role" dan "warehouse"
CREATE TABLE "employee" (
	"id"            SERIAL PRIMARY KEY,
	"user_id"       TEXT UNIQUE,
	"employee_name" TEXT,
	"password"      TEXT,
	"employee_code" TEXT UNIQUE,
	"id_role"       INTEGER NOT NULL,
	"warehouse_code"  TEXT NOT NULL,
	FOREIGN KEY ("id_role") REFERENCES "role" ("id") ON UPDATE CASCADE,
	FOREIGN KEY ("warehouse_code") REFERENCES "warehouse" ("warehouse_code") ON UPDATE CASCADE
);
