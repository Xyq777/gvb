package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_es"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var page req.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}

	list, count, err := srv_es.CommList(page.Key, page.Page, page.Limit)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.ElasticsearchOperateError, res.CodeMsg(res.ElasticsearchOperateError), c, err)
		return
	}

	callback.OK(res.ListData{
		ModelList: list,
		Count:     count,
	}, c)
}
