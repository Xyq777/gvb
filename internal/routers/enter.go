package routers

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/middleware"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routerGroup := RouterGroup{router}
	routerGroup.GET("/refresh", middleware.TokenRefresh)
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	routerGroup.MenusRouter()
	routerGroup.SettingUsersRouter()
	return router

}
