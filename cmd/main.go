package main

import (
	"gvb/internal/core"
	"gvb/internal/global"
	"gvb/internal/routers"
)

func main() {
	//初始化配置文件
	core.InitConfWithEnv()
	//初始化日志
	global.Log = core.InitLogger()
	//初始化数据库
	global.Db = core.InitGorm()
	//初始化路由
	router := routers.InitRouter()
	global.Log.Info("server run at ", global.Config.System.Addr())
	router.Run(global.Config.System.Addr())
}
