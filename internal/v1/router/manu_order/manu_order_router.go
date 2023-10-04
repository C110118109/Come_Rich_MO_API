package manu_order

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/manu_order"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("come-rich").Group("v1.0").Group("manu_order")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":manuOrderID", controller.GetByID)
		v10.DELETE(":manuOrderID", controller.Delete)
		v10.PATCH(":manuOrderID", controller.Updated)
	}

	return route
}
