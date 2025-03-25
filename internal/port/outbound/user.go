package outbound

import "anonsurvey/internal/adapter/model"

type UserRepo interface {
	CreateUser(model.UserModeler) error
	GetUserByEmail(string) model.User
	ListUsers() error
}
