package routers

import v1 "gvb/api/v1"

func (r *RouterGroup) MenusRouter() {
	group := r.Group("menu")
	menusApi := v1.ApiGroupApp.MenusApi
	group.POST("", menusApi.MenuCreateView)
}
