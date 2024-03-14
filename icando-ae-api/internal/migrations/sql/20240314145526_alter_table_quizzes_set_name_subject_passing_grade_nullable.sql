-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes
    ALTER COLUMN name DROP NOT NULL,
    ALTER COLUMN subject DROP NOT NULL,
    ALTER COLUMN passing_grade DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes
    ALTER COLUMN name SET NOT NULL,
    ALTER COLUMN subject SET NOT NULL,
    ALTER COLUMN passing_grade SET NOT NULL;
-- +goose StatementEnd
