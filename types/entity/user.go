package entity

import (
	"anonsurvey/internal/adapter/model"
	"anonsurvey/types"
	"errors"
)

type User struct {
	email      string
	name       string
	password   string
	permission types.Permission
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) UserModel() model.User {
	x := model.User{
		Name:     u.name,
		Email:    u.email,
		Password: u.password,
	}

	return x
}

type userBuilder struct {
	u        User
	setFuncs []func() error
}

func NewUser() *userBuilder {
	return &userBuilder{
		u: User{},
	}
}

func ModifyUser(user User) *userBuilder {
	return &userBuilder{
		u: user,
	}
}

func (b *userBuilder) SetName(name string) *userBuilder {
	setName := func() error {
		if name == "" {
			return errors.New("empty name")
		}

		b.u.name = name
		return nil
	}

	b.setFuncs = append(b.setFuncs, setName)
	return b
}

func (b *userBuilder) SetEmail(email string) *userBuilder {
	setEmail := func() error {
		if email == "" {
			return errors.New("empty email")
		}

		b.u.email = email
		return nil
	}

	b.setFuncs = append(b.setFuncs, setEmail)
	return b
}

func (b *userBuilder) SetPassword(password string) *userBuilder {
	setPass := func() error {
		if password == "" {
			return errors.New("empty password")
		}
		b.u.password = password
		return nil
	}

	b.setFuncs = append(b.setFuncs, setPass)
	return b
}

func (b *userBuilder) Build() (User, error) {
	for _, fn := range b.setFuncs {
		if err := fn(); err != nil {
			return User{}, err
		}
	}
	return b.u, nil
}
