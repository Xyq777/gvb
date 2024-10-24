package settings_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb/config/custom"
	"gvb/internal/core"
	"gvb/internal/models/res"
)

func (a *SettingsApi) SettingsUpdate(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "site":
		var InfoModel custom.SiteInfo
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			res.FAIL(res.InvalidParams, "参数错误", err, c)
			return
		}

	case "email":
		var InfoModel custom.Email
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			res.FAIL(res.InvalidParams, "参数错误", err, c)
			return
		}

	case "qq":
		fmt.Println("1")
		var InfoModel custom.QQ
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			res.FAIL(res.InvalidParams, "参数错误", err, c)
			return
		}
		fmt.Println(InfoModel)

	case "qiniu":
		var InfoModel custom.QiNiu
		err := c.ShouldBind(&InfoModel)
		if err != nil {
			res.FAIL(res.InvalidParams, "参数错误", err, c)
			return
		}

	default:
		res.FAIL(res.InvalidParams, "错误的路径参数", errors.New(""), c)
		return
	}

	err := core.SetToml()
	if err != nil {
		res.FAIL(res.FailedRewriteToml, "更新配置文件失败", err, c)
	}
	res.OK(struct{}{}, c)
}
