-- +goose Up
-- +goose StatementBegin
ALTER TABLE questions
ADD COLUMN "order" INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE questions
DROP COLUMN "order";
-- +goose StatementEnd