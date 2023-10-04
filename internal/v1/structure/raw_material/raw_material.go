package raw_material

import (
	model "eirc.app/internal/v1/structure"
)

type Table struct {
	RawMaterialID string `gorm:"column:raw_material_id;type:text;primary_key" json:"raw_material_id,omitempty"`
	MaterialName  string `gorm:"column:material_name;type:text" json:"material_name,omitempty"`
	Quantity      string `gorm:"column:quantity;type:text" json:"quantity,omitempty"`
	Percentage    string `gorm:"column:percentage;type:text" json:"percentage,omitempty"`
	Remark        string `gorm:"column:remark;type:text" json:"remark,omitempty"`
	ManuOrderID   string `gorm:"column:manu_order_id;type:text" json:"manu_order_id,omitempty"`
}

// 結構基底(SHOW/COPY)
type Base struct {
	RawMaterialID string `json:"raw_material_id"`
	MaterialName  string `json:"material_name"`
	Quantity      string `json:"quantity"`
	Percentage    string `json:"percentage"`
	Remark        string `json:"remark"`
}

// 清單顯示&查詢用
type Field struct {
	RawMaterialID string `json:"raw_material_id,omitempty"`
	MaterialName  string `json:"material_name,omitempty"`
	Quantity      string `json:"quantity,omitempty"`
	Percentage    string `json:"percentage,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	RawMaterials []*struct {
		Field
	} `json:"raw_materials"`
	model.OutPage
}

// 新增版本
type Created struct {
	RawMaterialID string `json:"raw_material_id,omitempty"`
	MaterialName  string `json:"material_name,omitempty"`
	Quantity      string `json:"quantity,omitempty"`
	Percentage    string `json:"percentage,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

type Updated struct {
	RawMaterialID string `json:"raw_material_id,omitempty"`
	MaterialName  string `json:"material_name,omitempty"`
	Quantity      string `json:"quantity,omitempty"`
	Percentage    string `json:"percentage,omitempty"`
	Remark        string `json:"remark,omitempty"`
}

func (a *Table) TableName() string {
	return "raw_material"
}
