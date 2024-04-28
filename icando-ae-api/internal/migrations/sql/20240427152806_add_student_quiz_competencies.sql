-- +goose Up
-- +goose StatementBegin
CREATE TABLE "student_quiz_competencies" (
    student_quiz_id uuid references student_quizzes(id),
    student_id uuid references students(id),
    competency_id uuid references competencies(id),
    correct_count int check(correct_count >= 0) not null,
    total_count int check(correct_count >= 0) not null,
    primary key (student_quiz_id, student_id, competency_id)
);

CREATE INDEX sqc_stident_id_idx ON student_quiz_competencies(student_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index sqc_stident_id_idx;
drop table student_quiz_competencies;
-- +goose StatementEnd
