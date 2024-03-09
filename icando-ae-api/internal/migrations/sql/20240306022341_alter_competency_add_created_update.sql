-- +goose Up
-- +goose StatementBegin
ALTER TABLE competencies
    ADD COLUMN created_at TIMESTAMPTZ,
    ADD COLUMN updated_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE competencies
DROP COLUMN created_at,
DROP COLUMN updated_at;
-- +goose StatementEnd
