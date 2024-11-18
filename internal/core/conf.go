package core

import (
	"github.com/BurntSushi/toml"
	"github.com/Mmx233/EnvConfig"
	"gvb/config"
	"gvb/internal/global"
	"os"
)

const ConfigFile = "config.toml"

func InitConf() {
	c := &config.Config{}
	initSystemConfWithEnv(c)
	initCustomWithToml(c)
	global.Config = c
}
func initSystemConfWithEnv(c *config.Config) {
	// 使用自定义的 EnvConfig 库从环境变量加载配置
	EnvConfig.Load("MYSQL_", &c.System.Mysql)   // 加载 MySQL 环境变量
	EnvConfig.Load("LOGGER_", &c.System.Logger) // 加载 Logger 环境变量
	EnvConfig.Load("SYSTEM_", &c.System.App)
	EnvConfig.Load("JWT_", &c.System.Jwt)
	EnvConfig.Load("UPLOAD_", &c.System.Upload)
	EnvConfig.Load("REDIS_", &c.System.Redis)
	EnvConfig.Load("EMAIL_", &c.System.Email)
	EnvConfig.Load("GITHUB_", &c.System.Github)
	EnvConfig.Load("ES_", &c.System.ES)
}
func initCustomWithToml(c *config.Config) {

	_, err := toml.DecodeFile(ConfigFile, &c.Custom)
	if err != nil {
		panic(err)
	}
}
func SetToml() error {
	file, err := os.OpenFile(ConfigFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		global.Log.Errorln(err)
		return err
	}
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(global.Config.Custom)
	if err != nil {
		global.Log.Errorln(err)
		return err
	}
	return nil
}
