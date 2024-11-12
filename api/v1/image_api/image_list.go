package image_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
)

func (a ImagesApi) ImageListApi(c *gin.Context) {
	var page req.Page
	err := c.ShouldBind(&page)
	if err != nil {
		callback.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	list, count, err := dao.GetList(dao.BannerModel{}, &page)
	if err != nil {
		callback.FAIL(res.FailedGetImageList, "获取数据失败", c, err)
		return
	}
	callback.OK(res.List{ModelList: list, Count: count}, c)
}
