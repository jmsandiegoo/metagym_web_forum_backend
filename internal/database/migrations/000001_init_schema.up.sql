-- ENUM Type for user table
CREATE TYPE EXPERIENCE_ENUM AS ENUM ('beginner', 'intermediate', 'expert');
-- user table
CREATE TABLE "user" (
    "user_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "username" VARCHAR(30) NOT NULL UNIQUE,
    "email" VARCHAR NOT NULL,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "rep" INT NOT NULL DEFAULT 0,
    "pfp_url" VARCHAR,
    "bio" VARCHAR(450),
    "experience" EXPERIENCE_ENUM DEFAULT 'beginner',
    "country" VARCHAR,
    "height" NUMERIC,
    "weight" NUMERIC,
    "age" INT,
    "isVerified" BOOLEAN DEFAULT FALSE,
    "create_timestamp" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "update_timestamp" TIMESTAMPTZ NOT NULL,
    -- OTHER CONSTRAINTS
    PRIMARY KEY ("user_id"),
    CONSTRAINT chk_user CHECK (
        (NOT "isVerified")
        OR (
            "country" IS NOT NULL
            AND "height" IS NOT NULL
            AND "weight" IS NOT NULL
            AND "age" IS NOT NULL
        )
    )
);
-- INDEXES for user Table for more efficient queries
CREATE UNIQUE INDEX "user_username_key" ON "user" ("username");
-- interest table
CREATE TABLE "interest" (
    "interest_id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR UNIQUE NOT NULL,
    "create_timestamp" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "update_timestamp" TIMESTAMPTZ NOT NULL,
    PRIMARY KEY ("interest_id")
);
-- interest table
CREATE TABLE "userinterest" (
    "user_id" UUID NOT NULL,
    "interest_id" UUID NOT NULL,
    PRIMARY KEY ("user_id", "interest_id"),
    CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES "user" ("user_id"),
    CONSTRAINT fk_interest FOREIGN KEY ("interest_id") REFERENCES "interest" ("interest_id")
);