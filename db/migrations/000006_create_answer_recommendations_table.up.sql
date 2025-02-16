CREATE TABLE IF NOT EXISTS answer_recommendations (
    answer_ulid TEXT NOT NULL,
    recommendation_ulid TEXT NOT NULL,
    PRIMARY KEY (answer_ulid, recommendation_ulid),
    FOREIGN KEY (answer_ulid) REFERENCES answers(ulid),
    FOREIGN KEY (recommendation_ulid) REFERENCES recommendations(ulid)
);