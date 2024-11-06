package dao

import (
	"gorm.io/gorm"
	"gvb/internal/global"
	"gvb/internal/models"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{DB: global.Db}
}
func (d UserDao) UpdateUserNickname(id uint, name string) error {
	return d.Model(&models.UserModel{}).Where("id = ?", id).Update("nickname", name).Error

}
