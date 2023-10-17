package order

import (
	"net/http"

	order "eirc.app/internal/v1/resolver/order"
	"github.com/gin-gonic/gin"
)

func CreateExcel(ctx *gin.Context) {

	codeMessage := order.CreateExcel()
	ctx.JSON(http.StatusOK, codeMessage)
}
