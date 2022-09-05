-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          text NOT NULL UNIQUE,
    username    text PRIMARY KEY,
    password    text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
