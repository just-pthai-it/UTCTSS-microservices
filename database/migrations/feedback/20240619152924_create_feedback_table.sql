-- +goose Up
-- +goose StatementBegin
CREATE TABLE feedback
(
    id         SERIAL PRIMARY KEY, -- Use SERIAL for auto-incrementing integer
    data       TEXT      NOT NULL,
    type       SMALLINT  NOT NULL, -- Adjust data type if necessary (e.g., INTEGER)
    is_bug     BOOLEAN  NOT NULL, -- Adjust data type if necessary (e.g., BOOLEAN)
    account_id INTEGER   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) WITH (OIDS = FALSE);
-- OIDS are not typically used in PostgreSQL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feedback;
-- +goose StatementEnd
