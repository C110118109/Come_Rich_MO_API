package sales_info

import (
	goodsDetailModel "eirc.app/internal/v1/structure/goods_detail"
	model "eirc.app/internal/v1/structure/sales_info"
	//"fmt"
	"gorm.io/gorm"
)

/* SP_SLIP_FG 控制單據類型
1 = 進貨單 ; 2 = 進貨退回單 ; 3 = 銷貨單
*/

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {

	db := e.db.Model(&model.Table{})
	//供搜尋
	if input.SalesNo != nil && *input.SalesNo != "" {
		db.Where("SP_NO like ?", "%"+*input.SalesNo+"%")
	}

	if input.CustomerNo != nil && *input.CustomerNo != "" {
		db.Where("SP_CTNO like ? or SP_CTNAME like ?", "%"+*input.CustomerNo+"%", "%"+*input.CustomerNo+"%")
	}

	if input.ProductName != nil && *input.ProductName != "" {
		//品項名稱搜尋, 找符合的單據編號
		subQuery := e.db.Model(&goodsDetailModel.Table{}).Where("SD_NAME like ?", "%"+*input.ProductName+"%").Select("SD_NO")
		db.Where("SP_NO IN (?)", subQuery)
	}

	err = db.Where("SP_SLIP_FG = ?", "2").Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("SP_NO desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).
		Preload("CustomerInfo").
		Preload("GoodsDetail", func(db *gorm.DB) *gorm.DB {
			return db.Where("SSLPDT.SD_SLIP_FG = ?", 2).Order("SSLPDT.SD_SEQ ASC")
		}).
		Where("SP_NO = ?", input.SalesNo).Where("SP_SLIP_FG = ?", "2")
	err = db.First(&output).Error

	return output, err
}

func (e *entity) GetItemNameByIDandSEQ(no string, seq string) (name string, err error) {
	var goodsDetail goodsDetailModel.Table
	err = e.db.Model(&goodsDetailModel.Table{}).Where("SD_NO = ? and SD_SEQ = ?", no, seq).First(&goodsDetail).Error
	return goodsDetail.Name, err
}
