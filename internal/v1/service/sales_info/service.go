package sales_info

import (
	entity "eirc.app/internal/v1/entity/sales_info"
	model "eirc.app/internal/v1/structure/sales_info"
	"gorm.io/gorm"
)

type Service interface {
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	GetItemNameByIDandSEQ(no string, seq string) (name string, err error) //查詢單獨項目
}

type service struct {
	Entity entity.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: entity.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
