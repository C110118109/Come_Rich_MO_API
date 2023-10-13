package raw_material

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/raw_material"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增使用者
// @description 新增使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body raw_material.Created true "新增使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/raw_material [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &raw_material.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.RawMaterialResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 條件搜尋使用者
// @description 條件使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param organizationID query string false "組織ID"
// @param raw_material query string false "帳號"
// @param chineseName query string false "中文名稱"
// @param roleName query string false "角色名稱"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=raw_material.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/raw_material [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &raw_material.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.RawMaterialResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一使用者
// @description 取得單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param manuOrderID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=manu_order.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/manu_order/{manuOrderID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	rawMaterialID := ctx.Param("rawMaterialID")
	//撈取新物件，套上搜尋值，執行搜尋
	input := &raw_material.Field{}
	input.RawMaterialID = util.PointerString(rawMaterialID)

	codeMessage := p.RawMaterialResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一使用者
// @description 刪除單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param manuOrderID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/raw_material/{manuOrderID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	rawMaterialID := ctx.Param("rawMaterialID")
	input := &raw_material.Updated{}
	input.RawMaterialID = rawMaterialID

	codeMessage := p.RawMaterialResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一使用者
// @description 更新單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param manuOrderID path string true "使用者ID"
// @param * body manu_order.Updated true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/raw_material/{manuOrderID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	//updatedBy := ctx.Value("userId").(string)
	rawMaterialID := ctx.Param("rawMaterialID")
	input := &raw_material.Updated{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	input.RawMaterialID = rawMaterialID
	codeMessage := p.RawMaterialResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
