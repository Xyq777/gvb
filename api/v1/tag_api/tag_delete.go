package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
)

func (a TagApi) TagDeleteApi(c *gin.Context) {
	var tagReq req.TagReq
	err := c.ShouldBind(&tagReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	var tagModel dao.TagModel
	tagModel.Title = tagReq.Title
	err = tagModel.Delete(global.Db)
	if err != nil {
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
	}
	callback.OK(tagModel, c)

}
