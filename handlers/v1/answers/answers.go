// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package answers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/tiagomelo/questionnaire-rest-api/db/answers"
	"github.com/tiagomelo/questionnaire-rest-api/validate"
	"github.com/tiagomelo/questionnaire-rest-api/web"
)

// For ease of unit testing.
var (
	// getAnswerFlow retrieves the answer flow from the database.
	getAnswerFlow = answers.GetAnswerFlow

	// getRecommendations retrieves the recommendations from the database.
	getRecommendations = answers.RecommendationsFromAnswers

	// jsonDecode decodes a JSON request body into a given struct.
	jsonDecode = func(r io.Reader, v any) error {
		return json.NewDecoder(r).Decode(v)
	}

	// validateAnswerFlow validates the answers against the answer flow.
	_validateAnswerFlow = validateAnswerFlow
)

// GetRecommendationsRequest represents the request body for the GetRecommendations handler.
type GetRecommendationsRequest struct {
	Answers []string `json:"answers" validate:"required,min=1"`
}

// handlers struct holds a database connection.
type handlers struct {
	db *sql.DB
}

// New initializes a new instance of handlers with a database connection.
func New(db *sql.DB) *handlers {
	return &handlers{
		db: db,
	}
}

// RecommendationsFromAnswers is an HTTP handler that returns recommendations based on the answers provided.
func (h *handlers) RecommendationsFromAnswers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var getRecommReq GetRecommendationsRequest
	if err := jsonDecode(r.Body, &getRecommReq); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, errors.Wrap(err, "decoding json").Error())
		return
	}
	if err := validate.Check(getRecommReq); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	answerFlow, err := getAnswerFlow(r.Context(), h.db)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := _validateAnswerFlow(getRecommReq.Answers, answerFlow); err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	recommendations, err := getRecommendations(r.Context(), h.db, getRecommReq.Answers)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	web.RespondWithJson(w, http.StatusOK, recommendations)
}
