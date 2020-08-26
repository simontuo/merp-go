package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"

	jwt2 "github.com/simontuo/merp-go/lib/jwt"
)

type AccessClams struct {
	UserPhone string
	jwt.StandardClaims
}

// 签发token
func SignToken(phone string) (token string, err error) {

	claims := AccessClams{UserPhone: phone}
	claims.ExpiresAt = jwt2.TokenExpiresAt()

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenObj.SignedString(jwt2.TokenKey())
}

func ParseToken(token string) (claims *AccessClams, err error) {
	claims = &AccessClams{}
	tokenObj, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwt2.TokenKey(), nil
	})

	// 返回结果
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token 错误")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New("token 过期")
			} else {
				return nil, errors.New("Cloud not handle this token " + err.Error())
			}
		} else {
			return nil, errors.New("无法解释token" + err.Error())
		}
	} else {
		if tokenObj != nil && tokenObj.Valid {
			return claims, nil
		} else {
			return nil, errors.New("无法解释token")
		}
	}
}