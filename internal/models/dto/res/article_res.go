package res

import "time"

type ArticleListRes struct {
	ID        string    `json:"id" structs:"id" gorm:"primarykey"` // es的id
	CreatedAt time.Time `jsons:"created_at" structs:"created_at"`  // 创建时间
	UpdatedAt time.Time `json:"updated_at" structs:"updated_at"`   // 更新时间

	Title   string `json:"title" structs:"title"`     // 文章标题
	Brief   string `json:"brief" structs:"brief"`     // 文章简介
	Content string `json:"content" structs:"content"` // 文章内容

	BannerUrl string `json:"banner_url" structs:"banner_url"` // 文章封面

}
