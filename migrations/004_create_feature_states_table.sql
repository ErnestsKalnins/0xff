CREATE TABLE feature_states
(
    id             BLOB PRIMARY KEY,
    feature_id     BLOB    NOT NULL,
    environment_id BLOB    NOT NULL,
    enabled        TINYINT NOT NULL,
    created_at     BIGINT  NOT NULL,
    updated_at     BIGINT  NOT NULL,
    FOREIGN KEY (feature_id) REFERENCES features (id) ON DELETE CASCADE,
    FOREIGN KEY (environment_id) REFERENCES environments (id) ON DELETE CASCADE
)