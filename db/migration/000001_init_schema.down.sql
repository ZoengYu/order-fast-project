ALTER TABLE IF EXISTS food_tag DROP CONSTRAINT IF EXISTS "menu_food_id";
ALTER TABLE IF EXISTS menu_food DROP CONSTRAINT IF EXISTS "menu_id";
ALTER TABLE IF EXISTS tables DROP CONSTRAINT IF EXISTS "tables_store_id";

DROP TABLE IF EXISTS food_tag;
DROP TABLE IF EXISTS menu_food;
DROP TABLE IF EXISTS menu;
DROP TABLE IF EXISTS tables;
DROP TABLE IF EXISTS stores;