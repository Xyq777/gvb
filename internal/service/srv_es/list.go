package srv_es

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/sirupsen/logrus"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/res"
)

func CommList(key string, page, limit int) (list []res.ArticleListRes, count int, err error) {

	boolSearch := &types.BoolQuery{}
	if key != "" {
		boolSearch.Must = append(boolSearch.Must, types.Query{
			Match: map[string]types.MatchQuery{"title": {Query: key}},
		})
	}
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	resp, err := global.ES.
		Search().
		From(page).
		Size(limit).
		Index(dao.ArticleModel{}.Index()).
		Query(
			&types.Query{
				Bool: boolSearch,
			}).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return nil, 0, err
	}
	//TODO more detailed error handle
	count = len(resp.Hits.Hits) //搜索到结果总条数
	demoList := make([]res.ArticleListRes, 0)
	for _, hit := range resp.Hits.Hits {
		var model res.ArticleListRes
		data, err := hit.Source_.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			count--
			continue
		}
		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			count--
			continue
		}
		model.ID = *hit.Id_
		demoList = append(demoList, model)
	}
	return demoList, count, err
}
