package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service"
	"gvb/internal/tools/claimx"
)

func (a *UsersApi) UserUpdatePasswordApi(c *gin.Context) {
	claims, err := claimx.GetClaim(c)
	if err != nil {
		global.Log.Error("未设置claims")
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
		return
	}

	var updatePasswordReq req.UserUpdatePasswordReq
	err = c.ShouldBind(&updatePasswordReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c)
		return
	}
	var userService = service.Srv.NewUserSrv(c)
	resp, err := userService.UpdatePassword(claims, &updatePasswordReq)
	if err != nil {
		callback.FAIL(resp.Code, resp.Msg, c, err)
		return
	}
	callback.OK(resp.Data, c)
}
