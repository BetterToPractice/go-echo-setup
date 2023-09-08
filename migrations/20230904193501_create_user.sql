-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    "id" bigint NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "email" varchar(255) NOT NULL,
    "password" varchar(100) NOT NULL,
    "username" varchar(100) NOT NULL,
    "created_at" TIMESTAMP NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL
);
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user" CASCADE;
COMMIT;
-- +goose StatementEnd