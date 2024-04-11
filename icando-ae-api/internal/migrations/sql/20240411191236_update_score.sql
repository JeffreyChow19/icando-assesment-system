-- +goose Up
-- +goose StatementBegin
ALTER TABLE "student_quizzes" DROP COLUMN "total_score";
ALTER TABLE "student_quizzes" ADD COLUMN total_score float;
ALTER TABLE "student_quizzes" ADD COLUMN correct_count int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "student_quizzes" DROP COLUMN "total_score";
ALTER TABLE "student_quizzes" ADD COLUMN total_score int;
ALTER TABLE "student_quizzes" DROP COLUMN correct_count;
-- +goose StatementEnd
