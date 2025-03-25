package adapter

import (
	"anonsurvey/internal/adapter/model"
)

type SurveyRepository struct{}

func NewSurveyRepository() SurveyRepository {
	return SurveyRepository{}
}

func (r SurveyRepository) CreateSurvey(s model.SurveyModeler) error {
	return nil
}
