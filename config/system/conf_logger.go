package system

type Logger struct {
	Level  string `toml:"level"` //info,debug,warn,error
	Prefix string `toml:"prefix"`
	//Director     string `toml:"director"`
	ShowLine     bool `toml:"show_line"`
	LogInConsole bool `toml:"log_in_console"`
}
