CREATE TABLE answers_flow (
    id SERIAL PRIMARY KEY,
    answer_ulid TEXT NOT NULL REFERENCES answers(ulid) ON DELETE CASCADE,
    previous_answer_ulid TEXT REFERENCES answers(ulid) ON DELETE CASCADE,
    next_question_ulid TEXT REFERENCES questions(ulid) ON DELETE CASCADE,
    CONSTRAINT answers_flow_unique UNIQUE (answer_ulid, previous_answer_ulid)
);
