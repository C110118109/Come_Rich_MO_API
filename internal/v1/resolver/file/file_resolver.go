package file

import (
	"strings"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"

	"encoding/json"
	"errors"

	model "eirc.app/internal/v1/structure/file"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *model.Created) interface{} {
	defer trx.Rollback()
	//Todo 檢查重複

	file, err := r.FileService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, file.FileID)
}

func (r *resolver) List(input *model.Fields) interface{} {
	output := &model.List{}
	output.Limit = input.Limit
	output.Page = input.Page

	quantity, file, err := r.FileService.List(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	//組裝資料，查品名與訂單資料
	for i, item := range file {
		id := strings.Split(item.Identifier, "-")
		if len(id) == 2 {
			file[i].ItemName, err = r.SalesInfoService.GetItemNameByIDandSEQ(id[0], id[1])
			if err != nil {
				log.Error(err)
				return code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	fileByte, err := json.Marshal(file)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(fileByte, &output.Files)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *model.Field) interface{} {
	file, err := r.FileService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontMessage := &model.Base{}
	messageByte, _ := json.Marshal(file)
	err = json.Unmarshal(messageByte, &frontMessage)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontMessage)
}

func (r *resolver) Deleted(input *model.Field) interface{} {
	_, err := r.FileService.GetByID(&model.Field{FileID: input.FileID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.FileService.Deleted(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}
