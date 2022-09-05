-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id text PRIMARY KEY,
    uid text REFERENCES users(id),
    term text NOT NULL,
    expected text NOT NULL,
    success boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
