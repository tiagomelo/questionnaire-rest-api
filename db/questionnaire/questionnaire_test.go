// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package questionnaire

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/questionnaire-rest-api/ptr"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name           string
		mockClosure    func() *sql.DB
		expectedOutput *Questionnaire
		expectedError  error
	}{
		{
			name: "happy path",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectQuery(regexp.QuoteMeta(getQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"q_ulid", "q_label", "q_text", "a_ulid", "a_text", "a_next_question_ulid"}).
							AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", "1", "Do you have difficulty getting or maintaining an erection?", "01JKZMRNXVQ4QSJ548FBSV1GJX", "Yes", "01JKZMRKVYH2SQXS5R9EXQQYWR").
							AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", "1", "Do you have difficulty getting or maintaining an erection?", "01JKZMRP780464QR5GRKKJFE38", "No", nil),
					)
				return db
			},
			expectedOutput: &Questionnaire{
				Questions: []*Question{
					{
						ULID:  "01JKZMRKJHW7MF4HKN0DAY0PAR",
						Label: "1",
						Text:  "Do you have difficulty getting or maintaining an erection?",
						Answers: []*Answer{
							{
								ULID:             "01JKZMRNXVQ4QSJ548FBSV1GJX",
								Text:             "Yes",
								NextQuestionULID: ptr.P("01JKZMRKVYH2SQXS5R9EXQQYWR"),
							},
							{
								ULID:             "01JKZMRP780464QR5GRKKJFE38",
								Text:             "No",
								NextQuestionULID: nil, // No next question
							},
						},
					},
				},
			},
		},
		{
			name: "error querying",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectQuery(regexp.QuoteMeta(getQuery)).
					WillReturnError(sql.ErrConnDone)
				return db
			},
			expectedError: errors.New("getting questionnaire: sql: connection is already closed"),
		},
		{
			name: "error iterating over rows",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectQuery(regexp.QuoteMeta(getQuery)).
					WillReturnRows(
						sqlmock.NewRows([]string{"q_ulid", "q_label", "q_text", "a_ulid", "a_text", "a_next_question_ulid"}).
							AddRow("01JKZMRKJHW7MF4HKN0DAY0PAR", "1", "Do you have difficulty getting or maintaining an erection?", "01JKZMRNXVQ4QSJ548FBSV1GJX", "Yes", "01JKZMRKVYH2SQXS5R9EXQQYWR").
							RowError(0, sql.ErrConnDone),
					)
				return db
			},
			expectedError: errors.New("iterating over rows: sql: connection is already closed"),
		},
		{
			name: "error on scan",
			mockClosure: func() *sql.DB {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"q_ulid", "q_label", "q_text", "a_ulid", "a_text", "a_next_question_ulid"}).
					AddRow(nil, nil, false, "here", "here", "here")
				mock.ExpectQuery(regexp.QuoteMeta(getQuery)).
					WillReturnRows(rows)

				return db
			},
			expectedError: errors.New("scanning row: sql: Scan error on column index 0, name \"q_ulid\": converting NULL to string is unsupported"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := tc.mockClosure()
			output, err := Get(context.TODO(), db)
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
