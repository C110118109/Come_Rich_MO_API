package login

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/v1/structure/jwe"
	loginModel "eirc.app/internal/v1/structure/login"
	"github.com/gin-gonic/gin"
)

// Web
// @Summary 使用者登入
// @description 使用者登入
// @Tags Login
// @version 1.0
// @Accept json
// @produce json
// @param * body logins.Login true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/login/web [post]
func (p *presenter) Web(ctx *gin.Context) {

	input := &loginModel.Login{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := p.LoginResolver.Web(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Refresh
// @Summary 換新的令牌
// @description 換新的令牌
// @Tags Login
// @version 1.0
// @Accept json
// @produce json
// @param * body jwe.Refresh true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/login/refresh [post]
func (p *presenter) Refresh(ctx *gin.Context) {
	input := &jwe.Refresh{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := p.LoginResolver.Refresh(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
