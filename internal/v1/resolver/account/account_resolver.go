package account

import (
	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	accountModel "eirc.app/internal/v1/structure/accounts"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *accountModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 角色名稱

	//查值會丟入物件，但實際又會在Entity針對input拆解出來一次
	_, err := r.AccountService.WithTrx(trx).GetByAccount(&accountModel.Field{Account: &input.Account,
		IsDeleted: util.PointerBool(false)})

	//檢查帳號是否重複
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//帳號不存在
			account, err := r.AccountService.WithTrx(trx).Created(input)
			if err != nil {
				log.Error(err)
				return code.GetCodeMessage(code.InternalServerError, err.Error())
			}

			trx.Commit()
			return code.GetCodeMessage(code.Successful, account.Account)

		} else {
			//其他錯誤
			log.Error(err)
			return code.GetCodeMessage(code.InternalServerError, err.Error())
		}

	} else {
		//帳號重複
		return code.GetCodeMessage(code.ItemExisted, err)
	}

}

func (r *resolver) List(input *accountModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &accountModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accounts, err := r.AccountService.List(input)
	accountsByte, err := json.Marshal(accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountsByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *accountModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	account, err := r.AccountService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &accountModel.Base{}
	accountByte, _ := json.Marshal(account)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *accountModel.Updated) interface{} {
	_, err := r.AccountService.GetByID(&accountModel.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.AccountService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *accountModel.Updated) interface{} {
	account, err := r.AccountService.GetByID(&accountModel.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.AccountService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	return code.GetCodeMessage(code.Successful, account.AccountId)
}
