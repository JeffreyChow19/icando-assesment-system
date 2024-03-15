-- +goose Up
-- +goose StatementBegin
ALTER TABLE students
  ALTER COLUMN class_id DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE students
  ALTER COLUMN class_id SET NOT NULL;
-- +goose StatementEnd
