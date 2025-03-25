package usecase

import "anonsurvey/types/entity"

type Survey interface {
	CreateSurvey(entity.Survey) error
	ListSurveys() []entity.Survey
}
