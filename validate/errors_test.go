// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package validate

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	testCases := []struct {
		name            string
		input           FieldErrors
		mockJsonMarshal func(v any) ([]byte, error)
		expectedOutput  string
	}{
		{
			name: "happy path",
			input: FieldErrors{
				{
					Field: "field",
					Error: "error",
				},
			},
			expectedOutput: `[{"field":"field","error":"error"}]`,
		},
		{
			name: "json marshal error",
			input: FieldErrors{
				{
					Field: "field",
					Error: "error",
				},
			},
			mockJsonMarshal: func(v any) ([]byte, error) {
				return nil, errors.New("json marshal error")
			},
			expectedOutput: "json marshal error",
		},
	}
	originalJsonMarshal := jsonMarshal
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() { jsonMarshal = originalJsonMarshal }()
			if tc.mockJsonMarshal != nil {
				jsonMarshal = tc.mockJsonMarshal
			}
			output := tc.input.Error()
			if output != tc.expectedOutput {
				t.Errorf("expected %s, got %s", tc.expectedOutput, output)
			}
		})
	}
}
