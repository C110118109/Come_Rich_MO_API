package raw_material

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/raw_material"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("come-rich").Group("v1.0").Group("raw_material")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":rawMaterialID", controller.GetByID)
		v10.DELETE(":rawMaterialID", controller.Delete)
		v10.PATCH(":rawMaterialID", controller.Updated)
	}

	return route
}
