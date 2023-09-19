package customer

import (
	model "eirc.app/internal/v1/structure/customer"
)

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	//供搜尋
	if input.No != nil {
		db.Where("CT_NO = ?", "%"+*input.No+"%")
	}

	if input.SName != nil {
		db.Where("CT_SNAME like ?", "%"+*input.SName+"%")
	}

	if input.Tel != nil {
		db.Where("CT_TEL like ?", "%"+*input.Tel+"%")
	}

	if input.Unino != nil {
		db.Where("CT_UNINO = ?", "%"+*input.Unino+"%")
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("CT_NO asc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("CT_NO = ?", input.No)
	err = db.First(&output).Error

	return output, err
}
