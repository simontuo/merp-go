package token

import "github.com/go-playground/validator/v10"

type SignRequest struct {
	Phone      string `form:"phone" json:"phone" binding:"required"`
	Password   string `form:"password" json:"password" binding:"SignPassword"`
	VerifyType string `json:"verify_type" form:"verify_type" binding:"required"`
	VerifyCode string `json:"verify_code" form:"verify_code" binding:"SignVerifyCode"`
}

func SignPassword(fl validator.FieldLevel) bool {
	var req *SignRequest = fl.Parent().Interface().(*SignRequest)
	if req.VerifyType == "password" && req.Password == "" {
		return false
	}

	return true
}

func SignVerifyCode(fl validator.FieldLevel) bool {
	var req *SignRequest = fl.Parent().Interface().(*SignRequest)
	if req.VerifyType == "code" && req.VerifyCode == "" {
		return false
	}

	return true
}
