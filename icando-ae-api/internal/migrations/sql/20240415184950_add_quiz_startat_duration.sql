-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes ADD COLUMN duration INTEGER;
ALTER TABLE quizzes ADD COLUMN start_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes DROP COLUMN duration;
ALTER TABLE quizzes DROP COLUMN start_at;
-- +goose StatementEnd

