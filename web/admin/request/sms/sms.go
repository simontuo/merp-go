package sms

type LoginVerifyCodeRequest struct {
	Phone string `json:"phone" form:"phone"`
}

type ResetPasswordVerifyCodeRequest struct {
	Phone string `json:"phone" form:"phone"`
}
