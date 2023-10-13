package raw_material

import (
	model "eirc.app/internal/v1/structure/raw_material"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("raw_material_id asc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{})
	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetByAccount(input *model.Field) (output *model.Table, err error) {

	db := e.db.Model(&model.Table{}).Where("raw_material = ?", input.RawMaterialID)

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("raw_material_id = ?", input.RawMaterialID).Save(&input).Error

	return err
}
