-- +goose Up
-- +goose StatementBegin
ALTER TABLE "student_answers"
    DROP CONSTRAINT student_answers_pkey;
ALTER TABLE "student_answers"
    DROP COLUMN id;
ALTER TABLE "student_answers"
    ADD PRIMARY KEY (student_quiz_id, question_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "student_answers"
    DROP CONSTRAINT student_answers_pkey;
ALTER TABLE "student_answers"
    ADD COLUMN id uuid default gen_random_uuid();
ALTER TABLE "student_answers"
    ADD PRIMARY KEY (id);
-- +goose StatementEnd
