// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package questionnaire

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// Questionnaire represents the entire questionnaire.
type Questionnaire struct {
	Questions []*Question `json:"questions"`
}

// Question represents a single question with its possible answers.
type Question struct {
	ULID    string    `json:"ulid"`
	Label   string    `json:"label"`
	Text    string    `json:"text"`
	Answers []*Answer `json:"answers"`
}

// Answer represents an answer choice, with a possible next question.
type Answer struct {
	ULID             string  `json:"ulid"`
	Text             string  `json:"text"`
	NextQuestionULID *string `json:"next_question_ulid"` // Nullable
}

// getQuery is the SQL query to retrieve the full questionnaire.
const getQuery = `
	SELECT q.ulid as q_ulid, 
		q.label as q_label, 
		q.text as q_text, 
		a.ulid as a_ulid, 
		a.text as a_text, 
		a.next_question_ulid as a_next_question_ulid
	FROM questions q
	LEFT JOIN answers a ON q.ulid = a.question_ulid
	ORDER BY q.id, a.id;
`

// Get retrieves the full questionnaire from the database.
func Get(ctx context.Context, db *sql.DB) (*Questionnaire, error) {
	rows, err := db.QueryContext(ctx, getQuery)
	if err != nil {
		return nil, errors.Wrap(err, "getting questionnaire")
	}
	defer rows.Close()
	questionMap := make(map[string]*Question)
	var questions []*Question
	for rows.Next() {
		var (
			qULID, qLabel, qText string
			aULID, aText         sql.NullString
			nextQULID            sql.NullString
		)
		if err := rows.Scan(&qULID, &qLabel, &qText, &aULID, &aText, &nextQULID); err != nil {
			return nil, errors.Wrap(err, "scanning row")
		}
		q, exists := questionMap[qULID]
		if !exists {
			q = &Question{
				ULID:    qULID,
				Label:   qLabel,
				Text:    qText,
				Answers: []*Answer{},
			}
			questionMap[qULID] = q
			questions = append(questions, q)
		}
		if aULID.Valid {
			answer := &Answer{
				ULID: aULID.String,
				Text: aText.String,
			}
			if nextQULID.Valid {
				answer.NextQuestionULID = &nextQULID.String
			}
			q.Answers = append(q.Answers, answer)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "iterating over rows")
	}
	return &Questionnaire{Questions: questions}, nil
}
