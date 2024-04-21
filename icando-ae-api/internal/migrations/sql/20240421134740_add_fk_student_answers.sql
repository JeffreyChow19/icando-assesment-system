-- +goose Up
-- +goose StatementBegin
ALTER TABLE "student_answers"
ADD CONSTRAINT fk_student_quiz FOREIGN KEY (student_quiz_id) references student_quizzes(id);
ALTER TABLE "student_answers"
ADD CONSTRAINT fk_question FOREIGN KEY (question_id) references questions(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "student_answers"
DROP CONSTRAINT fk_student_quiz;
ALTER TABLE "student_answers"
DROP CONSTRAINT fk_question;
-- +goose StatementEnd
