package customer

import (
	entity "eirc.app/internal/v1/entity/customer"
	model "eirc.app/internal/v1/structure/customer"
	"gorm.io/gorm"
)

type Service interface {
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
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
