package controller

import (
	"anonsurvey/internal/port/usecase"
	"anonsurvey/types/entity"
	"anonsurvey/types/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	svc usecase.User
}

func NewUserController(userSvc usecase.User) UserController {
	return UserController{
		svc: userSvc,
	}
}

func (c UserController) RegisterRoutes(e *echo.Echo) {
	e.POST("/user", c.CreateUser)
}

func (c UserController) CreateUser(ctx echo.Context) error {
	req := new(request.User)

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	u, err := entity.NewUser().
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetName(req.Name).
		Build()
	if err != nil {
		return err
	}

	if err := c.svc.CreateUser(u); err != nil {
		return err
	}

	resp := make(map[string]any)
	return ctx.JSON(http.StatusAccepted, resp)
}
