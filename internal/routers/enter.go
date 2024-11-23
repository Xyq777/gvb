package routers

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routerGroup := RouterGroup{router.Group("/api")}
	routerGroup.GET("/token", middleware.TokenRefresh)
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	routerGroup.MenusRouter()
	routerGroup.SettingUsersRouter()
	routerGroup.ArticleRouter()
	return router

}
