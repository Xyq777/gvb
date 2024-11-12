package user_api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_redis"
	"gvb/tools/jwt"
)

func (a *UsersApi) UserLogoutApi(c *gin.Context) {
	rt, err := c.Cookie("refreshToken")
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c, errors.New("cookie没有refreshToken"))
		return
	}
	claims, err := jwt.ParseToken(rt)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
		return
	}
	exist, err := srv_redis.ExistToken(claims.UserID, claims.ID)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.RedisGetFailed, res.CodeMsg(res.RedisGetFailed), c)
		return
	}
	if !exist {
		callback.FAIL(res.AlreadyLogout, res.CodeMsg(res.AlreadyLogout), c)
		return
	}
	err = srv_redis.LogoutToken(claims.UserID, claims.ID)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.RedisDelFailed, res.CodeMsg(res.RedisDelFailed), c)
		return
	}
	callback.OK(res.EmptyData, c)
}
