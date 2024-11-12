package srv_user

import (
	"github.com/google/uuid"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/service/srv_redis"
	"gvb/tools/jwt"
	"time"
)

func (s UserSrv) GenJwt(user *dao.UserModel) (refreshToken *string, accessToken *string, err error) {
	payload := jwt.Payload{
		UserID:   user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	jti := uuid.NewString()
	rtExp := time.Duration(global.Config.System.Jwt.RTExpires) * time.Second
	atExp := time.Duration(global.Config.System.Jwt.ATExpires) * time.Second
	rt, err := jwt.GenRefreshToken(payload, rtExp, jti)
	if err != nil {
		global.Log.Error(err)
		return nil, nil, err

	}
	at, err := jwt.GenAccessToken(payload, atExp)
	if err != nil {
		global.Log.Error(err)
		return nil, nil, err
	}
	err = srv_redis.SetToken(user.ID, jti, rtExp)
	if err != nil {
		global.Log.Error(err)
		return nil, nil, err
	}
	s.C.SetCookie(
		"refreshToken",       // Cookie 名称
		rt,                   // Cookie 值
		int(rtExp.Seconds()), // 过期时间，单位为秒
		"/",                  // 路径
		"localhost",          // 域名
		true,                 // 是否只能通过 HTTPS 传输
		true,                 // HttpOnly，禁止 JavaScript 访问
	)
	return &rt, &at, nil
}
