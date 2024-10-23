package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb/config/custom"
	"gvb/internal/core"
	"gvb/internal/global"
	"gvb/internal/models/res"
)

func (a *SettingsApi) SettingsUpdate(c *gin.Context) {
	var reqSiteInfoModel custom.SiteInfo
	err := c.ShouldBind(&reqSiteInfoModel)
	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", err, c)
	}
	global.Config.Custom.SiteInfo = reqSiteInfoModel
	err = core.SetToml()
	if err != nil {
		res.FAIL(res.FailedRewriteToml, "更新配置文件失败", err, c)
	}
	res.OK(global.Config.Custom.SiteInfo, c)
}
