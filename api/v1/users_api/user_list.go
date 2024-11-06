package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/serializition/req"
	"gvb/internal/models/serializition/res"
	"gvb/internal/service"
	"gvb/tools/jwt"
)

func (a *UsersApi) UserListApi(c *gin.Context) {
	_claims, exist := c.Get("claims")
	if !exist {
		global.Log.Error("未设置claims")
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
		return
	}
	claims := _claims.(*jwt.CustomClaims)
	var page req.Page
	err := c.ShouldBind(&page)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	var userSrv = service.Srv.NewUserSrv(c)
	resp, err := userSrv.UserList(page, claims.Role)
	if err != nil {
		callback.FAIL(resp.Code, resp.Msg, c, err)
		return
	}
	callback.OK(resp.Data, c)
}
