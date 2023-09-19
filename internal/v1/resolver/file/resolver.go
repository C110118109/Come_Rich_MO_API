package file

import (
	"eirc.app/internal/v1/service/file"
	sales_info "eirc.app/internal/v1/service/sales_info"
	model "eirc.app/internal/v1/structure/file"
	"gorm.io/gorm"
)

type Resolver interface {
	Created(trx *gorm.DB, input *model.Created) interface{}
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
	Deleted(input *model.Field) interface{}
}

type resolver struct {
	FileService      file.Service
	SalesInfoService sales_info.Service
}

func New(db *gorm.DB, db2 *gorm.DB) Resolver {

	return &resolver{
		FileService:      file.New(db),
		SalesInfoService: sales_info.New(db2),
	}
}
