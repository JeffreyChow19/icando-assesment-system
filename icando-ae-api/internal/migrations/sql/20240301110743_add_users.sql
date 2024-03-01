-- +goose Up
-- +goose StatementBegin
CREATE TABLE institutions (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR NOT NULL,
    nis VARCHAR NOT NULL,
    slug VARCHAR NOT NULL UNIQUE
);

CREATE TABLE teachers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    institution_id uuid NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES institutions(id)
);

CREATE TABLE learning_designers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    institution_id uuid NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES institutions(id)
);

CREATE TABLE classes (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR NOT NULL,
    grade VARCHAR NOT NULL,
    institution_id uuid NOT NULL,
    teacher_id uuid NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES institutions(id),
    FOREIGN KEY (teacher_id) REFERENCES teachers(id)
);
CREATE TABLE students (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    nisn VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    institution_id uuid NOT NULL,
    class_id uuid NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES institutions(id),
    FOREIGN KEY (class_id) REFERENCES classes(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE students;
DROP TABLE classes;
DROP TABLE learning_designers;
DROP TABLE teachers;
DROP TABLE institutions;
-- +goose StatementEnd
