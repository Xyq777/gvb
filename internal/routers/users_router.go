package routers

import (
	v1 "gvb/api/v1"
	"gvb/internal/middleware"
)

func (r *RouterGroup) SettingUsersRouter() {
	g := r.Group("/user")
	usersApi := v1.ApiGroupApp.UsersApi
	g.GET("/logout", usersApi.UserLogoutApi)
	g.POST("/login", usersApi.UserEmailLoginApi)

	g.DELETE("", middleware.JwtAuth(), usersApi.UserDeleteApi)
	g.PUT("", middleware.JwtAuth(), usersApi.UserUpdateApi)

	g.POST("/list", middleware.JwtAuth(), usersApi.UserListApi)
	g.POST("/email", middleware.JwtAuth(), usersApi.UserBindEmailApi)
	g.PUT("/password", middleware.JwtAuth(), usersApi.UserUpdatePasswordApi)

	g.GET("/github/login", usersApi.UserGithubLoginApi)
	g.GET("/github/callback", usersApi.UserGithubLoginCallback)

}
