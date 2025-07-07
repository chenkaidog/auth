package dto

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required,max=64"`
	Password string `form:"password" json:"password" binding:"required,max=128"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

type LogoutReq struct {
}

type LogoutResp struct {
}

type PasswordUpdateReq struct {
	Password    string `form:"password" json:"password" binding:"required,max=128"`
	PasswordNew string `form:"password_new" json:"password_new" binding:"required,min=8,max=128"`
}

type PasswordUpdateResp struct {
}
