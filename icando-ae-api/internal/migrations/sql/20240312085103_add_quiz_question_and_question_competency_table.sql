-- +goose Up
-- +goose StatementBegin
CREATE TABLE quizzes (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR NOT NULL,
    subject VARCHAR NOT NULL,
    passing_grade DECIMAL NOT NULL,
    parent_quiz uuid REFERENCES quizzes(id),
    created_by uuid NOT NULL,
    updated_by uuid,
    published_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE questions (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    numbering VARCHAR NOT NULL,
    choices jsonb,
    correct_choice_index INT,
    quiz_id uuid NOT NULL REFERENCES quizzes(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE question_competency (
    question_id uuid NOT NULL REFERENCES questions(id),
    competency_id uuid NOT NULL REFERENCES competencies(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE question_competency;
DROP TABLE questions;
DROP TABLE quizzes;
-- +goose StatementEnd
