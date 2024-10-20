package config

type Config struct {
	Mysql  Mysql  `toml:"mysql"`
	Logger Logger `toml:"logger"`
	System System `toml:"system"`
}
type Mysql struct {
	host     string `toml:"host"`
	port     int    `toml:"port"`
	db       string `toml:"db"`
	user     string `toml:"user"`
	password string `toml:"password"`
	logLevel string `toml:"log_level"` //输出sql日志等级 dev/release
}
type Logger struct {
	level        string `toml:"level"`
	prefix       string `toml:"prefix"`
	director     string `toml:"director"`
	showLine     bool   `toml:"show_line"`
	logInConsole bool   `toml:"log_in_console"`
}
type System struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Env  string `toml:"env"`
}
