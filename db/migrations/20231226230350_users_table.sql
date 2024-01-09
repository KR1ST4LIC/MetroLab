-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA users;

CREATE TABLE users.user
(
    "id"       serial,
    "tg_id"    bigint UNIQUE,
    "pay"      bool,
    "pubg_id"  bigint,
    "soprovod" text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users.user;

DROP SCHEMA users CASCADE;
-- +goose StatementEnd
