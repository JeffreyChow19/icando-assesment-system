-- +goose Up
-- +goose StatementBegin
CREATE TABLE "quiz_classes" (
    quiz_id uuid references quizzes(id),
    class_id uuid references classes(id),
    primary key (quiz_id, class_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "quiz_classes";
-- +goose StatementEnd
