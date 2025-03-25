package entity

import (
	"anonsurvey/internal/adapter/model"
	"fmt"
)

type Survey struct {
	Details     string
	Rating      uint
	SubmittedOn string
}

func (s Survey) SurveyModel() model.Survey {
	// TODO: implement
	return model.Survey{}
}

type surveyBuilder struct {
	s       Survey
	setters []func() error
}

func NewSurvey() *surveyBuilder {
	return &surveyBuilder{
		s: Survey{},
	}
}

func (b *surveyBuilder) SetDetails(details string) *surveyBuilder {
	fn := func() error {
		if details == "" {
			return fmt.Errorf("empty value")
		}
		b.s.Details = details
		return nil
	}

	b.setters = append(b.setters, fn)
	return b
}

func (b *surveyBuilder) SetRating(rating int) *surveyBuilder {
	fn := func() error {
		if rating < 1 || rating > 5 {
			return fmt.Errorf("invalid rating")
		}
		b.s.Rating = uint(rating)
		return nil
	}

	b.setters = append(b.setters, fn)
	return b
}

func (b *surveyBuilder) SetSubmittedOn(date string) *surveyBuilder {
	fn := func() error {
		if date == "" {
			return fmt.Errorf("empty value")
		}
		b.s.SubmittedOn = date
		return nil
	}

	b.setters = append(b.setters, fn)
	return b
}

func (b *surveyBuilder) Build() (Survey, error) {
	for _, fn := range b.setters {
		if err := fn(); err != nil {
			return Survey{}, err
		}
	}

	return b.s, nil
}
