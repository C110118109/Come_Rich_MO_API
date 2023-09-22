package login

import (
	"eirc.app/internal/v1/service/account"
	"eirc.app/internal/v1/service/jwe"

	jweModel "eirc.app/internal/v1/structure/jwe"
	loginModel "eirc.app/internal/v1/structure/login"
	"gorm.io/gorm"
)

type Resolver interface {
	Web(input *loginModel.Login) interface{}
	Refresh(input *jweModel.Refresh) interface{}
}

type resolver struct {
	Account account.Service
	JWE     jwe.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		Account: account.New(db),
		JWE:     jwe.New(),
	}
}
