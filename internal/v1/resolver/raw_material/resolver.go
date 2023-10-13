package raw_material

import (
	"eirc.app/internal/v1/service/raw_material"
	model "eirc.app/internal/v1/structure/raw_material"
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
	RawMaterialService raw_material.Service
}

func New(db *gorm.DB) Resolver {

	return &resolver{
		RawMaterialService: raw_material.New(db),
	}
}
