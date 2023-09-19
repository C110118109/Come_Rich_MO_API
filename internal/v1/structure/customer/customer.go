package customer

import (
	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 客戶號碼
	No string `gorm:"primaryKey;column:CT_NO;type:TEXT" json:"no,omitempty"`
	// 客戶全名
	Name string `gorm:"column:CT_NAME;type:TEXT" json:"name,omitempty"`
	// 客戶縮寫名字
	SName string `gorm:"column:CT_SNAME;type:TEXT" json:"sname,omitempty"`
	// 客戶地址1
	Address1 string `gorm:"column:CT_ADDR1;type:TEXT" json:"address1,omitempty"`
	// 客戶地址2
	Address2 string `gorm:"column:CT_ADDR2;type:TEXT" json:"address2,omitempty"`
	// 客戶地址3
	Address3 string `gorm:"column:CT_ADDR3;type:TEXT" json:"address3,omitempty"`
	// 客戶電話
	Tel string `gorm:"column:CT_TEL;type:TEXT" json:"tel,omitempty"`
	// 客戶傳真
	Fax string `gorm:"column:CT_FAX;type:TEXT" json:"fax,omitempty"`
	// 客戶統編
	Unino string `gorm:"column:CT_UNINO;type:TEXT" json:"unino,omitempty"`
	// 客戶備註
	Remark string `gorm:"column:CT_REMARK;type:TEXT" json:"remark,omitempty"`
}

type Base struct {
	// 客戶號碼
	No string `json:"no"`
	// 客戶全名
	Name string `json:"name"`
	// 客戶縮寫名字
	SName string `json:"sname"`
	// 客戶地址1
	Address1 string `json:"address1"`
	// 客戶地址2
	Address2 string `json:"address2"`
	// 客戶地址3
	Address3 string `json:"address3"`
	// 客戶電話
	Tel string `json:"tel"`
	// 客戶傳真
	Fax string `json:"fax"`
	// 客戶統編
	Unino string `json:"unino"`
	// 客戶備註
	Remark string `json:"remark"`
}

type Simple struct {
	// 客戶號碼
	No string `json:"no,omitempty"`
	// 客戶全名
	Name string `json:"name,omitempty"`
	// 客戶電話
	Tel string `json:"tel,omitempty"`
	// 客戶傳真
	Fax string `json:"fax,omitempty"`
	// 客戶統編
	Unino string `json:"unino,omitempty"`
}

type Field struct {
	// 客戶號碼
	No *string `json:"no"  form:"no"`
	// 客戶全名
	SName *string `json:"sname"  form:"sname"`
	// 客戶電話
	Tel *string `json:"tel"  form:"tel"`
	// 客戶統編
	Unino *string `json:"unino"  form:"unino"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Customers []*struct {
		Field
	} `json:"customers"`
	model.OutPage
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "PCUST"
}
