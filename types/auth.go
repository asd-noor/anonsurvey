package types

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ReqSignIn struct {
	Email    string
	Password string
}

func (r ReqSignIn) Validate() error {
	return v.ValidateStruct(&r,
		v.Field(&r.Email, is.Email),
		v.Field(&r.Password, v.Length(6, 32)),
	)
}
