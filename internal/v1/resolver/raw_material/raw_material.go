package raw_material

import (
	"encoding/json"
	"errors"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	rawMaterialModel "eirc.app/internal/v1/structure/raw_material"
	"gorm.io/gorm"
)

func (r *resolver) Created(trx *gorm.DB, input *rawMaterialModel.Created) interface{} {
	defer trx.Rollback()
	// Todo 檢查重複
	rawMaterial, err := r.RawMaterialService.WithTrx(trx).Created(input)
	if err != nil {
		log.Error(err)

		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	//成功的時候回傳 short name
	return code.GetCodeMessage(code.Successful, rawMaterial.RawMaterialID)
}

func (r *resolver) List(input *rawMaterialModel.Fields) interface{} {
	output := &rawMaterialModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, raw_materials, err := r.RawMaterialService.List(input)
	raw_materialsByte, err := json.Marshal(raw_materials)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(raw_materialsByte, &output.RawMaterials)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) GetByID(input *rawMaterialModel.Field) interface{} {
	raw_material, err := r.RawMaterialService.GetByID(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &rawMaterialModel.Base{}
	raw_materialByte, _ := json.Marshal(raw_material)
	err = json.Unmarshal(raw_materialByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *resolver) Deleted(input *rawMaterialModel.Updated) interface{} {
	// _, err := r.RawMaterialService.GetByID(&rawMaterialModel.Field{RawMaterialID: &input.RawMaterialID,
	// 	IsDeleted: util.PointerBool(false)}	)
	_, err := r.RawMaterialService.GetByID(&rawMaterialModel.Field{RawMaterialID: &input.RawMaterialID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RawMaterialService.Deleted(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (r *resolver) Updated(input *rawMaterialModel.Updated) interface{} {
	raw_material, err := r.RawMaterialService.GetByID(&rawMaterialModel.Field{RawMaterialID: &input.RawMaterialID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = r.RawMaterialService.Updated(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}
	return code.GetCodeMessage(code.Successful, raw_material.RawMaterialID)
}
