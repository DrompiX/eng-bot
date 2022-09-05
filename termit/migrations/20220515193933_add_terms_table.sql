-- +goose Up
-- +goose StatementBegin
CREATE TABLE terms (
    uid text REFERENCES users(id),
    term text,
    translation text NOT NULL,
    CONSTRAINT terms_pkey PRIMARY KEY (uid, term)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE terms;
-- +goose StatementEnd
