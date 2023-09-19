package sales_info

import (
	"eirc.app/internal/v1/presenter/sales_info"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB, db2 *gorm.DB) *gin.Engine {
	controller := sales_info.New(db, db2)
	v10 := route.Group("come-rich").Group("v1.0").Group("sales-info")
	{
		v10.GET("", controller.List)
		v10.GET(":salesNO", controller.GetByID)
	}

	return route
}
