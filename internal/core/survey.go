package core

import (
	"anonsurvey/internal/port/outbound"
	"anonsurvey/types/entity"
)

type SurveyService struct {
	repo outbound.SurveyRepo
}

func (svc SurveyService) CreateSurvey(s entity.Survey) error {
	svc.repo.CreateSurvey(s)
	return nil
}

func (svc SurveyService) ListSurveys() []entity.Survey {
	models := svc.repo.ListSurveys()
	entities := make([]entity.Survey, 0)
	for _, m := range models {
		e, err := entity.NewSurvey().
			SetDetails(m.Details).
			SetRating(int(m.Rating)).
			SetSubmittedOn(m.PostedOn).
			Build()
		if err != nil {
			entities = append(entities, e)
		}
	}
	return entities
}
