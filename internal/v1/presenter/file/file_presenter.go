package file

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	"eirc.app/internal/v1/structure/file"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增檔案
// @description 新增檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body file.Created true "新增檔案"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file [post]
func (p *presenter) Created(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	defer trx.Rollback()
	input := &file.Created{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	input.FileID = strconv.FormatInt(time.Now().Unix(), 10)
	id := strings.Split(input.Identifier, "-")
	if len(id) == 2 {
		input.SalesNo = id[0]
	}
	input.PathKey = input.Identifier + "_" + input.FileID + "." + input.Extension

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input.Base64))
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	b := buf.Bytes()

	err := ioutil.WriteFile("storage/"+input.PathKey, b, 0666)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusNotExtended, "上傳失敗")
	}

	input.PathKey = "/public/" + input.PathKey
	code := p.FileResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, code)

}

// List
// @Summary 條件搜尋檔案
// @description 條件檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=file.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &file.Fields{}

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}
	codeS3 := p.FileResolver.List(input)
	ctx.JSON(http.StatusOK, codeS3)
}

// GetByID
// @Summary 取得單一檔案
// @description 取得單一檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param FileID path string true "檔案ID"
// @success 200 object code.SuccessfulMessage{body=file.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/{FileID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	FileID := ctx.Param("FileID")
	input := &file.Field{}
	input.FileID = util.PointerString(FileID)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FileResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一檔案
// @description 刪除單一檔案
// @Tags File
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param FileID path string true "檔案ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/file/{FileID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	FileID := ctx.Param("FileID")

	input := &file.Field{}
	input.FileID = util.PointerString(FileID)

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.FileResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
