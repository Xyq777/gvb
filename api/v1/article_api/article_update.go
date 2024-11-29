package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"time"
)

func (a ArticleApi) UpdateArticleApi(c *gin.Context) {
	var updateReq req.ArticleUpdateReq
	err := c.ShouldBindJSON(&updateReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	var bannerUrl string
	if updateReq.BannerID != 0 {
		b, exist, err := dao.FindWithID(dao.BannerModel{}, updateReq.BannerID)
		if err != nil {
			callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
			return
		}
		if !exist {
			callback.FAIL(res.BannerNotExist, res.CodeMsg(res.BannerNotExist), c)
			return
		}
		bannerUrl = b.Path
	}

	article := dao.ArticleModel{
		ID:        updateReq.ID,
		UpdatedAt: time.Now(),
		Title:     updateReq.Title,
		Keyword:   updateReq.Title,
		Brief:     updateReq.Brief,
		Content:   updateReq.Content,
		Category:  updateReq.Category,
		Source:    updateReq.Source,
		Link:      updateReq.Link,
		BannerID:  updateReq.BannerID,
		BannerUrl: bannerUrl,
		Tags:      updateReq.Tags,
	}
	err = article.UpdateInES()
	if err != nil {
		global.Log.Error(err)

		callback.FAIL(res.ElasticsearchOperateError, res.CodeMsg(res.ElasticsearchOperateError), c, err)
		return
	}
	callback.OK(res.EmptyData, c)
}
