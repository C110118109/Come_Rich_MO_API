package login

import (
	"eirc.app/internal/v1/presenter/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := login.New(db)
	route.Group("come-rich").Group("v1.0").
		Group("login").POST("web", controller.Web)
	route.Group("come-rich").Group("v1.0").
		Group("refresh").POST("", controller.Refresh)

	return route
}
