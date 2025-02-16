// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package answers

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/questionnaire-rest-api/ptr"
)

func TestRecommendationsFromAnswers(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		mockClosure    func() *sql.DB
		mockScanRows   func(rows *sql.Rows, dest ...interface{}) error
		expectedOutput *RecommendationsResponse
		expectedError  error
	}{
		{
			name:  "happy path",
			input: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				expectedQuery := fmt.Sprintf(cteQueryTemplate, "'01JKZMRKJHW7MF4HKN0DAY0PAR'")
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"exclusion_reason", "recommendation", "next_question_ulid"}).
							AddRow(nil, "Recommendation 1", nil),
					)
				return db
			},
			expectedOutput: &RecommendationsResponse{
				Recommendations: []string{"Recommendation 1"},
			},
		},
		{
			name:  "exclusion reason",
			input: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				expectedQuery := fmt.Sprintf(cteQueryTemplate, "'01JKZMRKJHW7MF4HKN0DAY0PAR'")
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"exclusion_reason", "recommendation", "next_question_ulid"}).
							AddRow("exclusion reason", nil, nil),
					)
				return db
			},
			expectedOutput: &RecommendationsResponse{
				ExclusionReason: ptr.P("exclusion reason"),
			},
		},
		{
			name:  "next question",
			input: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				expectedQuery := fmt.Sprintf(cteQueryTemplate, "'01JKZMRKJHW7MF4HKN0DAY0PAR'")
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"exclusion_reason", "recommendation", "next_question_ulid"}).
							AddRow(nil, nil, "01JKZMRKJHW7MF4HKN0DAY0PAR"),
					)
				return db
			},
			expectedOutput: &RecommendationsResponse{
				NextQuestionULID: ptr.P("01JKZMRKJHW7MF4HKN0DAY0PAR"),
			},
		},
		{
			name: "no answers provided",
			mockClosure: func() *sql.DB {
				db, _, _ := sqlmock.New()
				return db
			},
			expectedError: errors.New("no answers provided"),
		},
		{
			name:  "error querying",
			input: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				expectedQuery := fmt.Sprintf(cteQueryTemplate, "'01JKZMRKJHW7MF4HKN0DAY0PAR'")
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnError(errors.New("error querying"))
				return db
			},
			expectedError: errors.New("executing recommendations query: error querying"),
		},
		{
			name:  "error scanning",
			input: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				expectedQuery := fmt.Sprintf(cteQueryTemplate, "'01JKZMRKJHW7MF4HKN0DAY0PAR'")
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"exclusion_reason", "recommendation", "next_question_ulid"}).
							AddRow(nil, "Recommendation 1", nil),
					)

				return db
			},
			mockScanRows: func(rows *sql.Rows, dest ...interface{}) error {
				return errors.New("error scanning")
			},
			expectedError: errors.New("scanning recommendations row: error scanning"),
		},
		{
			name:  "error iterating over rows",
			input: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				expectedQuery := fmt.Sprintf(cteQueryTemplate, "'01JKZMRKJHW7MF4HKN0DAY0PAR'")
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"exclusion_reason", "recommendation", "next_question_ulid"}).
							AddRow(nil, "Recommendation 1", nil).
							RowError(0, errors.New("error iterating")),
					)

				return db
			},
			expectedError: errors.New("iterating over recommendations rows: error iterating"),
		},
	}
	originalScanRows := scanRows
	for _, tc := range testCases {
		defer func() { scanRows = originalScanRows }()
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockScanRows != nil {
				scanRows = tc.mockScanRows
			}
			db := tc.mockClosure()
			defer db.Close()
			output, err := RecommendationsFromAnswers(context.TODO(), db, tc.input)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no error, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v", got nil`, tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

func TestGetAnswerFlow(t *testing.T) {
	testCases := []struct {
		name           string
		mockClosure    func() *sql.DB
		mockScanRows   func(rows *sql.Rows, dest ...interface{}) error
		expectedOutput map[string]AnswerFlow
		expectedError  error
	}{
		{
			name: "happy path",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectQuery(regexp.QuoteMeta(getAnswerFlowQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"answer_ulid", "previous_answers", "next_question_ulid"}).
							AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", pq.StringArray{"01JKZMRKJHW7MF4HKN0DAY0PAR"}, nil),
					)
				return db
			},
			expectedOutput: map[string]AnswerFlow{
				"01JKZMRKJHW7MF4HKN0DAY0PAR": {
					PreviousAnswers: []string{"01JKZMRKJHW7MF4HKN0DAY0PAR"},
				},
			},
		},
		{
			name: "handles multiple previous answers",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectQuery(regexp.QuoteMeta(getAnswerFlowQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"answer_ulid", "previous_answers", "next_question_ulid"}).
							AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", pq.StringArray{"01JKZMRQ3GS5731G5MJ9Y79DMH"}, nil).
							AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", pq.StringArray{"01JKZMRPGPWW4GF1B55BBS3R9Z"}, nil),
					)
				return db
			},
			expectedOutput: map[string]AnswerFlow{
				"01JKZMRKJHW7MF4HKN0DAY0PAR": {
					PreviousAnswers: []string{
						"01JKZMRQ3GS5731G5MJ9Y79DMH", // First previous answer
						"01JKZMRPGPWW4GF1B55BBS3R9Z", // Second previous answer
					},
				},
			},
		},
		{
			name: "error querying",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectQuery(regexp.QuoteMeta(getAnswerFlowQuery)).
					WillReturnError(sql.ErrConnDone)
				return db
			},
			expectedError: errors.New("querying answers flow: sql: connection is already closed"),
		},
		{
			name: "error scanning",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"answer_ulid", "previous_answers", "next_question_ulid"}).
					AddRow(nil, nil, nil)
				mock.ExpectQuery(regexp.QuoteMeta(getAnswerFlowQuery)).
					WillReturnRows(rows)
				return db
			},
			mockScanRows: func(rows *sql.Rows, dest ...interface{}) error {
				return errors.New("scanning row")
			},
			expectedError: errors.New("scanning answer flow row: scanning row"),
		},
		{
			name: "error iterating over rows",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"answer_ulid", "previous_answers", "next_question_ulid"}).
					AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", pq.StringArray{"01JKZMRKJHW7MF4HKN0DAY0PAR"}, nil).
					RowError(0, errors.New("error iterating"))
				mock.ExpectQuery(regexp.QuoteMeta(getAnswerFlowQuery)).
					WillReturnRows(rows)
				return db
			},
			expectedError: errors.New("iterating over answers flow rows: error iterating"),
		},
	}
	originalScanRows := scanRows
	for _, tc := range testCases {
		defer func() { scanRows = originalScanRows }()
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockScanRows != nil {
				scanRows = tc.mockScanRows
			}
			db := tc.mockClosure()
			defer db.Close()
			output, err := GetAnswerFlow(context.TODO(), db)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no error, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v", got nil`, tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}
