// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package answers

import (
	"context"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/questionnaire-rest-api/db/answers"
)

func TestRecommendationsFromAnswers(t *testing.T) {
	testCases := []struct {
		name                   string
		input                  string
		mockGetAnswerFlow      func(ctx context.Context, db *sql.DB) (map[string]answers.AnswerFlow, error)
		mockValidateAnswerFlow func(answerUUIDs []string, answerFlow map[string]answers.AnswerFlow) error
		mockGetRecommendations func(ctx context.Context, db *sql.DB, answerUUIDs []string) (*answers.RecommendationsResponse, error)
		mockJsonDecode         func(r io.Reader, v any) error
		expectedOutput         string
		expectedStatusCode     int
	}{
		{
			name:  "happy path",
			input: `{"answers":["ulid-1"]}`,
			mockGetAnswerFlow: func(ctx context.Context, db *sql.DB) (map[string]answers.AnswerFlow, error) {
				return map[string]answers.AnswerFlow{}, nil
			},
			mockValidateAnswerFlow: func(answerUUIDs []string, answerFlow map[string]answers.AnswerFlow) error {
				return nil
			},
			mockGetRecommendations: func(ctx context.Context, db *sql.DB, answerUUIDs []string) (*answers.RecommendationsResponse, error) {
				return &answers.RecommendationsResponse{
					Recommendations: []string{"no recommendations"},
				}, nil
			},
			expectedOutput:     `{"recommendations":["no recommendations"]}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "error decoding json",
			mockJsonDecode: func(r io.Reader, v any) error {
				return io.EOF
			},
			expectedOutput:     `{"error":"decoding json: EOF"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "error validating request - missing answers",
			input:              `{}`,
			expectedOutput:     `{"error":"[{\"field\":\"answers\",\"error\":\"answers is a required field\"}]"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "error validating request - empty answers array",
			input:              `{"answers":[]}`,
			expectedOutput:     `{"error":"[{\"field\":\"answers\",\"error\":\"answers must contain at least 1 item\"}]"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "error validating answers",
			input: `{"answers":["ulid-1"]}`,
			mockGetAnswerFlow: func(ctx context.Context, db *sql.DB) (map[string]answers.AnswerFlow, error) {
				return map[string]answers.AnswerFlow{}, nil
			},
			mockValidateAnswerFlow: func(answerUUIDs []string, answerFlow map[string]answers.AnswerFlow) error {
				return errors.New("validation error")
			},
			expectedOutput:     `{"error":"validation error"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "error getting answer flow",
			input: `{"answers":["ulid-1"]}`,
			mockGetAnswerFlow: func(ctx context.Context, db *sql.DB) (map[string]answers.AnswerFlow, error) {
				return nil, errors.New("error getting answer flow")
			},
			expectedOutput:     `{"error":"error getting answer flow"}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "error getting recommendations",
			input: `{"answers":["ulid-1"]}`,
			mockGetAnswerFlow: func(ctx context.Context, db *sql.DB) (map[string]answers.AnswerFlow, error) {
				return map[string]answers.AnswerFlow{}, nil
			},
			mockValidateAnswerFlow: func(answerUUIDs []string, answerFlow map[string]answers.AnswerFlow) error {
				return nil
			},
			mockGetRecommendations: func(ctx context.Context, db *sql.DB, answerUUIDs []string) (*answers.RecommendationsResponse, error) {
				return nil, errors.New("error getting recommendations")
			},
			expectedOutput:     `{"error":"error getting recommendations"}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	originalJsonDecode := jsonDecode
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				jsonDecode = originalJsonDecode
			}()
			_validateAnswerFlow = tc.mockValidateAnswerFlow
			getAnswerFlow = tc.mockGetAnswerFlow
			getRecommendations = tc.mockGetRecommendations
			if tc.mockJsonDecode != nil {
				jsonDecode = tc.mockJsonDecode
			}
			req, err := http.NewRequest(http.MethodPost, "answers", strings.NewReader(tc.input))
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			h := New(nil)
			handler := http.HandlerFunc((h).RecommendationsFromAnswers)
			handler.ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatusCode, rr.Code)
			require.Equal(t, tc.expectedOutput, rr.Body.String())
		})
	}
}
