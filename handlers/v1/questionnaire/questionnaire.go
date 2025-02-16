// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package questionnaire

import (
	"database/sql"
	"net/http"

	"github.com/tiagomelo/questionnaire-rest-api/db/questionnaire"
	"github.com/tiagomelo/questionnaire-rest-api/web"
)

// For ease of unit testing.
var getQuestionnaire = questionnaire.Get

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

// Get returns the questionnaire.
func (h *handlers) Get(w http.ResponseWriter, r *http.Request) {
	questionnaire, err := getQuestionnaire(r.Context(), h.db)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	web.RespondWithJson(w, http.StatusOK, questionnaire)
}
