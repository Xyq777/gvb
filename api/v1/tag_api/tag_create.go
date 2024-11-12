package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
)

func (a TagApi) TagCreateApi(c *gin.Context) {
	var tagCreateReq req.TagCreateReq
	err := c.ShouldBind(&tagCreateReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	var tagModel dao.TagModel
	tagModel.Title = tagCreateReq.Title
	exist, err := tagModel.FindWithTitle(global.Db)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseFailedCreate), c, err)
		return
	}
	if exist {
		callback.FAIL(res.TagAlreadyExist, res.CodeMsg(res.TagAlreadyExist), c)
		return
	}
	err = tagModel.Create(global.Db)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseFailedCreate), c, err)
		return
	}
	callback.OK(tagModel, c)
}
