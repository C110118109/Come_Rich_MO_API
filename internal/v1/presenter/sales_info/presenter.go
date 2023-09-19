package sales_info

import (
	"eirc.app/internal/v1/resolver/sales_info"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type presenter struct {
	SalesInfoResolver sales_info.Resolver
}

func New(db *gorm.DB, db2 *gorm.DB) Presenter {
	return &presenter{
		SalesInfoResolver: sales_info.New(db, db2),
	}
}
