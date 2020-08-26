package user

import "github.com/go-playground/validator/v10"

type ResetPasswordRequest struct {
	Phone           string `form:"phone" json:"phone" binding:"required"`
	VerifyCode      string `json:"verify_code" form:"verify_code" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required,ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

func ConfirmPassword(fl validator.FieldLevel) bool {
	var req *ResetPasswordRequest = fl.Parent().Interface().(*ResetPasswordRequest)
	if req.Password != req.ConfirmPassword {
		return false
	}

	return true
}
