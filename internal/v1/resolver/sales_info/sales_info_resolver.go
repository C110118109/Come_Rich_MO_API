package sales_info

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	fileModel "eirc.app/internal/v1/structure/file"
	salesInfoModel "eirc.app/internal/v1/structure/sales_info"
	"gorm.io/gorm"
)

func (r *resolver) List(input *salesInfoModel.Fields) interface{} {

	output := &salesInfoModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, salesInfos, err := r.SalesInfoService.List(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	allManuOrder, err := r.FileService.GetAllManuOrder()
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, item := range salesInfos {
		for _, manuOrder := range allManuOrder {
			if manuOrder.SalesNo == item.SalesNo {
				salesInfos[i].HasManuOrder = true
				break
			}
		}
	}

	salesInfosByte, err := json.Marshal(salesInfos)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	output.RecordNumbers = quantity
	err = json.Unmarshal(salesInfosByte, &output.SalesInfos)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *salesInfoModel.Field) interface{} {

	base, err := r.SalesInfoService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	//最後整理 合併製令下載路徑
	for i, item := range base.GoodsDetail {
		amount, findPath, err := r.FileService.List(&fileModel.Fields{Field: fileModel.Field{Identifier: &item.Identifier}})
		if err != nil {
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err)
		}
		if amount > 0 {
			base.GoodsDetail[i].ImgPath = findPath[0].PathKey
			base.GoodsDetail[i].FileId = findPath[0].FileID
		}
	}

	frontCustomer := &salesInfoModel.Base{}
	salesInfoByte, _ := json.Marshal(base)
	err = json.Unmarshal(salesInfoByte, &frontCustomer)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontCustomer)
}
