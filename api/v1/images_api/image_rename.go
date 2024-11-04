package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/dao"
	"gvb/internal/models/req"
	"gvb/internal/models/res"
)

func (a ImagesApi) ImageRenameApi(c *gin.Context) {
	var updateReq req.UpdateImageNameReq
	err := c.ShouldBind(&updateReq)
	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	image, err := dao.UpdateImage(updateReq)
	if err != nil {
		res.FAIL(res.InvalidParams, "修改失败", c, err)
		return
	}
	res.OK(image, c)
}
