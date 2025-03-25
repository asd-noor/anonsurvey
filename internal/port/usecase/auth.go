package usecase

import "anonsurvey/types/entity"

type Auth interface {
	SignIn(entity.User) error
	SignOut()
}
