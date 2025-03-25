package adapter

import (
	"anonsurvey/internal/adapter/model"
	"anonsurvey/internal/port/outbound"
)

type UserRepository struct {
	db outbound.DBClient
}

func NewUserRepository(db outbound.DBClient) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) GetUserByEmail(email string) model.User {
	return model.User{}
}

func (r UserRepository) CreateUser(user model.UserModeler) error {
	_ = user.UserModel()
	return nil
}

func (us UserRepository) ListUsers() error {
	panic("not implemented") // TODO: Implement
}
