-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_book" (
    "user_id" SERIAL NOT NULL,
    "book_id" SERIAL NOT NULL,
    "borrowed_at" TIMESTAMP NOT NULL,
    "duration" INT NOT NULL,
    "duration_unit" VARCHAR(1) NOT NULL,
    "returned_at" TIMESTAMP

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_book"
-- +goose StatementEnd
