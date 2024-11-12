package v1

import (
	"gvb/api/v1/image_api"
	"gvb/api/v1/menu_api"
	"gvb/api/v1/setting_api"
	"gvb/api/v1/tag_api"
	"gvb/api/v1/user_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   image_api.ImagesApi
	MenusApi    menu_api.MenusApi
	UsersApi    user_api.UsersApi
	TagApi      tag_api.TagApi
}

var ApiGroupApp = new(ApiGroup)
