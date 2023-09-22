package account

import (
	"encoding/json"

	"errors"

	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	model "eirc.app/internal/v1/structure/accounts"
	"gorm.io/gorm"
)

func (s *service) Created(input *model.Created) (output *model.Base, err error) {

	key := "423CD5C09F7DD58950F1E494099EB075"

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	//將 JSON 字串處理成對應的結構
	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	pwd := util.HmacSha512(input.Pwd, key)
	log.Debug(input.Pwd)
	password, err := util.AesEncryptOFB([]byte(pwd), []byte(key))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//output.AccountId = util.GenerateUUID() //隨機產生key
	output.Pwd = util.Base64BydEncode(password)
	output.CreatedAt = util.NowToUTC()
	output.UpdatedAt = util.PointerTime(util.NowToUTC())
	output.UpdatedBy = &input.CreatedBy //沿用presenter 設定的使用者uuid
	output.IsDeleted = false
	output.Status = "Created"

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

	err = s.Entity.Created(table)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) List(input *model.Fields) (quantity int64, output []*model.Base, err error) {
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

// 帳號搜尋
func (s *service) GetByAccount(input *model.Field) (output *model.Base, err error) {
	field, err := s.Entity.GetByAccount(input)
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

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)

		return err
	}

	field.UpdatedBy = input.UpdatedBy
	field.UpdatedAt = util.PointerTime(util.NowToUTC())
	field.IsDeleted = true
	err = s.Entity.Updated(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {

	key := "423CD5C09F7DD58950F1E494099EB075"

	field, err := s.Entity.GetByID(&model.Field{AccountID: input.AccountID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)

		return err
	}

	//	密碼更新
	if input.Pwd != "" {
		pwd := util.HmacSha512(input.Pwd, key)
		log.Debug(input.Pwd)
		password, err := util.AesEncryptOFB([]byte(pwd), []byte(key))

		if err != nil {
			log.Error(err)
			return err
		}
		input.Pwd = util.Base64BydEncode(password)
		input.PwdUpdatedAt = util.PointerTime(util.NowToUTC())
	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)

		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)

		return err
	}

	err = s.Entity.Updated(field)

	return err
}

func (s *service) AcknowledgeAccount(input *model.Field) (acknowledge bool, output *model.Base, err error) {

	key := "423CD5C09F7DD58950F1E494099EB075"

	//查詢資料庫資料
	field, err := s.Entity.GetByAccount(input)
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, errors.New("Account error")
		} else {
			return false, nil, err
		}
	}

	password, err := util.AesDecryptOFB(util.Base64StdDecode(field.Pwd), []byte(key))
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	input.Pwd = util.HmacSha512(input.Pwd, key)
	if string(password) != input.Pwd {
		return false, nil, errors.New("Incorrect password")
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return false, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return false, output, err
	}

	return true, output, nil
}
