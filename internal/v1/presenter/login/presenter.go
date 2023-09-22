package login

import (
	"eirc.app/internal/v1/resolver/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Presenter interface {
	Web(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type presenter struct {
	LoginResolver login.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		LoginResolver: login.New(db),
	}
}
