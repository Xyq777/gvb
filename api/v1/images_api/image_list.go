package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/dao"
	"gvb/internal/models"
	"gvb/internal/models/req"
	"gvb/internal/models/res"
)

func (a ImagesApi) ImageListApi(c *gin.Context) {
	var page req.Page
	err := c.ShouldBind(&page)
	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	list, count, err := dao.GetList(models.BannerModel{}, &page)
	if err != nil {
		res.FAIL(res.FailedGetImageList, "获取数据失败", c, err)
		return
	}
	res.OK(res.List{ModelList: list, Count: count}, c)
}
