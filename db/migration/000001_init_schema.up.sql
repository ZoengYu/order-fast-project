CREATE TABLE IF NOT EXISTS stores (
  "id" bigserial PRIMARY KEY,
  "store_name" varchar NOT NULL,
  "store_address" varchar NOT NULL,
  "store_phone" varchar NOT NULL,
  "store_owner" varchar NOT NULL,
  "store_manager" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "stores" ("store_name");
