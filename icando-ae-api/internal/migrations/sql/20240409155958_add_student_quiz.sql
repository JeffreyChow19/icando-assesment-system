-- +goose Up
-- +goose StatementBegin
CREATE TYPE QUIZ_STATUS AS ENUM ('NOT_STARTED', 'STARTED', 'SUBMITTED');

CREATE TABLE "student_quizzes" (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    total_score int,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    status QUIZ_STATUS not null default 'NOT_STARTED',
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now(),
    quiz_id uuid NOT NULL REFERENCES quizzes(id)
);

-- NOTES:
-- ASSUME question_competency wont change after the quiz are released,
-- we don't need to create student_answer_competencies table.
-- instead, just create a column in student_answers that will contain array of
-- competency_id and is_passed status

CREATE TABLE "student_answers" (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now(),
    question_id uuid not null references questions(id),
    student_quiz_id uuid not null references student_quizzes(id),
    answer_id int not null,
    is_correct bool default null,
    competencies jsonb -- array of object that have attribute of competency_id and is_passed
);

-- CREATE TABLE "student_answer_competencies" (
--     student_answer_id uuid not null references student_answers(id),
--     question_id uuid not null references questions(id),
--     competency_id uuid not null references competencies(id),
--     is_passed bool not null,
--     primary key (student_answer_id, competency_id)
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE student_answers;
DROP TABLE student_quizzes;
DROP TYPE QUIZ_STATUS;
-- +goose StatementEnd
