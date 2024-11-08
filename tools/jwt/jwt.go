package jwt

import (
	"errors"
	_jwt "github.com/golang-jwt/jwt/v5"
	"gvb/internal/global"
	"gvb/internal/models/ctype"
	"time"
)

type Payload struct {
	Username string     `json:"username"`
	Nickname string     `json:"nickname"`
	UserID   uint       `json:"user_id"`
	Role     ctype.Role `json:"role"`
}
type CustomClaims struct {
	_jwt.RegisteredClaims
	Payload
}

var Secret []byte

func GenRefreshToken(payload Payload, exp time.Duration, jti string) (string, error) {
	rt, err := generateToken(payload, exp, jti)
	if err != nil {
		global.Log.Error(err.Error())
		return "", err
	}
	return rt, nil
}
func GenAccessToken(payload Payload, exp time.Duration) (string, error) {
	rt, err := generateToken(payload, exp, "")
	if err != nil {
		global.Log.Error(err.Error())
		return "", err
	}
	return rt, nil

}
func generateToken(payload Payload, exp time.Duration, jti string) (string, error) {
	Secret = []byte(global.Config.System.Jwt.Secret)
	claims := CustomClaims{
		Payload: payload,
		RegisteredClaims: _jwt.RegisteredClaims{
			ExpiresAt: _jwt.NewNumericDate(time.Now().Add(exp)),
			Issuer:    global.Config.System.Jwt.Issuer,
			ID:        jti,
		},
	}
	t := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
	return t.SignedString(Secret)
}
func ParseToken(token string) (*CustomClaims, error) {
	Secret = []byte(global.Config.System.Jwt.Secret)
	//如果想减少服务器操作量，可以在keyFunc中对token签名方法进行判断
	t, err := _jwt.ParseWithClaims(token, &CustomClaims{}, func(token *_jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		global.Log.Error(err.Error() + "解析token失败")
		return nil, err
	}
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("token无效")

}
