CREATE TABLE IF NOT EXISTS answers (
    id SERIAL PRIMARY KEY,
    ulid TEXT UNIQUE NOT NULL CHECK (ulid ~ '^[0-9A-HJKMNP-TV-Z]{26}$'),
    question_ulid TEXT NOT NULL,
    text TEXT NOT NULL,
    next_question_ulid TEXT,
    previous_question_ulid TEXT,
    CONSTRAINT fk_question FOREIGN KEY (question_ulid) REFERENCES questions(ulid),
    CONSTRAINT fk_next_question FOREIGN KEY (next_question_ulid) REFERENCES questions(ulid),
    CONSTRAINT fk_previous_question FOREIGN KEY (previous_question_ulid) REFERENCES questions(ulid)
);
