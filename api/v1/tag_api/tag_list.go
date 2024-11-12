package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
)

func (a TagApi) TagListApi(c *gin.Context) {
	var page req.Page
	err := c.ShouldBind(&page)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}

	tags, count, err := dao.GetList(dao.TagModel{}, &page)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	listData := res.ListData{
		ModelList: tags,
		Count:     count,
	}
	callback.OK(listData, c)
}
