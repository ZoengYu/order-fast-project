CREATE TABLE IF NOT EXISTS stores (
  "id" bigserial PRIMARY KEY,
  "store_name" varchar(60) NOT NULL,
  "store_address" varchar(120) NOT NULL,
  "store_phone" varchar(10) UNIQUE NOT NULL,
  "store_owner" varchar NOT NULL,
  "store_manager" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS tables (
  "id" bigserial PRIMARY KEY,
  "store_id" bigint NOT NULL,
  "table_id" bigint NOT NULL,
  "table_name" varchar(60)  NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tables" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
CREATE INDEX ON "stores" ("store_name");

CREATE INDEX ON "tables" ("store_id");
