package raw_material

import (
	"eirc.app/internal/v1/resolver/raw_material"
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
	RawMaterialResolver raw_material.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		RawMaterialResolver: raw_material.New(db),
	}
}
