package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service"
)

func (a *MenusApi) MenuRemoveApi(c *gin.Context) {
	var deleteReq = make([]req.DeleteReq, 0)
	err := c.ShouldBindJSON(&deleteReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	var srv = service.Srv.NewMenuSrv(c)
	resp, err := srv.Delete(deleteReq)
	if err != nil {
		callback.FAIL(resp.Code, resp.Msg, c, err)
		return
	}

	callback.OK(resp.Data, c)

}
