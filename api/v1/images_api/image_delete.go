package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/serializition/req"
	"gvb/internal/models/serializition/res"
)

func (a ImagesApi) ImageDeleteApi(c *gin.Context) {
	var deleteImagesReq []req.DeleteReq
	err := c.ShouldBind(&deleteImagesReq)

	if err != nil {
		callback.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	var deleteReqList req.DeleteReqList = deleteImagesReq
	imageList, count, err := dao.FindWithIDs(models.BannerModel{}, deleteReqList)
	if err != nil {
		callback.FAIL(res.FailedGetImageList, "数据库查询失败", c, err)
		return
	}
	if count != len(deleteImagesReq) {
		callback.FAIL(res.NotFoundImages, "图片不存在", c, err)
		return
	}
	err = global.Db.Delete(&imageList).Error
	if err != nil {
		callback.FAIL(res.FailedDeleteImages, "删除图片失败", c, err)
		return
	}

	callback.OK(deleteReqList, c)

}
