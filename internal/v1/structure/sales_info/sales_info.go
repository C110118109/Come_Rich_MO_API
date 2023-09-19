package sales_info

import (
	"time"

	model "eirc.app/internal/v1/structure"
	customer "eirc.app/internal/v1/structure/customer"
	goods_detail "eirc.app/internal/v1/structure/goods_detail"
)

type Table struct {
	// 單據號碼
	SalesNo string `gorm:"primaryKey;column:SP_NO;type:TEXT" json:"sales_no,omitempty"`
	// 單據類型
	OrderType string `gorm:"column:SP_SLIP_FG;type:TEXT" json:"order_type,omitempty"`
	// 客戶名稱
	CustomerName string `gorm:"column:SP_CTNAME;type:TEXT" json:"customer_name,omitempty"`
	// 客戶代號
	CustomerNo string `gorm:"column:SP_CTNO;type:TEXT" json:"customer_no,omitempty"`
	// 送貨地址
	DeliveryAddress string `gorm:"column:SP_CTADD2;type:TEXT" json:"delivery_address,omitempty"`
	// 單據備註
	SalesRemark string `gorm:"column:SP_REM;type:TEXT" json:"sales_remark,omitempty"`
	// 單據日期
	SalesDate time.Time `gorm:"column:SP_DATE;type:DATETIME" json:"sales_date,omitempty"`
	// 發票號碼
	InvoiceNumber string `gorm:"column:SP_INVOICE;type:TEXT" json:"invoice_number,omitempty"`

	// 客戶資料
	CustomerInfo customer.Table `gorm:"foreignkey:CT_NO;references:SP_CTNO"`

	// 商品詳細
	GoodsDetail []goods_detail.Table `gorm:"foreignkey:SD_NO;references:SP_NO"`
}

type Base struct {

	// 客戶名稱
	CustomerName string `json:"customer_name"`
	// 客戶代號
	CustomerNo string `json:"customer_no"`
	// 單據日期
	SalesDate time.Time `json:"sales_date"`
	// 單據號碼
	SalesNo string `json:"sales_no"`
	// 發票號碼
	InvoiceNumber string `json:"invoice_number"`
	// 送貨地址
	DeliveryAddress string `json:"delivery_address"`
	// 合計金(客製)
	Total float64 `json:"total,omitempty"`
	// 單據備註
	SalesRemark string `json:"sales_remark"`

	// 客戶資料(Preload)
	CustomerInfo customer.Simple
	// 商品詳細(Preload)
	GoodsDetail []goods_detail.Base
	// 是否存在製令單
	HasManuOrder bool `json:"has_manu_order,omitempty"`
}

type Field struct {
	// 單據號碼
	SalesNo *string `json:"sales_no,omitempty"   form:"sales_no"`
	// 客戶名稱
	CustomerName *string `json:"customer_name,omitempty"`
	// 客戶代號
	CustomerNo *string `json:"customer_no,omitempty"  form:"customer_no"`
	// 是否存在製令單
	HasManuOrder bool `json:"has_manu_order"`

	// 搜尋產品名稱
	ProductName *string `json:"product_name,omitempty" form:"product_name"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	SalesInfos []*struct {
		Field
	} `json:"sales_infos"`
	model.OutPage
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "SSLIP"
}
