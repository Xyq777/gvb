package flag

import (
	"fmt"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/ctype"
	"gvb/tools/Encryptor"
)

func CreateUser() {
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		userName string
		nickName string
		password string
		email    string
		avatar   = "/uploads/avatar/default.jpg"
	)
	for true {
		fmt.Printf("请输入用户名：")
		fmt.Scan(&userName)
		count := global.Db.Find(&models.UserModel{}, "user_name = ?", userName).RowsAffected
		if count != 0 {
			fmt.Println("用户名已存在")
			continue
		}
		break
	}

	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	hashPwd := Encryptor.Md5([]byte(password))
	err := global.Db.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       ctype.PermissionAdmin,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err.Error() + "创建用户失败")
		return
	}
	global.Log.Infoln("创建用户成功")

}