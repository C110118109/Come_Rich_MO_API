package file

import (
	"eirc.app/internal/v1/entity/file"
	model "eirc.app/internal/v1/structure/file"
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Field) (err error)
	GetAllManuOrder() (output []*model.Base, err error)
}

type service struct {
	Entity file.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: file.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
