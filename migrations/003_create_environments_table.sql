CREATE TABLE environments
(
    id         BLOB PRIMARY KEY,
    project_id BLOB NOT NULL,
    name       TEXT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects (id)
)