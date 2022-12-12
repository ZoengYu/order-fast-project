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

CREATE TABLE IF NOT EXISTS food (
  "id" bigserial PRIMARY KEY,
  "menu_id" bigint NOT NULL,
  "name" varchar(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS food_tag (
  "id" bigserial PRIMARY KEY,
  "food_id" bigint NOT NULL,
  "food_tag" varchar(60) NOT NULL
);

ALTER TABLE "tables" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
ALTER TABLE "food" ADD FOREIGN KEY ("menu_id") REFERENCES "menu" ("id");
ALTER TABLE "food_tag" ADD FOREIGN KEY ("food_id") REFERENCES "food" ("id");

CREATE INDEX ON "stores" ("store_name");

CREATE INDEX ON "tables" ("store_id");
