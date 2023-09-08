-- +goose Up
-- +goose StatementBegin
CREATE TABLE "profile" (
    "id" bigint NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "phone_number" varchar(15) NOT NULL,
    "gender" varchar NULL,
    "user_id" bigint NOT NULL UNIQUE,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
ALTER TABLE "profile" ADD CONSTRAINT "profile_user_id_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "user" ("id") DEFERRABLE INITIALLY DEFERRED;
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "profile" CASCADE;
COMMIT;
-- +goose StatementEnd
