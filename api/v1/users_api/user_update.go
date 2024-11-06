package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_user"
	"gvb/internal/tools/claimx"
)

func (a *UsersApi) UserUpdateApi(c *gin.Context) {
	var userUpdateReq req.UserUpdateReq
	err := c.ShouldBind(&userUpdateReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	claims, err := claimx.GetClaim(c)
	if err != nil {
		global.Log.Error("未设置claims")
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c)
		return
	}
	var userSrv srv_user.UserSrv
	resp, err := userSrv.UpdateNickname(claims, &userUpdateReq)
	if err != nil {
		callback.FAIL(resp.Code, resp.Msg, c, err)
		return
	}
	callback.OK(resp.Data, c)

}
