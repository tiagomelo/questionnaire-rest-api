package doc

import (
	answersModel "github.com/tiagomelo/questionnaire-rest-api/db/answers"
	"github.com/tiagomelo/questionnaire-rest-api/db/questionnaire"
	"github.com/tiagomelo/questionnaire-rest-api/handlers/v1/answers"
)

// swagger:route GET /api/v1/questionnaire questionnaire questionnaire
// Get questionnaire.
// ---
// responses:
//		200: getQuestionnaireResponse
//		500: description: internal server error

// swagger:response getQuestionnaireResponse
type GetQuestionnaireResponse struct {
	// in:body
	Body *questionnaire.Questionnaire
}

// swagger:route POST /api/v1/questionnaire/answers answers Answer
// Answers the questionnaire.
// ---
// responses:
//		201: answerQuestionnaireResponse
//		400: description: missing required fields
//		500: description: internal server error

// swagger:parameters Answer
type PostQuestionnaireParamsWrapper struct {
	// in:body
	Body answers.GetRecommendationsRequest
}

// swagger:response answerQuestionnaireResponse
type AnswerQuestionnaireResponseWrapper struct {
	// in:body
	Body *answersModel.RecommendationsResponse
}
