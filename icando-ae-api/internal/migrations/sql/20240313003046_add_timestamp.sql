-- +goose Up
-- +goose StatementBegin
ALTER TABLE institutions
    ADD COLUMN created_at TIMESTAMPTZ,
    ADD COLUMN updated_at TIMESTAMPTZ;

ALTER TABLE classes
    ADD COLUMN created_at TIMESTAMPTZ,
    ADD COLUMN updated_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE institutions
    DROP COLUMN created_at,
    DROP COLUMN updated_at;

ALTER TABLE classes
    DROP COLUMN created_at,
    DROP COLUMN updated_at;
-- +goose StatementEnd
