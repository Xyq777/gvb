package dao

import (
	"gorm.io/gorm"
	"gvb/internal/global"
	"gvb/internal/models/ctype"
)

// UserModel 用户表
type UserModel struct {
	gorm.Model
	Nickname       string           `gorm:"size:36" json:"nickname"`             // 昵称
	Username       string           `gorm:"size:36" json:"username"`             // 用户名
	Password       string           `gorm:"size:128" json:"-"`                   // 密码
	Avatar         string           `gorm:"size:256" json:"avatar_id"`           // 头像id
	Email          string           `gorm:"size:128" json:"email"`               // 邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                  // 手机号
	Addr           string           `gorm:"size:64" json:"addr"`                 // 地址
	Token          string           `gorm:"size:64" json:"token"`                // 其他平台的唯一id
	IP             string           `gorm:"size:20" json:"ip"`                   // ip地址
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`        // 权限  1 管理员  2 普通用户  3 游客
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"` // 注册来源
	CollectsModels []CollectModel   `gorm:"foreignKey:UserID" json:"-"`          // 收藏的文章列表
}

func (u *UserModel) Update(tx *gorm.DB) error {
	return tx.Model(&UserModel{}).Where("id = ?", u.ID).Updates(u).Error

}
func (u *UserModel) Delete(tx *gorm.DB) error {
	err := tx.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
		err := tx.Delete(&u).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}
func (u *UserModel) Create(tx *gorm.DB) error {
	return tx.Create(&u).Error
}
