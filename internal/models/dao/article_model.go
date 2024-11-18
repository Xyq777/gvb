package dao

import (
	"gvb/internal/models/ctype"
	"time"
)

//es8的官方库的typedAi，文档很少，对于部分更新，目前只想到把json标签添加omitempty

type ArticleModel struct {
	ID        uint      `json:"id,omitempty" structs:"id" gorm:"primarykey"`            // es的id
	CreatedAt time.Time `json:"created_at,omitempty" structs:"created_at"`              // 创建时间
	UpdatedAt time.Time `json:"updated_at,omitempty" structs:"updated_at"`              // 更新时间
	DeletedAt time.Time `json:"deleted_at,omitempty" structs:"deleted_at" gorm:"index"` // 删除时间

	Title   string `json:"title,omitempty" structs:"title"`     // 文章标题
	Keyword string `json:"keyword,omitempty" structs:"keyword"` // 关键字
	Brief   string `json:"brief,omitempty" structs:"brief"`     // 文章简介
	Content string `json:"content,omitempty" structs:"content"` // 文章内容

	LookCount     int `json:"look_count,omitempty" structs:"look_count"`         // 浏览量
	CommentCount  int `json:"comment_count,omitempty" structs:"comment_count"`   // 评论量
	DiggCount     int `json:"digg_count,omitempty" structs:"digg_count"`         // 点赞量
	CollectsCount int `json:"collects_count,omitempty" structs:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id,omitempty" structs:"user_id"`               // 用户id
	UserNickName string `json:"user_nick_name,omitempty" structs:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"user_avatar,omitempty" structs:"user_avatar"`       // 用户头像

	Category string `json:"category,omitempty" structs:"category"` // 文章分类
	Source   string `json:"source,omitempty" structs:"source"`     // 文章来源
	Link     string `json:"link,omitempty" structs:"link"`         // 原文链接

	Banner    BannerModel `json:"-,omitempty"`                               // 文章封面
	BannerID  uint        `json:"banner_id,omitempty" structs:"banner_id"`   // 文章封面id
	BannerUrl string      `json:"banner_url,omitempty" structs:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags,omitempty" structs:"tags"` // 文章标签
}
