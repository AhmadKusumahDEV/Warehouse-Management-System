
CREATE TABLE "product" (
	"id"                  SERIAL PRIMARY KEY,
	"product_name"        TEXT NOT NULL,
	"price"               INTEGER NOT NULL,
	"description_product" TEXT,
	"product_code"        TEXT NOT NULL UNIQUE,
	"id_category"         INTEGER NOT NULL,
	FOREIGN KEY ("id_category") REFERENCES "category" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE "transactions" (
	"id"                      SERIAL PRIMARY KEY,
	"code_transaksi"          TEXT NOT NULL UNIQUE,
	"origin_entity_name"      TEXT NOT NULL,
	"destination_entity_name" TEXT,
	"employee_code"           TEXT NOT NULL,
	"id_status"               INTEGER NOT NULL,
	"created_at"              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Menggunakan TIMESTAMPTZ untuk praktik terbaik
	"tipe_transaksi"          TEXT,
	FOREIGN KEY ("employee_code") REFERENCES "employee" ("employee_code") ON DELETE RESTRICT,
	FOREIGN KEY ("id_status") REFERENCES "status" ("id") ON DELETE RESTRICT
);

-- Tabel yang bergantung pada tabel yang baru dibuat di atas
CREATE TABLE "product_detail" (
	"id"           SERIAL PRIMARY KEY,
	"code_product" TEXT NOT NULL,
	"id_size"      INTEGER NOT NULL,
	"barcode"      TEXT NOT NULL UNIQUE,
	UNIQUE ("code_product", "id_size"),
	FOREIGN KEY ("code_product") REFERENCES "product" ("product_code") ON DELETE CASCADE,
	FOREIGN KEY ("id_size") REFERENCES "size" ("id") ON DELETE RESTRICT
);

CREATE TABLE "inventory" (
	"id"             SERIAL PRIMARY KEY,
	"quantity"       INTEGER NOT NULL DEFAULT 0,
	"code_product"   TEXT NOT NULL,
	"id_size"        INTEGER NOT NULL,
	"code_warehouse" TEXT NOT NULL,
	UNIQUE ("code_product", "id_size", "code_warehouse"),
	FOREIGN KEY ("code_product") REFERENCES "product" ("product_code") ON DELETE CASCADE,
	FOREIGN KEY ("code_warehouse") REFERENCES "warehouse" ("warehouse_code") ON DELETE CASCADE,
	FOREIGN KEY ("id_size") REFERENCES "size" ("id") ON DELETE RESTRICT
);

-- Tabel yang paling banyak memiliki ketergantungan
CREATE TABLE "detail_transactions" (
	"id"                SERIAL PRIMARY KEY,
	"id_transaction"    INTEGER NOT NULL,
	"id_detail_product" INTEGER NOT NULL,
	"quantity"          INTEGER NOT NULL,
	"scanner_quantity"  INTEGER NOT NULL DEFAULT 0,
	UNIQUE ("id_transaction", "id_detail_product"),
	FOREIGN KEY ("id_detail_product") REFERENCES "product_detail" ("id") ON DELETE RESTRICT,
	FOREIGN KEY ("id_transaction") REFERENCES "transactions" ("id") ON DELETE CASCADE
);