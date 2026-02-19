CREATE TABLE IF NOT EXISTS transactions (
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "total_buy" NUMERIC NOT NULL
);