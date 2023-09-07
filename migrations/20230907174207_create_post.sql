-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "post" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    user_id bigint NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

ALTER TABLE "post" ADD CONSTRAINT "post_user_id_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "post_user_id" ON "post" ("user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "post" CASCADE;
-- +goose StatementEnd
