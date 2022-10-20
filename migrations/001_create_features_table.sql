CREATE TABLE features
(
    id             BLOB PRIMARY KEY,
    project_id     BLOB,
    technical_name TEXT   NOT NULL,
    display_name   TEXT,
    description    TEXT,
    created_at     BIGINT NOT NULL,
    updated_at     BIGINT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects (id)
);