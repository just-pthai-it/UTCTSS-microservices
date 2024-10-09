-- +goose Up
-- +goose StatementBegin
CREATE TABLE mail_templates
(
    id         SERIAL PRIMARY KEY, -- Use SERIAL for auto-incrementing integer
    code       VARCHAR(50)  NOT NULL UNIQUE,
    subject    VARCHAR(100) NOT NULL UNIQUE,
    content    text         NOT NULL DEFAULT '',
    layout     VARCHAR(50)  NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    deleted_at TIMESTAMP             DEFAULT NULL
) WITH (OIDS = FALSE);
-- OIDS are not typically used in PostgreSQL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS mail_templates;
-- +goose StatementEnd
