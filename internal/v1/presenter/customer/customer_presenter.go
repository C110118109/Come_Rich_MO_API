package customer

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/customer"
	"github.com/gin-gonic/gin"
)

// List
// @Summary 顯示顧客清單
func (p *presenter) List(ctx *gin.Context) {
	input := &customer.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.CustomerResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一顧客資料
func (p *presenter) GetByID(ctx *gin.Context) {
	customerID := ctx.Param("cusNO")
	input := &customer.Field{}
	input.No = util.PointerString(customerID)

	codeMessage := p.CustomerResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
