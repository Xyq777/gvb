package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/internal/models/res"
)

func (a *SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OK(&global.Config, c)
}
