CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar(60) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS stores (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "store_name" varchar(60) NOT NULL,
  "store_address" varchar(120) NOT NULL,
  "store_phone" varchar(10) UNIQUE NOT NULL,
  "store_manager" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS tables (
  "id" bigserial PRIMARY KEY,
  "store_id" bigint NOT NULL,
  "table_id" bigint NOT NULL,
  "table_name" varchar(60) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS menu (
  "id" bigserial PRIMARY KEY,
  "store_id" bigint NOT NULL,
  "menu_name" varchar(60) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS item (
  "id" bigserial PRIMARY KEY,
  "menu_id" bigint NOT NULL,
  "name" varchar(60) NOT NULL,
  "price" integer NOT NULL
);

CREATE TABLE IF NOT EXISTS item_tag (
  "id" bigserial PRIMARY KEY,
  "item_id" bigint NOT NULL,
  "item_tag" varchar(60) NOT NULL
);

ALTER TABLE "stores" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "tables" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
ALTER TABLE "item" ADD FOREIGN KEY ("menu_id") REFERENCES "menu" ("id");
ALTER TABLE "item_tag" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("id");

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "stores" ("store_name");

CREATE INDEX ON "tables" ("store_id");
