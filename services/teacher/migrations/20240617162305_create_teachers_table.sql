-- +goose Up
-- +goose StatementBegin
CREATE TABLE teachers
(
    id                        SERIAL PRIMARY KEY,                  -- Use SERIAL for auto-incrementing integer
    code                      VARCHAR(50)  NOT NULL UNIQUE,
    name                      VARCHAR(100) NOT NULL,
    is_female                 BOOLEAN      NOT NULL DEFAULT FALSE, -- Use BOOLEAN for true/false values
    birth                     DATE         NOT NULL,
    university_teacher_degree VARCHAR(200) NOT NULL DEFAULT '',
    is_head_of_department     BOOLEAN      NOT NULL DEFAULT FALSE,
    is_head_of_faculty        BOOLEAN      NOT NULL DEFAULT FALSE,
    is_active                 BOOLEAN      NOT NULL DEFAULT FALSE,
    department_id             VARCHAR(50)  NOT NULL,
    deleted_at                TIMESTAMP             DEFAULT NULL
) WITH (OIDS = FALSE);
-- OIDS are not typically used in PostgreSQL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teachers;
-- +goose StatementEnd
