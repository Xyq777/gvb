package routers

import (
	v1 "gvb/api/v1"
	"gvb/internal/middleware"
)

func (r *RouterGroup) SettingUsersRouter() {
	g := r.Group("/user")
	usersApi := v1.ApiGroupApp.UsersApi
	g.GET("/logout", usersApi.UserLogoutApi)
	g.POST("/list", middleware.JwtAuth(), usersApi.UserListApi)
	g.POST("/login", usersApi.UserEmailLoginApi)
	g.PUT("", middleware.JwtAuth(), usersApi.UserUpdateApi)
	g.PUT("/password", middleware.JwtAuth(), usersApi.UserUpdatePasswordApi)
	g.DELETE("", middleware.JwtAuth(), usersApi.UserDeleteApi)

}
