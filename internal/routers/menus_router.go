package routers

import v1 "gvb/api/v1"

func (r *RouterGroup) MenusRouter() {
	group := r.Group("menu")
	menusApi := v1.ApiGroupApp.MenusApi
	group.POST("", menusApi.MenuCreateView)
	group.GET("", menusApi.MenuListView)
	group.PUT("/:id", menusApi.MenuUpdateAPi)
	group.DELETE("", menusApi.MenuRemoveApi)
}
