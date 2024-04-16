-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes RENAME COLUMN deadline TO end_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes RENAME COLUMN end_at TO deadline;
-- +goose StatementEnd
