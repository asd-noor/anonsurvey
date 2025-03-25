package model

type SurveyModeler interface {
	SurveyModel() Survey
}

type Survey struct {
	ID       int
	Rating   uint
	Details  string
	PostedOn string
}
