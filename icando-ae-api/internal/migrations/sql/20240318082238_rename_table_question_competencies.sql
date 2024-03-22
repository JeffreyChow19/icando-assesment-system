-- +goose Up
-- +goose StatementBegin
ALTER TABLE question_competency RENAME TO question_competencies;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE question_competencies RENAME TO question_competency;
-- +goose StatementEnd
