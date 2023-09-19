package file

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// 檔案編號
	FileID string `gorm:"column:file_id;type:TEXT;primary_key" json:"file_id,omitempty"`
	// 單據項目識別碼
	Identifier string `gorm:"column:identifier;type:TEXT" json:"identifier,omitempty"`
	// 路徑
	PathKey string `gorm:"column:path_key;type:TEXT" json:"path_key,omitempty"`
	//副檔名
	Extension string `gorm:"column:extension;type:TEXT" json:"extension,omitempty"`
	//檔案大小
	Size string `gorm:"column:size;type:TEXT" json:"size,omitempty"`
	//對應單據
	SalesNo string `gorm:"column:sales_no;type:TEXT" json:"sales_no,omitempty"`

	// 創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:boolean;default:false" json:"is_deleted,omitempty"`
}

type Base struct {
	// 檔案編號
	FileID string `json:"file_id"`
	// 單據項目識別碼
	Identifier string `json:"identifier"`
	// 路徑
	PathKey string `json:"path_key"`
	//副檔名
	Extension string `json:"extension"`
	//檔案大小
	Size string `json:"size"`
	//對應單據
	SalesNo string `json:"sales_no"`

	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`

	// 品名(客製)
	ItemName string `json:"item_name,omitempty"`
}

type Created struct {
	// 檔案編號
	FileID string `json:"file_id,omitempty"`
	// 單據項目識別碼
	Identifier string `json:"identifier,omitempty"`
	// 路徑
	PathKey string `json:"path_key,omitempty"`
	//副檔名
	Extension string `json:"extension,omitempty"`
	//檔案大小
	Size string `json:"size,omitempty"`
	//對應單據
	SalesNo string `json:"sales_no,omitempty"`

	// 創建時間
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`

	//前端讀入的base64 格式
	Base64 string `json:"base64,omitempty" gorm:"-"`
}

type Field struct {
	// 檔案編號
	FileID *string `json:"file_id"`
	// 單據項目識別碼
	Identifier *string `json:"identifier"  form:"identifier"`
	// 路徑
	PathKey *string `json:"path_key"`
	//副檔名
	Extension *string `json:"extension"`
	//檔案大小
	Size *string `json:"size"`
	// 單據號碼
	SalesNo *string `json:"sales_no"  form:"sales_no"`

	// 創建時間
	CreatedAt *time.Time `json:"created_at"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty"  form:"del"`

	// 品名(客製)
	ItemName *string `json:"item_name,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Files []*struct {
		Field
	} `json:"files"`
	model.OutPage
}

func (a *Table) TableName() string {
	return "file"
}
