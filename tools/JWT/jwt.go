package JWT

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"gvb/internal/global"
	"time"
)

type Payload struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	UserID   uint   `json:"user_id"`
	Role     int    `json:"role"`
}
type CustomClaims struct {
	jwt.RegisteredClaims
	Payload
}

var Secret []byte

func GenerateToken(payload Payload) (string, error) {
	Secret = []byte(global.Config.System.Jwt.Secret)
	claims := CustomClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.Config.System.Jwt.Expires))),
			Issuer:    global.Config.System.Jwt.Issuer,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(Secret)
}
func ParseToken(token string) (*CustomClaims, error) {
	Secret = []byte(global.Config.System.Jwt.Secret)
	//如果想减少服务器操作量，可以在keyFunc中对token签名方法进行判断
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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
