package flag

import sys_flag "flag"

type Option struct {
	M bool //migrate table
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("m", false, "是否进行gorm的表迁移")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		M: *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.M {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.M {
		Makemigrations()
	}
}
