package customer

import (
	"eirc.app/internal/v1/presenter/customer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := customer.New(db)
	v10 := route.Group("come-rich").Group("v1.0").Group("customer")
	{
		v10.GET("", controller.List)
		v10.GET(":cusNO", controller.GetByID)
	}

	return route
}
