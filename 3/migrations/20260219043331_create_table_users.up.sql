CREATE TABLE IF NOT EXISTS users (
    "id" SERIAL PRIMARY KEY,
    "country" VARCHAR(255) NOT NULL,
    "credit_card_type" VARCHAR(100) NOT NULL,
    "credit_card_number" VARCHAR(100) NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255)
);