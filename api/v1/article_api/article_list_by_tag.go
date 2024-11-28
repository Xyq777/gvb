package article_api

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_es"
)

func (a ArticleApi) GetArticleListByTag(c *gin.Context) {
	aggName := "tags"
	subAggName := "articles"
	buckets, err := srv_es.TagAggregateSearch(dao.ArticleModel{}.Index(), aggName, subAggName)
	if err != nil {
		callback.FAIL(res.ElasticsearchOperateError, res.CodeMsg(res.ElasticsearchOperateError), c, err)
		global.Log.Error(err)
		return
	}
	var resp []res.ArticleListByTagRes
	for _, bucket := range buckets {
		r := res.ArticleListByTagRes{}
		r.TagName = bucket.Key.(string)
		r.Count = int(bucket.DocCount)
		for _, hit := range bucket.Aggregations[subAggName].(*types.TopHitsAggregate).Hits.Hits {
			article := res.ArticleByTagResp{}
			err := json.Unmarshal(hit.Source_, &article)
			if err != nil {
				callback.FAIL(res.ElasticsearchOperateError, res.CodeMsg(res.ElasticsearchOperateError), c, err)
				global.Log.Error(err)
				return
			}
			r.ArticleList = append(r.ArticleList, res.ArticleByTagResp{
				ID:    *hit.Id_,
				Title: article.Title,
			})
		}
		resp = append(resp, r)
	}
	callback.OK(resp, c)
}
