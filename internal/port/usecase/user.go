package usecase

import "anonsurvey/types/entity"

type User interface {
	CreateUser(entity.User) error
	ListUsers() error
}
