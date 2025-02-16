// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package answers

import (
	"github.com/pkg/errors"
	"github.com/tiagomelo/questionnaire-rest-api/db/answers"
)

// ValidateAnswerFlow validates the answer flow based on the provided answers.
func validateAnswerFlow(answerULIDs []string, answerFlow map[string]answers.AnswerFlow) error {
	if len(answerULIDs) == 0 {
		return errors.New("no answers provided")
	}
	answeredSet := make(map[string]bool)
	for i, ansULID := range answerULIDs {
		flow, exists := answerFlow[ansULID]
		if !exists {
			return errors.Errorf("invalid answer ULID: %s", ansULID)
		}
		// first answer is only valid if it doesn't require previous answers.
		if i == 0 && len(flow.PreviousAnswers) == 0 {
			answeredSet[ansULID] = true
			continue
		}
		// if no required previous answers, allow it (Terminal Answers).
		if len(flow.PreviousAnswers) == 0 {
			answeredSet[ansULID] = true
			continue
		}
		// alow valid previous answers.
		validPreviousFound := false
		for _, prev := range flow.PreviousAnswers {
			if answeredSet[prev] {
				validPreviousFound = true
				break
			}
		}
		// allow dynamic jumps from `answerFlow`.
		if !validPreviousFound {
			for prevAns, prevFlow := range answerFlow {
				if prevFlow.Next != nil && *prevFlow.Next == ansULID && answeredSet[prevAns] {
					validPreviousFound = true
					break
				}
			}
		}
		if !validPreviousFound {
			return errors.Errorf("invalid answer sequence. Expected one of: %v before %s", flow.PreviousAnswers, ansULID)
		}
		answeredSet[ansULID] = true
	}
	return nil
}
