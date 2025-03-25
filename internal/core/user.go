package core

import (
	"anonsurvey/internal/port/outbound"
	"anonsurvey/types/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo outbound.UserRepo
}

func NewUserService(userRepo outbound.UserRepo) UserService {
	return UserService{}
}

func (s UserService) SignIn(u entity.User) error {
	u2 := s.repo.GetUserByEmail(u.Email())
	if !s.isSamePassword(u.Password(), u2.Password) {
		return errors.New("invalid password")
	}
	return nil
}

func (s UserService) SignOut() {}

func (s UserService) CreateUser(u entity.User) error {
	hashedPass := s.generatePassword(u.Password())
	u2, err := entity.ModifyUser(u).SetPassword(hashedPass).Build()
	if err != nil {
		return err
	}
	if err := s.repo.CreateUser(u2); err != nil {
		return err
	}

	return nil
}

func (s UserService) ListUsers() error {
	return nil
}

func (s UserService) generatePassword(plainPass string) string {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	return string(hashedPass)
}

func (s UserService) isSamePassword(plainPass string, hashedPass string) bool {
	hb, pb := []byte(hashedPass), []byte(plainPass)
	return bcrypt.CompareHashAndPassword(hb, pb) == nil
}
