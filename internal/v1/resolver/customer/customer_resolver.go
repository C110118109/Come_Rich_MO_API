package customer

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	customerModel "eirc.app/internal/v1/structure/customer"
	"gorm.io/gorm"
)

func (r *resolver) List(input *customerModel.Fields) interface{} {

	output := &customerModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, customers, err := r.CustomerService.List(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	customersByte, err := json.Marshal(customers)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(customersByte, &output.Customers)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *customerModel.Field) interface{} {

	base, err := r.CustomerService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	frontCustomer := &customerModel.Base{}
	customerByte, _ := json.Marshal(base)
	err = json.Unmarshal(customerByte, &frontCustomer)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, frontCustomer)
}
