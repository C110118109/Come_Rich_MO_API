package jwe

type JWE struct {
	//中文名稱
	Name string `json:"name,omitempty"`
	//使用者UUID
	AccountID string `json:"account_id,omitempty"`
	//帳號
	Account string `json:"account,omitempty"`
	//員工編號
	CompanyID string `json:"company_id,omitempty"`
	//角色
	Role string `json:"role,omitempty"`
}

type Token struct {
	//授權令牌
	AccessToken string `json:"access_token,omitempty"`
	//刷新令牌
	RefreshToken string `json:"refresh_token,omitempty"`
	//使用者UUID
	AccountID string `json:"account_id,omitempty"`
	//使用者名稱
	Name string `json:"name,omitempty"`
	//使用者名稱(英文)
	EngName string `json:"eng_name,omitempty"`
	//角色
	Role string `json:"role,omitempty"`
	//部門
	DepName string `json:"dep_name,omitempty"`
}

type Refresh struct {
	//刷新令牌
	RefreshToken string `json:"refresh_token,omitempty" binding:"required" validate:"required"`
}
