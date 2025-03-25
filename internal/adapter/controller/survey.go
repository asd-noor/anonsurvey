package controller

import (
	"anonsurvey/internal/port/usecase"
	"anonsurvey/types/entity"
	"anonsurvey/types/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SurveyController struct {
	svc usecase.Survey
}

func NewSurveyController() SurveyController {
	return SurveyController{}
}

func (c SurveyController) RegisterRoutes(e *echo.Echo) {
	e.POST("/survey", c.CreateSurvey)
}

func (c SurveyController) CreateSurvey(ctx echo.Context) error {
	req := new(request.Survey)

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	s, err := entity.NewSurvey().
		SetDetails(req.Details).
		SetRating(req.Rating).
		SetSubmittedOn(req.SubmittedOn).
		Build()
	if err != nil {
		return err
	}

	if err := c.svc.CreateSurvey(s); err != nil {
		return err
	}

	resp := make(map[string]any)
	return ctx.JSON(http.StatusAccepted, resp)
}

func (c SurveyController) ListSurveys(ctx echo.Context) error {
	resp := make(map[string]any)
	return ctx.JSON(http.StatusAccepted, resp)
}
