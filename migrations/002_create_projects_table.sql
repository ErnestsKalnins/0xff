CREATE TABLE projects
(
    id         BLOB PRIMARY KEY,
    name       TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);