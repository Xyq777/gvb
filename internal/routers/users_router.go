package routers

import (
	v1 "gvb/api/v1"
	"gvb/internal/middleware"
)

func (r *RouterGroup) SettingUsersRouter() {
	g := r.Group("/user")
	usersApi := v1.ApiGroupApp.UsersApi

	g.POST("/list", middleware.JwtAuth(), usersApi.UserListApi)
	g.POST("/login", usersApi.UserEmailLoginApi)
	g.POST("", middleware.JwtAuth(), usersApi.UserUpdateApi)
}
