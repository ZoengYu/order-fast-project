CREATE TABLE users (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS stores (
  "id" bigserial PRIMARY KEY,
  "owner" varchar(60) NOT NULL,
  "name" varchar(60) NOT NULL,
  "address" varchar(120) NOT NULL,
  "phone" varchar(10) UNIQUE NOT NULL,
  "manager" varchar NOT NULL,
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

ALTER TABLE "stores" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
ALTER TABLE "tables" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");
ALTER TABLE "item" ADD FOREIGN KEY ("menu_id") REFERENCES "menu" ("id");
ALTER TABLE "item_tag" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("id");

CREATE INDEX ON "stores" ("id", "owner");

CREATE INDEX ON "tables" ("store_id");
