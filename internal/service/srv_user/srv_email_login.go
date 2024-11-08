package srv_user

import (
	"errors"
	"github.com/google/uuid"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_redis"
	"gvb/tools/encryptor"
	"gvb/tools/jwt"
	"time"
)

func (s UserSrv) EmailLogin(user *req.UserEmailLoginReq) (*res.Response, error) {
	userModel, exist, err := dao.FindWithUsername(dao.UserModel{}, user.Username)
	if !exist {
		resp := res.NewResponse(res.UserNotExist, res.EmptyData, res.CodeMsg(res.UserNotExist))
		return resp, errors.New(res.CodeMsg(res.UserNotExist))
	}
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	if !encryptor.CheckMd5([]byte(user.Password), userModel.Password) {
		resp := res.NewResponse(res.PasswordNotMatched, res.EmptyData, res.CodeMsg(res.PasswordNotMatched))
		return resp, err
	}
	//jwt
	payload := jwt.Payload{
		UserID:   userModel.ID,
		Username: userModel.Username,
		Nickname: userModel.Nickname,
		Role:     userModel.Role,
	}
	jti := uuid.NewString()
	rtExp := time.Duration(global.Config.System.Jwt.RTExpires) * time.Second
	atExp := time.Duration(global.Config.System.Jwt.ATExpires) * time.Second
	rt, err := jwt.GenRefreshToken(payload, rtExp, jti)
	if err != nil {
		global.Log.Error(err)
		resp := res.NewResponse(res.TokenGenerateFailed, res.EmptyData, res.CodeMsg(res.TokenGenerateFailed))
		return resp, err
	}
	at, err := jwt.GenAccessToken(payload, atExp)
	if err != nil {
		global.Log.Error(err)
		resp := res.NewResponse(res.TokenGenerateFailed, res.EmptyData, res.CodeMsg(res.TokenGenerateFailed))
		return resp, err
	}
	err = srv_redis.SetToken(userModel.ID, jti, rtExp)
	if err != nil {
		global.Log.Error(err)
		resp := res.NewResponse(res.RedisSetFailed, res.EmptyData, res.CodeMsg(res.RedisSetFailed))
		return resp, err
	}
	dataResp := res.LoginRes{
		AccessToken:  at,
		RefreshToken: rt,
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

	resp := res.NewResponse(res.Success, dataResp, res.CodeMsg(res.Success))
	return resp, nil
}
