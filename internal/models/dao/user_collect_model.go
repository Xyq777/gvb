package dao

import "time"

// UserCollectModel 自定义第三张表  记录用户什么时候收藏了什么文章
type UserCollectModel struct {
	UserID    uint `gorm:"primaryKey"`
	ArticleID uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
