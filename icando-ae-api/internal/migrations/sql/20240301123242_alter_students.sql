-- +goose Up
-- +goose StatementBegin
ALTER TABLE institutions
ADD CONSTRAINT "institutions_slug_unique" UNIQUE (slug);

ALTER TABLE teachers
ADD CONSTRAINT "teachers_email_unique" UNIQUE (email);

ALTER TABLE learning_designers
ADD CONSTRAINT "learning_designers_email_unique" UNIQUE (email);

ALTER TABLE Students
ADD CONSTRAINT "students_nisn_unique" UNIQUE (nisn),
ADD CONSTRAINT "students_email_unique" UNIQUE (email),
DROP COLUMN password;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE institutions
DROP CONSTRAINT "institutions_slug_unique";

ALTER TABLE teachers
DROP CONSTRAINT "teachers_email_unique";

ALTER TABLE learning_designers
DROP CONSTRAINT "learning_designers_email_unique";

ALTER TABLE Students
DROP CONSTRAINT "students_nisn_unique",
DROP CONSTRAINT "students_email_unique",
ADD COLUMN password VARCHAR(255); 
-- +goose StatementEnd
