package manu_order

import (
	"eirc.app/internal/v1/service/manu_order"
	model "eirc.app/internal/v1/structure/manu_order"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Updated) interface{}
	Updated(input *model.Updated) interface{}
}

type resolver struct {
	ManuOrderService manu_order.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		ManuOrderService: manu_order.New(db),
	}
}
