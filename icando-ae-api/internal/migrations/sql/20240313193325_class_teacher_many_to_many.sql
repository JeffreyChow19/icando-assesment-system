-- +goose Up
-- +goose StatementBegin
CREATE TABLE "class_teacher" (
    class_id uuid not null references classes(id) on delete cascade on update cascade,
    teacher_id uuid not null references teachers(id) on delete cascade on update cascade,
    primary key (class_id, teacher_id)
);

INSERT INTO "class_teacher"(class_id, teacher_id)
SELECT id as class_id, teacher_id
FROM classes;

ALTER TABLE "classes" DROP COLUMN teacher_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "classes"
ADD COLUMN teacher_id uuid,
ADD CONSTRAINT "class_teacher_fk" FOREIGN KEY (teacher_id) REFERENCES teachers(id);
-- +goose StatementEnd
