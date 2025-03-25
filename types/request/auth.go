package request

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SignIn struct {
	Email    string
	Password string
}

func (r SignIn) Validate() error {
	return v.ValidateStruct(&r,
		v.Field(&r.Email, is.Email),
		v.Field(&r.Password, v.Length(6, 32)),
	)
}
