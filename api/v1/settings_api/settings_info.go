package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/res"
)

func (a *SettingsApi) SettingsInfoView(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "site":
		callback.OK(&global.Config.Custom.SiteInfo, c)
	case "email":
		callback.OK(&global.Config.Custom.Email, c)
	case "qq":
		callback.OK(&global.Config.Custom.QQ, c)
	case "qiniu":
		callback.OK(&global.Config.Custom.QiNiu, c)
	default:
		callback.FAIL(res.InvalidParams, "错误的路径参数", c)
	}

}
