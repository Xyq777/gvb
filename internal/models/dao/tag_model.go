package dao

import "gorm.io/gorm"

// TagModel 标签表
type TagModel struct {
	gorm.Model
	Title string `gorm:"size:16;comment:标签的名称" json:"title"` // 标签的名称
}
