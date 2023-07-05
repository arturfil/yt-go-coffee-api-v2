CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS coffees (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" varchar NOT NULL,
    "roast" varchar NOT NULL,
    "region" varchar NOT NULL,
    "image" varchar NOT NULL,
    "price" FLOAT NOT NULL,
    "grind_unit" INT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

