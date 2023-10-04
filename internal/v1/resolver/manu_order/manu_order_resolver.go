package manu_order

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	manuOrderModel "eirc.app/internal/v1/structure/manu_order"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *manuOrderModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	manuOrder, err := r.ManuOrderService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	//成功的時候回傳 short name
	return code.GetCodeMessage(code.Successful, manuOrder.ManuOrderID)
}

func (r *resolver) List(input *manuOrderModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &manuOrderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, manu_orders, err := r.ManuOrderService.List(input)
	manu_ordersByte, err := json.Marshal(manu_orders)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(manu_ordersByte, &output.ManuOrders)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *manuOrderModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	manu_order, err := r.ManuOrderService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &manuOrderModel.Base{}
	manu_orderByte, _ := json.Marshal(manu_order)
	err = json.Unmarshal(manu_orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *manuOrderModel.Updated) interface{} {
	_, err := r.ManuOrderService.GetByID(&manuOrderModel.Field{ManuOrderID: &input.ManuOrderID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManuOrderService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *manuOrderModel.Updated) interface{} {
	manu_order, err := r.ManuOrderService.GetByID(&manuOrderModel.Field{ManuOrderID: &input.ManuOrderID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.ManuOrderService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	return code.GetCodeMessage(code.Successful, manu_order.ManuOrderID)
}
