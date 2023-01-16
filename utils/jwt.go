package utils

import (
	"SH-admin/global"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Uid int64 `json:"uid"`
	jwt.RegisteredClaims
}

// CreateClaims 創建claims
func CreateClaims(uid int64) MyClaims {
	c := MyClaims{
		uid,
		jwt.RegisteredClaims{
			Issuer: "sh-admin",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Duration(global.Config.Jwt.ExpiredTime) * time.Second),
			},
		},
	}
	return c
}

// CreateToken 透過claims產生token
func CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.Jwt.SigningKey))
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.SigningKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
