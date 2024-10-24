package routers

import (
	v1 "gvb/api/v1"
)

func (r *RouterGroup) SettingsRouter() {
	group := r.Group("/settings")
	settingsApi := v1.ApiGroupApp.SettingsApi
	group.GET("/:name", settingsApi.SettingsInfoView)
	group.PUT("/:name", settingsApi.SettingsUpdate)
}
