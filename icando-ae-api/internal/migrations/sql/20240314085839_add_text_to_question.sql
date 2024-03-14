-- +goose Up
-- +goose StatementBegin
ALTER TABLE questions
    ADD COLUMN text TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE questions
    DROP COLUMN text;
-- +goose StatementEnd
