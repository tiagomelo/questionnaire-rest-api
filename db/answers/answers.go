// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package answers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// For ease of unit testing.
var scanRows = func(rows *sql.Rows, dest ...interface{}) error {
	return rows.Scan(dest...)
}

const (
	// cteQueryTemplate is a CTE query that retrieves the next question, recommendations, or exclusions.
	cteQueryTemplate = `
	WITH user_answers AS (
		SELECT unnest(ARRAY[%s]::TEXT[]) AS answer_ulid
	),
	exclusions AS (
		-- Find if any of the given answers trigger an exclusion
		SELECT reason AS exclusion_reason
		FROM exclusions
		WHERE answer_ulid IN (SELECT answer_ulid FROM user_answers)
		LIMIT 1 -- Stop at the first exclusion
	),
	latest_answer AS (
		-- Get the last answered question (if no exclusion exists)
		SELECT a.next_question_ulid
		FROM answers a
		WHERE a.ulid IN (SELECT answer_ulid FROM user_answers)
		ORDER BY a.id DESC
		LIMIT 1
	),
	next_question AS (
		-- Determine the next expected question
		SELECT next_question_ulid
		FROM latest_answer
		WHERE next_question_ulid IS NOT NULL
		AND next_question_ulid NOT IN (SELECT answer_ulid FROM user_answers)
		AND NOT EXISTS (SELECT 1 FROM exclusions) -- Skip if exclusion exists
	),
	recommendations AS (
		-- Retrieve product recommendations if no exclusions exist
		SELECT p.name AS recommendation
		FROM answer_recommendations ar
		JOIN recommendations r ON ar.recommendation_ulid = r.ulid
		JOIN products p ON r.product_ulid = p.ulid
		WHERE ar.answer_ulid IN (SELECT answer_ulid FROM user_answers)
		AND NOT EXISTS (SELECT 1 FROM exclusions) -- Skip if exclusion exists
		AND NOT EXISTS (SELECT 1 FROM next_question) -- Skip if next question exists
	)

	-- Prioritize Exclusions
	SELECT exclusion_reason, NULL::TEXT AS recommendation, NULL::TEXT AS next_question_ulid
	FROM exclusions

	UNION ALL

	-- Next Question if No Exclusion
	SELECT NULL::TEXT AS exclusion_reason, NULL::TEXT AS recommendation, next_question_ulid
	FROM next_question
	WHERE NOT EXISTS (SELECT 1 FROM exclusions)

	UNION ALL

	-- Recommendations if No Exclusion & No Next Question
	SELECT NULL::TEXT AS exclusion_reason, recommendation, NULL::TEXT AS next_question_ulid
	FROM recommendations
	WHERE NOT EXISTS (SELECT 1 FROM exclusions)
	AND NOT EXISTS (SELECT 1 FROM next_question)
`

	// getAnswerFlowQuery retrieves the answer flow for each answer.
	getAnswerFlowQuery = `
		SELECT 
			af.answer_ulid, 
			COALESCE(array_agg(af.previous_answer_ulid) FILTER (WHERE af.previous_answer_ulid IS NOT NULL), ARRAY[]::TEXT[]) AS previous_answers, 
			af.next_question_ulid
		FROM answers_flow af
		GROUP BY af.answer_ulid, af.next_question_ulid
	`
)

// RecommendationsResponse represents the result of processing user answers.
type RecommendationsResponse struct {
	ExclusionReason  *string  `json:"exclusion_reason,omitempty"`
	Recommendations  []string `json:"recommendations,omitempty"`
	NextQuestionULID *string  `json:"next_question_ulid,omitempty"`
}

// RecommendationsFromAnswers retrieves the next question, recommendations, or exclusions.
func RecommendationsFromAnswers(ctx context.Context, db *sql.DB, answerULIDs []string) (*RecommendationsResponse, error) {
	if len(answerULIDs) == 0 {
		return nil, errors.New("no answers provided")
	}
	cteQuery := fmt.Sprintf(cteQueryTemplate, "'"+strings.Join(answerULIDs, "','")+"'")
	rows, err := db.QueryContext(ctx, cteQuery)
	if err != nil {
		return nil, errors.Wrap(err, "executing recommendations query")
	}
	defer rows.Close()
	var response RecommendationsResponse
	for rows.Next() {
		var exclusionReason, recommendation sql.NullString
		var nextQuestionULID sql.NullString
		if err := scanRows(rows, &exclusionReason, &recommendation, &nextQuestionULID); err != nil {
			return nil, errors.Wrap(err, "scanning recommendations row")
		}
		if exclusionReason.Valid {
			response.ExclusionReason = &exclusionReason.String
		}
		if recommendation.Valid {
			response.Recommendations = append(response.Recommendations, recommendation.String)
		}
		if nextQuestionULID.Valid {
			response.NextQuestionULID = &nextQuestionULID.String
		}
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "iterating over recommendations rows")
	}
	return &response, nil
}

// AnswerFlow represents the flow of answers.
type AnswerFlow struct {
	PreviousAnswers []string
	Next            *string
}

// GetAnswerFlow retrieves the answer flow for each answer.
func GetAnswerFlow(ctx context.Context, db *sql.DB) (map[string]AnswerFlow, error) {
	rows, err := db.QueryContext(ctx, getAnswerFlowQuery)
	if err != nil {
		return nil, errors.Wrap(err, "querying answers flow")
	}
	defer rows.Close()
	answerFlow := make(map[string]AnswerFlow)
	for rows.Next() {
		var answerULID string
		var previousAnswers pq.StringArray
		var nextQuestionULID *string
		if err := scanRows(rows, &answerULID, (*pq.StringArray)(&previousAnswers), &nextQuestionULID); err != nil {
			return nil, errors.Wrap(err, "scanning answer flow row")
		}
		// if answer exists, append previous answers instead of overwriting.
		if flow, exists := answerFlow[answerULID]; exists {
			flow.PreviousAnswers = append(flow.PreviousAnswers, previousAnswers...)
			answerFlow[answerULID] = flow
		} else {
			answerFlow[answerULID] = AnswerFlow{
				PreviousAnswers: previousAnswers,
				Next:            nextQuestionULID,
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "iterating over answers flow rows")
	}
	return answerFlow, nil
}
