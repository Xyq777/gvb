package settings_api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/internal/models/res"
)

func (a *SettingsApi) SettingsInfoView(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "site":
		res.OK(&global.Config.Custom.SiteInfo, c)
	case "email":
		res.OK(&global.Config.Custom.Email, c)
	case "qq":
		res.OK(&global.Config.Custom.QQ, c)
	case "qiniu":
		res.OK(&global.Config.Custom.QiNiu, c)
	default:
		res.FAIL(res.InvalidParams, "错误的路径参数", errors.New(""), c)
	}

}
