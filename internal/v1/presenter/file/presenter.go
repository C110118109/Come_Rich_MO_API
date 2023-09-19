package file

import (
	"eirc.app/internal/v1/resolver/file"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type presenter struct {
	FileResolver file.Resolver
}

func New(db *gorm.DB, db2 *gorm.DB) Presenter {
	return &presenter{
		FileResolver: file.New(db, db2),
	}
}
