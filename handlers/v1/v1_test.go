// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package v1_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/questionnaire-rest-api/config"
	"github.com/tiagomelo/questionnaire-rest-api/db"
	"github.com/tiagomelo/questionnaire-rest-api/handlers"
)

var (
	testDb     *sql.DB
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	cfg, err := config.ReadFromEnvFile("../../.env")
	if err != nil {
		fmt.Println("error when reading configuration:", err)
		os.Exit(1)
	}
	testDb, err = db.ConnectToPsql(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresTestHost, cfg.PostgresDb)
	if err != nil {
		fmt.Println("error when connecting to the test database:", err)
		os.Exit(1)
	}
	log := slog.New(slog.NewJSONHandler(io.Discard, nil))
	apiMux := handlers.NewApiMux(&handlers.ApiMuxConfig{
		Db:  testDb,
		Log: log,
	})

	testServer = httptest.NewServer(apiMux)
	defer testServer.Close()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestV1GetQuestionnaire(t *testing.T) {
	expectedOutput, err := os.ReadFile("../../testdata/expected_questionnaire.json")
	require.NoError(t, err)
	resp, err := http.Get(testServer.URL + "/api/v1/questionnaire")
	require.NoError(t, err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.JSONEq(t, string(expectedOutput), string(b), "Expected and actual JSON do not match")
}

func TestV1PostAnswers(t *testing.T) {
	testCases := []struct {
		name           string
		inputFilePath  string
		outputFilePath string
	}{
		{
			name:           "recommends Sildenafil 50mg",
			inputFilePath:  "../../testdata/answers/input/recommends_sildenafil_50mg.json",
			outputFilePath: "../../testdata/answers/output/recommends_sildenafil_50mg.json",
		},
		{
			name:           "recommends Tadalafil 20mg",
			inputFilePath:  "../../testdata/answers/input/recommends_tadalafil_20mg.json",
			outputFilePath: "../../testdata/answers/output/recommends_tadalafil_20mg.json",
		},
		{
			name:           "recommends Tadalafil 10mg",
			inputFilePath:  "../../testdata/answers/input/recommends_tadalafil_10mg.json",
			outputFilePath: "../../testdata/answers/output/recommends_tadalafil_10mg.json",
		},
		{
			name:           "recommends Sildenafil 100mg",
			inputFilePath:  "../../testdata/answers/input/recommends_sildenafil_100mg.json",
			outputFilePath: "../../testdata/answers/output/recommends_sildenafil_100mg.json",
		},
		{
			name:           "recommends Sildenafil 100mg as preferred treatment",
			inputFilePath:  "../../testdata/answers/input/recommends_sildenafil_100mg_preferred.json",
			outputFilePath: "../../testdata/answers/output/recommends_sildenafil_100mg_preferred.json",
		},
		{
			name:           "recommends Tadalafil 20mg as preferred treatment",
			inputFilePath:  "../../testdata/answers/input/recommends_tadalafil_20mg_preferred.json",
			outputFilePath: "../../testdata/answers/output/recommends_tadalafil_20mg_preferred.json",
		},
		{
			name:           "recommends Sildenafil 100mg or Tadalafil 20mg as preferred treatment",
			inputFilePath:  "../../testdata/answers/input/recommends_sildenafil_100mg_tadalafil_20mg_preferred.json",
			outputFilePath: "../../testdata/answers/output/recommends_sildenafil_100mg_tadalafil_20mg_preferred.json",
		},
		{
			name:           "recommends Sildenafil 50mg or Tadalafil 10mg",
			inputFilePath:  "../../testdata/answers/input/recommends_sildenafil_50mg_tadalafil_10mg.json",
			outputFilePath: "../../testdata/answers/output/recommends_sildenafil_50mg_tadalafil_10mg.json",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := os.ReadFile(tc.inputFilePath)
			require.NoError(t, err)
			expectedOutput, err := os.ReadFile(tc.outputFilePath)
			require.NoError(t, err)
			resp, err := http.Post(testServer.URL+"/api/v1/questionnaire/answers", "application/json", bytes.NewBuffer([]byte(input)))
			require.NoError(t, err)
			defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.JSONEq(t, string(expectedOutput), string(b), "Expected and actual JSON do not match")
		})
	}
}

func TestV1PostAnswers_validation_errors(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		outputFilePath string
	}{
		{
			name:           "missing answers",
			input:          "{}",
			outputFilePath: "../../testdata/answers/output/missing_answers.json",
		},
		{
			name:           "empty answers",
			input:          `{"answers": []}`,
			outputFilePath: "../../testdata/answers/output/empty_answers.json",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expectedOutput, err := os.ReadFile(tc.outputFilePath)
			require.NoError(t, err)
			resp, err := http.Post(testServer.URL+"/api/v1/questionnaire/answers", "application/json", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)
			defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.JSONEq(t, string(expectedOutput), string(b), "Expected and actual JSON do not match")
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
	}
}

func TestV1PostAnswers_exclusions(t *testing.T) {
	groupedTestCases := map[string][]struct {
		name           string
		inputFilePath  string
		outputFilePath string
	}{
		"User responded 'No' to Q1 (No erection difficulty)": {
			{
				name:           "no difficulty getting or maintaining an erection",
				inputFilePath:  "../../testdata/answers/input/exclusion_no_difficulty.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded 'Yes' to Q3 (Heart or neurological conditions)": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_preferred_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_preferred_heart_neuro_conds.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded 'Significant liver problems (such as cirrhosis of the liver) or kidney problems' to Q4": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_significant.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded 'Currently prescribed GTN...' to Q4": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_gtn.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded 'Abnormal blood pressure' to Q4": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_abnormal.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded with condition to Q4": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_condition.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded with some drug to Q5": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_alpha.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded with 'Riociguat...' to Q5": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_rio.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded with 'Saquinavir...' to Q5": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
		"User responded with 'Cimetidine...' to Q5": {
			{
				name:           "should recommend Sildenafil 50mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 10mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_10mg_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_preferred_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_tadalafil_20mg_preferred_saqui.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 100mg or Tadalafil 20mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_100mg_tadalafil_20mg_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
			{
				name:           "should recommend Sildenafil 50mg or Tadalafil 10mg as preferred treatment but excludes it",
				inputFilePath:  "../../testdata/answers/input/excludes_sildenafil_50mg_tadalafil_10mg_cime.json",
				outputFilePath: "../../testdata/answers/output/no_products_available.json",
			},
		},
	}
	for group, testCases := range groupedTestCases {
		t.Run(group, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					input, err := os.ReadFile(tc.inputFilePath)
					require.NoError(t, err)

					expectedOutput, err := os.ReadFile(tc.outputFilePath)
					require.NoError(t, err)

					resp, err := http.Post(testServer.URL+"/api/v1/questionnaire/answers", "application/json", bytes.NewBuffer([]byte(input)))
					require.NoError(t, err)
					defer resp.Body.Close()

					b, err := io.ReadAll(resp.Body)
					require.NoError(t, err)

					require.JSONEq(t, string(expectedOutput), string(b), "Expected and actual JSON do not match")
				})
			}
		})
	}
}

func TestV1PostAnswers_invalid_answers_sequence(t *testing.T) {
	testCases := []struct {
		name           string
		inputFilePath  string
		outputFilePath string
	}{
		{
			name:           "invalid backward flow (e.g. Q3 -> Q1)",
			inputFilePath:  "../../testdata/answers/input/invalid_backward_flow.json",
			outputFilePath: "../../testdata/answers/output/error_backward_flow.json",
		},
		{
			name:           "invalid forward flow (e.g. Q1 -> Q5)",
			inputFilePath:  "../../testdata/answers/input/invalid_forward_flow.json",
			outputFilePath: "../../testdata/answers/output/error_forward_flow.json",
		},
		{
			name:           "invalid answer ULID",
			inputFilePath:  "../../testdata/answers/input/invalid_answer_uuid.json",
			outputFilePath: "../../testdata/answers/output/error_non_existent_answer.json",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := os.ReadFile(tc.inputFilePath)
			require.NoError(t, err)

			expectedOutput, err := os.ReadFile(tc.outputFilePath)
			require.NoError(t, err)

			resp, err := http.Post(testServer.URL+"/api/v1/questionnaire/answers", "application/json", bytes.NewBuffer([]byte(input)))
			require.NoError(t, err)
			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			require.JSONEq(t, string(expectedOutput), string(b), "Expected and actual JSON do not match")
		})
	}
}
