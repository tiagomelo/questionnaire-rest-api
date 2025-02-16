// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package questionnaire

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/questionnaire-rest-api/db/questionnaire"
	"github.com/tiagomelo/questionnaire-rest-api/ptr"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name                 string
		mockGetQuestionnaire func(ctx context.Context, db *sql.DB) (*questionnaire.Questionnaire, error)
		expectedOutput       string
		expectedStatusCode   int
	}{
		{
			name: "happy path",
			mockGetQuestionnaire: func(ctx context.Context, db *sql.DB) (*questionnaire.Questionnaire, error) {
				return &questionnaire.Questionnaire{
					Questions: []*questionnaire.Question{
						{
							ULID:  "ulid-1",
							Label: "label-1",
							Text:  "text-1",
							Answers: []*questionnaire.Answer{
								{
									ULID:             "ulid-1",
									Text:             "text-1",
									NextQuestionULID: ptr.P("ulid-2"),
								},
							},
						},
						{
							ULID:  "ulid-2",
							Label: "label-2",
							Text:  "text-2",
							Answers: []*questionnaire.Answer{
								{
									ULID:             "ulid-1",
									Text:             "text-1",
									NextQuestionULID: nil,
								},
							},
						},
					},
				}, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedOutput:     "{\"questions\":[{\"ulid\":\"ulid-1\",\"label\":\"label-1\",\"text\":\"text-1\",\"answers\":[{\"ulid\":\"ulid-1\",\"text\":\"text-1\",\"next_question_ulid\":\"ulid-2\"}]},{\"ulid\":\"ulid-2\",\"label\":\"label-2\",\"text\":\"text-2\",\"answers\":[{\"ulid\":\"ulid-1\",\"text\":\"text-1\",\"next_question_ulid\":null}]}]}",
		},
		{
			name: "error",
			mockGetQuestionnaire: func(ctx context.Context, db *sql.DB) (*questionnaire.Questionnaire, error) {
				return nil, errors.New("some error")
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedOutput:     "{\"error\":\"some error\"}",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			getQuestionnaire = tc.mockGetQuestionnaire
			req, err := http.NewRequest(http.MethodGet, "questionnaire", nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			h := New(nil)
			handler := http.HandlerFunc((h).Get)
			handler.ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatusCode, rr.Code)
			require.Equal(t, tc.expectedOutput, rr.Body.String())
		})
	}
}
