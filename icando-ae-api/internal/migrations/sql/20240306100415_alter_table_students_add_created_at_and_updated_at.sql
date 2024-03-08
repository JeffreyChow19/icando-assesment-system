-- +goose Up
-- +goose StatementBegin
ALTER TABLE students
    ADD COLUMN created_at TIMESTAMPTZ,
    ADD COLUMN updated_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE students
    DROP COLUMN updated_at,
    DROP COLUMN created_at;
-- +goose StatementEnd
