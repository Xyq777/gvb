package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/req"
	"gvb/internal/models/res"
)

func (a ImagesApi) ImageDeleteApi(c *gin.Context) {
	var deleteImagesID []req.DeleteReq
	err := c.ShouldBind(&deleteImagesID)

	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	IDList := req.GetIDList(deleteImagesID)
	imageList, count, err := dao.FindWithIDs(models.BannerModel{}, IDList)
	if err != nil {
		res.FAIL(res.FailedGetImageList, "数据库查询失败", c, err)
		return
	}
	if count != len(IDList) {
		res.FAIL(res.NotFoundImages, "图片不存在", c, err)
		return
	}
	err = global.Db.Delete(&imageList).Error
	if err != nil {
		res.FAIL(res.FailedDeleteImages, "删除图片失败", c, err)
		return
	}
	res.OK(IDList, c)

}
