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
	var deleteImagesReq []req.DeleteReq
	err := c.ShouldBind(&deleteImagesReq)

	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	var deleteReqList req.DeleteReqList = deleteImagesReq
	imageList, count, err := dao.FindWithIDs(models.BannerModel{}, deleteReqList)
	if err != nil {
		res.FAIL(res.FailedGetImageList, "数据库查询失败", c, err)
		return
	}
	if count != len(deleteImagesReq) {
		res.FAIL(res.NotFoundImages, "图片不存在", c, err)
		return
	}
	err = global.Db.Delete(&imageList).Error
	if err != nil {
		res.FAIL(res.FailedDeleteImages, "删除图片失败", c, err)
		return
	}

	res.OK(deleteReqList, c)

}
