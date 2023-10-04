package manu_order

import (
	"eirc.app/internal/v1/resolver/manu_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	ManuOrderResolver manu_order.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		ManuOrderResolver: manu_order.New(db),
	}
}
