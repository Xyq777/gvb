package req

import "gvb/internal/models/ctype"

type ArticleReq struct {
	Title    string      `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Brief    string      `json:"brief"`                                   // 文章简介
	Content  string      `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string      `json:"category"`                                // 文章分类
	Source   string      `json:"source"`                                  // 文章来源
	Link     string      `json:"link"`                                    // 原文链接
	BannerID uint        `json:"banner_id"`                               // 文章封面id
	Tags     ctype.Array `json:"tags"`                                    // 文章标签
}
