package v1

import (
	"gvb/api/v1/images_api"
	"gvb/api/v1/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
