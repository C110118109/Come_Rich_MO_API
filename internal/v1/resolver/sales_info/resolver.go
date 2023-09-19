package sales_info

import (
	file "eirc.app/internal/v1/service/file"
	"eirc.app/internal/v1/service/sales_info"
	model "eirc.app/internal/v1/structure/sales_info"
	"gorm.io/gorm"
)

type Resolver interface {
	List(input *model.Fields) interface{}
	GetByID(input *model.Field) interface{}
}

type resolver struct {
	SalesInfoService sales_info.Service
	FileService      file.Service
}

func New(db *gorm.DB, db2 *gorm.DB) Resolver {
	return &resolver{
		SalesInfoService: sales_info.New(db),
		FileService:      file.New(db2),
	}
}
