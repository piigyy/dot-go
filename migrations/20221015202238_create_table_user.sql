-- +goose Up
CREATE TABLE "user" (
	"id" SERIAL PRIMARY KEY,
    "fullname" VARCHAR NOT NULL,
	"email" VARCHAR(255) UNIQUE NOT NULL,
    "birth_date" TIMESTAMP NOT NULL
);
CREATE INDEX "user_id_idx" ON "user" USING BTREE (id);
CREATE INDEX "user_email_idx" ON "user" USING BTREE (email);

-- +goose Down
DROP INDEX "user_email_idx";
DROP INDEX "user_id_idx";
DROP TABLE IF EXISTS "user";
