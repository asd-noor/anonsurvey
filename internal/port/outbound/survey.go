package outbound

import "anonsurvey/internal/adapter/model"

type SurveyRepo interface {
	CreateSurvey(model.SurveyModeler) error
	ListSurveys() []model.Survey
}
