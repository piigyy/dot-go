-- +goose Up
-- +goose StatementBegin
CREATE TABLE "book" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "isbn" VARCHAR(13) NOT NULL,
    "author" VARCHAR NOT NULL,
    "published_at" TIMESTAMP,
    "stock" INT NOT NULL
);
CREATE INDEX "book_id_idx" ON "book" USING BTREE (id);
CREATE INDEX "book_title_idx" ON "book" USING BTREE (title);
CREATE INDEX "book_isbn_idx" ON "book" USING BTREE (isbn);
CREATE INDEX "book_author_idx" ON "book" USING BTREE (author);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "book_id_idx"
DROP INDEX "book_title_idx"
DROP INDEX "book_isbn_idx"
DROP INDEX "book_author_idx"
DROP TABLE IF EXISTS "book"
-- +goose StatementEnd
