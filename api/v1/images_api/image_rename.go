package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/dao"
	"gvb/internal/models/serializition/req"
	"gvb/internal/models/serializition/res"
)

func (a ImagesApi) ImageRenameApi(c *gin.Context) {
	var updateReq req.UpdateImageNameReq
	err := c.ShouldBind(&updateReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	image, err := dao.UpdateImage(updateReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, "修改失败", c, err)
		return
	}
	callback.OK(image, c)
}
