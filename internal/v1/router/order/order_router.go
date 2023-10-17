package order

import (
	controller "eirc.app/internal/v1/presenter/order"
	"github.com/gin-gonic/gin"
)

func GetRoute(route *gin.Engine) *gin.Engine {
	v10 := route.Group("come-rich").Group("v1.0").Group("order")
	{

		v10.GET("", controller.CreateExcel)
	}
	//產生PI excel檔案
	return route
}
