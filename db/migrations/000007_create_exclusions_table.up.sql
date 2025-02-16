CREATE TABLE IF NOT EXISTS exclusions (
    answer_ulid TEXT PRIMARY KEY,
    reason TEXT NOT NULL,
    FOREIGN KEY (answer_ulid) REFERENCES answers(ulid) ON DELETE CASCADE
);
