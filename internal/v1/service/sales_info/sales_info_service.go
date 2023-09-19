package sales_info

import (
	"encoding/json"
	"strconv"

	"eirc.app/internal/pkg/log"
	model "eirc.app/internal/v1/structure/sales_info"
	"github.com/shopspring/decimal"
)

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

	var totalCount float64
	for i, item := range output.GoodsDetail {
		//編上識別碼
		output.GoodsDetail[i].Identifier = field.GoodsDetail[i].SalesNo + "-" + strconv.Itoa(field.GoodsDetail[i].ItemSeq)
		//編上序號
		output.GoodsDetail[i].ItemNum = i + 1
		//總額計算
		if item.Quantity != 0 && item.Price != 0 {
			output.GoodsDetail[i].Subtotal, _ = decimal.NewFromFloat32(item.Quantity).Mul(decimal.NewFromFloat32(item.Price)).Round(2).Float64()
			totalCount, _ = decimal.NewFromFloat(totalCount).Add(decimal.NewFromFloat(output.GoodsDetail[i].Subtotal)).Round(2).Float64()
		}
	}
	output.Total = totalCount

	return output, nil
}

func (s *service) GetItemNameByIDandSEQ(no string, seq string) (name string, err error) {
	name, err = s.Entity.GetItemNameByIDandSEQ(no, seq)
	if err != nil {
		log.Error(err)

		return "", err
	}

	return name, nil
}
