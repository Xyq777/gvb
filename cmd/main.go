package main

import (
	"gvb/internal/core"
	"gvb/internal/flag"
	"gvb/internal/global"
	"gvb/internal/routers"
)

func main() {
	//初始化配置文件
	core.InitConf()

	//初始化日志
	global.Log = core.InitLogger()
	//初始化数据库
	global.Db = core.InitGorm()
	global.Redis = core.ConnectRedis()
	//读取命令行参数
	op := flag.Parse()
	if flag.IsWebStop(op) {
		flag.SwitchOption(op)
		return
	} else {
		flag.SwitchOption(op)
	}
	//初始化路由
	router := routers.InitRouter()
	global.Log.Info("server run at ", global.Config.System.App.Addr())
	router.Run(global.Config.System.App.Addr())
}
