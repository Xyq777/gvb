package system

import "fmt"

type App struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	LogLevel string `toml:"log_level"`
}

func (s App) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
