package config

type Config struct {
	Mysql  Mysql  `toml:"mysql"`
	Logger Logger `toml:"logger"`
	System System `toml:"system"`
}

type Logger struct {
	Level        string `toml:"level"`
	Prefix       string `toml:"prefix"`
	Director     string `toml:"director"`
	ShowLine     bool   `toml:"show_line"`
	LogInConsole bool   `toml:"log_in_console"`
}
