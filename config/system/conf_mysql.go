package system

import "C"
import (
	"fmt"
)

type Mysql struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Db       string `toml:"db"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	LogLevel string `toml:"log_level"` //输出sql日志等级 dev/release
}

func (m *Mysql) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Db)
	return dsn
}
