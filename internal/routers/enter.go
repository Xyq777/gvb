package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routerGroup := RouterGroup{router}
	routerGroup.SettingsRouter()
	routerGroup.ImagesRouter()
	return router

}
