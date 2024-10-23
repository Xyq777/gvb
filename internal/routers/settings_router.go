package routers

import (
	v1 "gvb/api/v1"
)

func (r *RouterGroup) SettingsRouter() {
	group := r.Group("/settings")
	settingsApi := v1.ApiGroupApp.SettingsApi
	group.GET("", settingsApi.SettingsInfoView)
	group.PUT("", settingsApi.SettingsUpdate)
}
