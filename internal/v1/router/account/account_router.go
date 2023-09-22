package account

import (
	//"eirc.app/internal/v1/middleware"
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/account"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("come-rich").Group("v1.0").Group("account")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":accountID", controller.GetByID)
		v10.DELETE(":accountID", controller.Delete)
		v10.PATCH(":accountID", controller.Updated)
	}

	return route
}
