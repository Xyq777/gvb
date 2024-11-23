package dao

import "time"

// UserCollectModel 自定义第三张表  记录用户什么时候收藏了什么文章

type CollectModel struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint
	ArticleID string    `gorm:"size:32;comment:文章的es id"`
	CreatedAt time.Time `gorm:"comment:收藏的时间"`
}
