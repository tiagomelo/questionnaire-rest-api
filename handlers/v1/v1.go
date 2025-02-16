// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package v1

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiagomelo/questionnaire-rest-api/handlers/v1/answers"
	"github.com/tiagomelo/questionnaire-rest-api/handlers/v1/questionnaire"
	"github.com/tiagomelo/questionnaire-rest-api/middleware"
)

// Config struct holds the database connection and logger.
type Config struct {
	JwtKey string
	Db     *sql.DB
	Log    *slog.Logger
}

// Routes initializes and returns a new router with configured routes.
func Routes(c *Config) *mux.Router {
	router := mux.NewRouter()
	initializeRoutes(c.Db, router)
	router.Use(
		func(h http.Handler) http.Handler {
			return middleware.Logger(c.Log, h)
		},
		middleware.Compress,
		middleware.PanicRecovery,
	)
	return router
}

// initializeRoutes sets up the routes for book operations.
func initializeRoutes(db *sql.DB, router *mux.Router) {
	questionnaireHandlers := questionnaire.New(db)
	answersHandlers := answers.New(db)
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/questionnaire", questionnaireHandlers.Get).Methods(http.MethodGet)
	apiRouter.HandleFunc("/questionnaire/answers", answersHandlers.RecommendationsFromAnswers).Methods(http.MethodPost)
}
