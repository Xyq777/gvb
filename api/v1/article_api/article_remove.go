package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_es"
)

func (a ArticleApi) RemoveArticleApi(c *gin.Context) {
	var removeReq req.ArticleRemoveReq
	err := c.ShouldBind(&removeReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	count, err := srv_es.DeleteArticleInBulk(removeReq.IDList)
	if err != nil {
		callback.FAIL(res.ElasticsearchOperateError, res.CodeMsg(res.ElasticsearchOperateError), c, err)
		return
	}

	callback.OK(
		res.ListData{
			ModelList: res.EmptyData,
			Count:     count,
		}, c)
}
