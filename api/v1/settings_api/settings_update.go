package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb/config/custom"
	"gvb/internal/callback"
	"gvb/internal/core"
	"gvb/internal/models/dto/res"
)

func (a *SettingsApi) SettingsUpdate(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "site":
		var InfoModel custom.SiteInfo
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			callback.FAIL(res.InvalidParams, "参数错误", c, err)
			return
		}

	case "email":
		var InfoModel custom.Email
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			callback.FAIL(res.InvalidParams, "参数错误", c, err)
			return
		}

	case "qq":
		fmt.Println("1")
		var InfoModel custom.QQ
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			callback.FAIL(res.InvalidParams, "参数错误", c, err)
			return
		}
		fmt.Println(InfoModel)

	case "qiniu":
		var InfoModel custom.QiNiu
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			callback.FAIL(res.InvalidParams, "参数错误", c, err)
			return
		}

	default:
		callback.FAIL(res.InvalidParams, "错误的路径参数", c)
		return
	}

	err := core.SetToml()
	if err != nil {
		callback.FAIL(res.FailedRewriteToml, "更新配置文件失败", c, err)
	}
	callback.OK(struct{}{}, c)
}
