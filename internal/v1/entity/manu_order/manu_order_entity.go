package manu_order

import (
	model "eirc.app/internal/v1/structure/manu_order"
)

func (e *entity) Created(input *model.Table) (err error) {

	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{}).Preload("RawMaterial")
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", *input.IsDeleted)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Preload("RawMaterial").Where("manu_order_id = ?", input.ManuOrderID)
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByAccount(input *model.Field) (output *model.Table, err error) {

	db := e.db.Model(&model.Table{}).Where("manu_order = ?", input.ManuOrderID)
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("manu_order_id = ?", input.ManuOrderID).Save(&input).Error

	return err
}
