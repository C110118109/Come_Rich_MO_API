package file

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/file"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB, db2 *gorm.DB) *gin.Engine {
	controller := presenter.New(db, db2)
	v10 := route.Group("come-rich").Group("v1.0").Group("file")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":FileID", controller.GetByID)
		v10.DELETE(":FileID", controller.Delete)
	}

	return route
}
