-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes
    ADD COLUMN deadline TIMESTAMPTZ;

ALTER TABLE questions
    DROP COLUMN numbering;

ALTER TABLE questions
    RENAME COLUMN correct_choice_index TO answer_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes
    DROP COLUMN deadline;

ALTER TABLE questions
    ADD COLUMN numbering VARCHAR NOT NULL;

ALTER TABLE questions
    RENAME COLUMN answer_id TO correct_choice_index;
-- +goose StatementEnd
