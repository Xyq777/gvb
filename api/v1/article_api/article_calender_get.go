package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_es"
)

func (a ArticleApi) GetArticleCalender(c *gin.Context) {
	buckets, err := srv_es.DateAggregateSearch(dao.ArticleModel{}.Index(), "created_at", "day")
	if err != nil {
		callback.FAIL(res.ElasticsearchOperateError, res.CodeMsg(res.ElasticsearchOperateError), c, err)
		return
	}
	calender := make([]res.ArticleCalenderRes, 0)
	for _, bucket := range buckets {
		calender = append(calender,
			res.ArticleCalenderRes{
				Day:   *bucket.KeyAsString,
				Count: int(bucket.DocCount),
			},
		)
	}
	callback.OK(
		res.ListData{
			ModelList: calender,
			Count:     len(calender),
		}, c)
}
