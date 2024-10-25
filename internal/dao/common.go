package dao

import (
	"gvb/internal/global"
	"gvb/internal/models/req"
)

func GetList[T any](model T, page *req.Page) (list []T, count int, err error) {
	offset := (page.Page - 1) * page.Limit
	if offset < 0 {
		offset = 0
	}
	if page.Sort == "" {
		page.Sort = "created_at desc"
	}
	result := global.Db.Limit(page.Limit).Order(page.Sort).Offset(offset).Find(&list)
	err = result.Error
	count = int(result.RowsAffected)
	return list, count, err
}
