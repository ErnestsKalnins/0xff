CREATE TABLE features
(
    id             BLOB PRIMARY KEY,
    technical_name TEXT NOT NULL,
    display_name   TEXT,
    description    TEXT,
    enabled        TINYINT NOT NULL DEFAULT 0,
    created_at     BIGINT NOT NULL,
    updated_at     BIGINT NOT NULL,
);