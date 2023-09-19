package file

import (
	model "eirc.app/internal/v1/structure/file"
)

func (e *entity) Created(input *model.Table) (fileId string, err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return input.FileID, err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	if input.Identifier != nil {
		db.Where("identifier like ?", *input.Identifier)
	}

	if input.SalesNo != nil {
		db.Where("sales_no like ?", *input.SalesNo)
	}

	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", *input.IsDeleted)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{})
	db.Where("file_id = ?", input.FileID)

	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", *input.IsDeleted)
	}
	err = db.First(&output).Error

	return output, err
}

func (e *entity) Updated(input *model.Table) (err error) {
	db := e.db.Model(&model.Table{})
	db.Where("file_id = ?", input.FileID)

	err = db.Save(&input).Error

	return err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).
		Where("file_id = ?", input.FileID).
		Delete(&input).Error

	return err
}

func (e *entity) GetAllManuOrder() (output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	db.Where("is_deleted = ?", false)
	err = db.Distinct("sales_no").Order("created_at desc").Find(&output).Error

	return output, err
}
