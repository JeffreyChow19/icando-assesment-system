-- +goose Up
-- +goose StatementBegin
DROP TABLE learning_designers;
CREATE TYPE TEACHER_ROLE AS ENUM ('LEARNING_DESIGNER', 'REGULAR');
ALTER TABLE teachers
    ADD COLUMN role TEACHER_ROLE NOT NULL DEFAULT 'REGULAR',
    ADD COLUMN created_at TIMESTAMPTZ,
    ADD COLUMN updated_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE teachers
    DROP COLUMN updated_at,
    DROP COLUMN created_at,
    DROP COLUMN role;
CREATE TABLE learning_designers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    institution_id uuid NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES institutions(id)
);
-- +goose StatementEnd
