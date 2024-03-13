-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes
    ADD CONSTRAINT fk_quizzes_created_by FOREIGN KEY (created_by) REFERENCES teachers(id),
    ADD CONSTRAINT fk_quizzes_updated_by FOREIGN KEY (updated_by) REFERENCES teachers(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes
    DROP CONSTRAINT fk_quizzes_created_by,
    DROP CONSTRAINT fk_quizzes_updated_by;
-- +goose StatementEnd
