package customer

import (
	"eirc.app/internal/v1/service/customer"
	model "eirc.app/internal/v1/structure/customer"
	"gorm.io/gorm"
)

type Resolver interface {
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
}

type resolver struct {
	CustomerService customer.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		CustomerService: customer.New(db),
	}
}
