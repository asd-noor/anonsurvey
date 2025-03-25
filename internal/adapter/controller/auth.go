package controller

import (
	"anonsurvey/internal/port/usecase"
	"anonsurvey/types/entity"
	"anonsurvey/types/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	svc usecase.Auth
}

func NewAuthController(svc usecase.Auth) AuthController {
	return AuthController{
		svc: svc,
	}
}

func (c AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/sign-in", c.signIn)
}

func (c AuthController) signIn(ctx echo.Context) error {
	req := new(request.SignIn)
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	u, err := entity.NewUser().
		SetEmail(req.Email).
		SetPassword(req.Password).
		Build()
	if err != nil {
		return err
	}

	if err := c.svc.SignIn(u); err != nil {
		return err
	}

	res := make(map[string]any)
	return ctx.JSON(http.StatusOK, res)
}
