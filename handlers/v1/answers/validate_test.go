// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package answers

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/questionnaire-rest-api/db/answers"
	"github.com/tiagomelo/questionnaire-rest-api/ptr"
)

func TestValidateAnswerFlow(t *testing.T) {
	cases := []struct {
		name          string
		answerULIDs   []string
		answerFlow    map[string]answers.AnswerFlow
		expectedError error
	}{
		{
			name:        "Valid flow",
			answerULIDs: []string{"A", "B"},
			answerFlow: map[string]answers.AnswerFlow{
				"A": {PreviousAnswers: []string{}},    // first answer, valid.
				"B": {PreviousAnswers: []string{"A"}}, // requires A before B.
			},
		},
		{
			name:          "empty answerULIDs",
			expectedError: errors.New("no answers provided"),
		},
		{
			name:        "terminal answer in sequence is allowed",
			answerULIDs: []string{"A", "Z"},
			answerFlow: map[string]answers.AnswerFlow{
				"A": {PreviousAnswers: []string{}}, // first answer, always valid.
				"Z": {PreviousAnswers: []string{}}, // terminal answer (no dependencies).
			},
		},
		{
			name:        "invalid answer ULID",
			answerULIDs: []string{"A", "B"},
			answerFlow: map[string]answers.AnswerFlow{
				"A": {PreviousAnswers: []string{}}, // first answer, valid.
			},
			expectedError: errors.New("invalid answer ULID: B"),
		},
		{
			name:        "missing required previous answer",
			answerULIDs: []string{"B"},
			answerFlow: map[string]answers.AnswerFlow{
				"B": {PreviousAnswers: []string{"A"}}, // requires A, but it's missing.
			},
			expectedError: errors.New("invalid answer sequence. Expected one of: [A] before B"),
		},
		{
			name:        "allow dynamic jump from B to D",
			answerULIDs: []string{"A", "B", "D"},
			answerFlow: map[string]answers.AnswerFlow{
				"A": {PreviousAnswers: []string{}},                      // first valid answer.
				"B": {PreviousAnswers: []string{"A"}, Next: ptr.P("D")}, // B normally goes to D.
				"C": {PreviousAnswers: []string{"B"}},                   // C requires B but is skipped.
				"D": {PreviousAnswers: []string{"C"}},                   // normally requires C, but allowed from B via `Next`.
			},
		},
		{
			name:        "Skipping optional branch explicitly allowed",
			answerULIDs: []string{"A", "D"},
			answerFlow: map[string]answers.AnswerFlow{
				"A": {PreviousAnswers: []string{}, Next: ptr.P("B")},    // A normally leads to B.
				"B": {PreviousAnswers: []string{"A"}, Next: ptr.P("C")}, // B normally leads to C.
				"C": {PreviousAnswers: []string{"B"}},                   // C normally requires B.
				"D": {PreviousAnswers: []string{"C", "A"}},              // ensure D allows skipping B and C.
			},
		},

		{
			name:        "Skipping optional step when parent allows it",
			answerULIDs: []string{"A", "D"},
			answerFlow: map[string]answers.AnswerFlow{
				"A": {PreviousAnswers: []string{}, Next: ptr.P("B")},    // A leads to B but it's optional.
				"B": {PreviousAnswers: []string{"A"}, Next: ptr.P("C")}, // B normally leads to C.
				"C": {PreviousAnswers: []string{"B"}},                   // C requires B explicitly.
				"D": {PreviousAnswers: []string{"C", "A"}, Next: nil},   // allowing D to come from either A or C.
			},
		},

		{
			name:        "Final answer without issues",
			answerULIDs: []string{"Z"},
			answerFlow: map[string]answers.AnswerFlow{
				"Z": {PreviousAnswers: []string{}}, // final answer.
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateAnswerFlow(tc.answerULIDs, tc.answerFlow)
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no error, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v", got nil`, tc.expectedError)
				}
			}

		})
	}
}
