package srv_es

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"gvb/internal/global"
	"gvb/internal/models/dao"
)

func DeleteArticleInBulk(IDList []string) (int, error) {
	count := len(IDList)
	bulk := global.ES.Bulk().Index(dao.ArticleModel{}.Index())
	for _, ID := range IDList {
		err := bulk.DeleteOp(
			types.DeleteOperation{
				Id_: &ID,
			},
		)
		if err != nil {
			global.Log.Error(err)
			count--
		}
	}
	_, err := bulk.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return 0, err
	}
	if count == 0 {
		return 0, errors.New("all delete failed")
	}
	return count, nil
}
