package file

import (
	"encoding/json"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/file"
)

func (s *service) Created(input *model.Created) (output *model.Base, err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	input.CreatedAt = util.NowToUTC()
	input.IsDeleted = false

	marshal, err = json.Marshal(output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	table := &model.Table{}
	err = json.Unmarshal(marshal, &table)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fileId, err := s.Entity.Created(table)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	output.FileID = fileId
	return output, nil
}

func (s *service) List(input *model.Fields) (quantity int64, output []*model.Base, err error) {

	if input.IsDeleted == nil {
		//如果未指定檔案刪除狀態
		input.IsDeleted = util.PointerBool(false)
	}

	if input.Limit == 0 && input.Page == 0 {
		input.Limit = 1
		input.Page = 1
	}

	amount, fields, err := s.Entity.List(input)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	return amount, output, err
}

func (s *service) GetByID(input *model.Field) (output *model.Base, err error) {
	field, err := s.Entity.GetByID(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) Deleted(input *model.Field) (err error) {
	field, err := s.Entity.GetByID(&model.Field{FileID: input.FileID})
	if err != nil {
		log.Error(err)
		return err
	}

	field.IsDeleted = true
	err = s.Entity.Updated(field)

	return err
}

func (s *service) GetAllManuOrder() (output []*model.Base, err error) {
	field, err := s.Entity.GetAllManuOrder()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}
