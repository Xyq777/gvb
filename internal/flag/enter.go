package flag

import (
	sys_flag "flag"
)

type Option struct {
	M bool //migrate table
	U bool //user
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("m", false, "-m 是否进行gorm的表迁移")
	user := sys_flag.Bool("u", false, "-u 数据库用户名")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		U: *user,
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
		return
	}
	if option.U {
		CreateUser()
		return
	}

}
