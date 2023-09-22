package login

type Login struct {
	//帳號
	Account string `json:"account,omitempty" binding:"required" validate:"required"`
	//密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
}
