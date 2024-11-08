package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/ctype"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service"
	"gvb/internal/tools/claimx"
)

func (a *UsersApi) UserDeleteApi(c *gin.Context) {
	var userDeleteReq req.DeleteReq
	err := c.ShouldBind(&userDeleteReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
	}
	claims, err := claimx.GetClaim(c)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c, err)
		return
	}
	if claims.Role != ctype.PermissionAdmin {
		if claims.UserID != userDeleteReq.ID {
			callback.FAIL(res.PermissionDenied, res.CodeMsg(res.PermissionDenied), c, err)
			return
		}
	}

	var userSrv = service.Srv.NewUserSrv(c)
	resp, err := userSrv.DeleteUser(userDeleteReq.ID)
	if err != nil {
		callback.FAIL(resp.Code, resp.Msg, c, err)
		return
	}
	callback.OK(resp.Data, c)

}
