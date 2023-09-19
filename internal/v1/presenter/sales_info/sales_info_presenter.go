package sales_info

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	salesInfoModel "eirc.app/internal/v1/structure/sales_info"
	"github.com/gin-gonic/gin"
)

// List
// @Summary 顯示顧客清單
func (p *presenter) List(ctx *gin.Context) {
	input := &salesInfoModel.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.SalesInfoResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一訂單資料
func (p *presenter) GetByID(ctx *gin.Context) {
	salesNO := ctx.Param("salesNO")
	input := &salesInfoModel.Field{}
	input.SalesNo = util.PointerString(salesNO)

	codeMessage := p.SalesInfoResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
