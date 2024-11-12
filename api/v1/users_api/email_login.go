package users_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service"
)

func (a *UsersApi) UserEmailLoginApi(c *gin.Context) {
	var userLoginReq = &req.UserEmailLoginReq{}
	err := c.ShouldBind(&userLoginReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
	}
	var userSrv = service.Srv.NewUserSrv(c)
	resp, err := userSrv.EmailLogin(userLoginReq)
	if err != nil {
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	callback.OK(resp.Data, c)
}
