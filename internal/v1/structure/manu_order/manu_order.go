package manu_order

import (
	"time"

	model "eirc.app/internal/v1/structure"
	rawMaterial "eirc.app/internal/v1/structure/raw_material"
)

type Table struct {
	//製令流水號
	ManuOrderID string `gorm:"column:manu_order_id;type:text;primary_key" json:"manu_order_id,omitempty"`
	//
	GoodsName      string               `gorm:"column:goods_name;type:text" json:"goods_name,omitempty"`
	Identifier     string               `gorm:"column:identifier;type:text" json:"identifier,omitempty"`
	TotalQuantity  string               `gorm:"column:total_quantity;type:text" json:"total_quantity,omitempty"`
	CompletionDate string               `gorm:"column:completion_date;type:datetime" json:"completion_date,omitempty"`
	DateOfIssuance string               `gorm:"column:date_of_issuance;type:datetime" json:"date_of_issuance,omitempty"`
	CreatedAt      time.Time            `gorm:"column:created_at;type:timestamp" json:"created_at,omitempty"`
	UpdatedAt      time.Time            `gorm:"column:updated_at;type:timestamp" json:"updated_at,omitempty"`
	IsDeleted      bool                 `gorm:"column:is_deleted;type:bool" json:"is_deleted,omitempty"`
	RawMaterial    []*rawMaterial.Table `gorm:"foreignkey:manu_order_id;references:manu_order_id"`
}

// 結構基底(SHOW/COPY)
type Base struct {
	//製令流水號
	ManuOrderID    string    `json:"manu_order_id"`
	GoodsName      string    `json:"goods_name"`
	Identifier     string    `json:"identifier"`
	TotalQuantity  string    `json:"total_quantity"`
	CompletionDate string    `json:"completion_date"`
	DateOfIssuance string    `json:"date_of_issuance"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
	RawMaterial    []*rawMaterial.Base
}

// 清單顯示&查詢用
type Field struct {
	//製令流水號
	ManuOrderID    *string   `json:"manu_order_id,omitempty"`
	GoodsName      string    `json:"goods_name,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	TotalQuantity  string    `json:"total_quantity,omitempty"`
	CompletionDate string    `json:"completion_date,omitempty"`
	DateOfIssuance string    `json:"date_of_issuance,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	IsDeleted      *bool     `json:"is_deleted,omitempty"`
	RawMaterial    []*rawMaterial.Base
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	ManuOrders []*struct {
		Field
	} `json:"manu_orders"`
	model.OutPage
}

// 新增版本
type Created struct {
	//製令流水號
	ManuOrderID    string    `json:"manu_order_id,omitempty"`
	GoodsName      string    `json:"goods_name,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	TotalQuantity  string    `json:"total_quantity,omitempty"`
	CompletionDate string    `json:"completion_date,omitempty"`
	DateOfIssuance string    `json:"date_of_issuance,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	IsDeleted      bool      `json:"is_deleted,omitempty"`
}

type Updated struct {
	//製令流水號
	ManuOrderID    string    `json:"manu_order_id,omitempty"`
	GoodsName      string    `json:"goods_name,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	TotalQuantity  string    `json:"total_quantity,omitempty"`
	CompletionDate string    `json:"completion_date,omitempty"`
	DateOfIssuance string    `json:"date_of_issuance,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	IsDeleted      bool      `json:"is_deleted,omitempty"`
}

func (a *Table) TableName() string {
	return "manu_order"
}
