-- +goose Up
-- +goose StatementBegin
CREATE TABLE competencies (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    numbering VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE competencies;
-- +goose StatementEnd
