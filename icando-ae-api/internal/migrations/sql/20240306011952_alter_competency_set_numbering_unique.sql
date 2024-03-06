-- +goose Up
-- +goose StatementBegin
ALTER TABLE competencies ADD CONSTRAINT numbering_unique UNIQUE (numbering);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE competencies
DROP CONSTRAINT numbering_unique;
-- +goose StatementEnd
