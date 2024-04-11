-- +goose Up
-- +goose StatementBegin
ALTER TABLE "student_quizzes" ADD COLUMN student_id uuid references students(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "student_quizzes" DROP COLUMN student_id;
-- +goose StatementEnd
