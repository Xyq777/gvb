package core

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Mmx233/EnvConfig"
	"gvb/config"
	"gvb/internal/global"
)

var Conf config.Config

func InitConf() {
	const ConfigFile = "config.toml"
	_, err := toml.DecodeFile(ConfigFile, &Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(Conf)
}
func InitConfWithEnv() {
	c := &config.Config{}
	// 使用自定义的 EnvConfig 库从环境变量加载配置
	EnvConfig.Load("MYSQL_", &c.Mysql)   // 加载 MySQL 环境变量
	EnvConfig.Load("LOGGER_", &c.Logger) // 加载 Logger 环境变量
	EnvConfig.Load("SYSTEM_", &c.System)
	global.Config = c
}
