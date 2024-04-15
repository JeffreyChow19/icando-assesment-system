-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes
    ALTER COLUMN subject TYPE text[] USING ARRAY [subject];
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes
    ALTER COLUMN subject TYPE varchar;
-- +goose StatementEnd
