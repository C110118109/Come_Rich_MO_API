package accounts

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

type Table struct {
	// uuid員工序號
	AccountId string `gorm:"primaryKey;column:account_id;type:TEXT;" json:"account_id,omitempty"`
	// 員工編號
	CompanyId string `gorm:"column:company_id;type:VARCHAR;" json:"company_id,omitempty"`
	// 帳號
	Account string `gorm:"column:account;type:VARCHAR;" json:"account,omitempty"`
	// 姓名
	Name string `gorm:"column:name;type:VARCHAR;" json:"name,omitempty"`
	// 英文姓名
	EngName string `gorm:"column:eng_name;type:VARCHAR;" json:"eng_name,omitempty"`
	// 密碼
	Pwd string `gorm:"column:pwd;type:VARCHAR;" json:"password,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;default:false" json:"is_deleted,omitempty"`
	// 手機
	Phone string `gorm:"column:phone;type:TEXT;" json:"phone,omitempty"`
	// 信箱
	Email string `gorm:"column:email;type:TEXT;" json:"email,omitempty"`
	// 狀態
	Status string `gorm:"column:status;type:VARCHAR" json:"status,omitempty"`
	// 權限名稱
	RoleName string `gorm:"column:role_name;type:TEXT;" json:"role_name,omitempty"`
	// 密碼更新時間
	PwdUpdatedAt *time.Time `gorm:"column:pwd_updated_at;type:TIMESTAMP;" json:"pwd_updated_at,omitempty"`
	// 建立時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:TEXT;" json:"created_by,omitempty"`
	// 資料更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 資料更新人
	UpdatedBy *string `gorm:"column:updated_by;type:TEXT;" json:"updated_by,omitempty"`

}

// 結構基底(SHOW/COPY)
type Base struct {
	// uuid員工序號
	AccountId string `json:"account_id"`
	// 員工編號
	CompanyId string `json:"company_id"`
	// 帳號
	Account string `json:"account"`
	// 姓名
	Name string `json:"name"`
	// 英文姓名
	EngName string `json:"eng_name"`
	// 密碼
	Pwd string `json:"password,omitempty"`
	// 手機
	Phone string `json:"phone"`
	// 信箱
	Email string `json:"email"`
	// 狀態
	Status string `json:"status"`
	// 權限名稱
	RoleName string `json:"role_name"`
	// 密碼更新時間
	PwdUpdatedAt *time.Time `json:"pwd_updated_at"`
	// 建立時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by"`
	// 資料更新時間
	UpdatedAt *time.Time `json:"updated_at"`
	// 資料更新人
	UpdatedBy *string `json:"updated_by"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`

}

// 清單顯示&查詢用
type Field struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 員工編號
	CompanyID *string `json:"company_id,omitempty" form:"company_id"`
	// 帳號
	Account *string `json:"account,omitempty" form:"account"`
	// 密碼
	Pwd string `json:"password,omitempty"`
	// 中文名稱
	Name *string `json:"name,omitempty"   form:"name"`
	// 英文姓名
	EngName *string `json:"eng_name,omitempty" form:"eng_name"`
	// 狀態
	Status *string `json:"status,omitempty" swaggerignore:"true"`
	// 權限名稱
	RoleName *string `json:"role_name,omitempty"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"   form:"del"`


	// 密碼更新時間
	PwdUpdatedAt *time.Time `json:"pwd_updated_at,omitempty"`
	// 資料更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Accounts []*struct {
		Field
	} `json:"accounts"`
	model.OutPage
}

// 新增版本
type Created struct {
	// 員工編號
	CompanyId string `json:"company_id,omitempty"`
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 姓名
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 英文姓名
	EngName string `json:"eng_name,omitempty"`
	// 密碼
	Pwd string `json:"password,omitempty" binding:"required" validate:"required"`
	// 手機
	Phone string `json:"phone,omitempty"`
	// 信箱
	Email string `json:"email,omitempty"`
	// 權限名稱
	RoleName string `json:"role_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

type Updated struct {
	// 編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 帳號
	Account string `json:"account,omitempty"`
	// 員工編號
	CompanyID string `json:"company_id,omitempty"`
	// 中文名稱
	Name string `json:"name,omitempty"`
	// 英文姓名
	EngName string `json:"eng_name,omitempty"`
	// 密碼
	Pwd string `json:"password,omitempty"`
	// 手機
	Phone string `json:"phone,omitempty"`
	// 信箱
	Email string `json:"email,omitempty"`
	// 狀態
	Status string `json:"status,omitempty" swaggerignore:"true"`
	// 權限名稱
	RoleName string `json:"role_name,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" swaggerignore:"true"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty" swaggerignore:"true"`
	// 密碼更新時間
	PwdUpdatedAt *time.Time `json:"pwd_updated_at,omitempty"`
}

type Login struct {
	// 帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password" binding:"required" validate:"required"`
}

type Token struct {
}

func (a *Table) TableName() string {
	return "accounts"
}