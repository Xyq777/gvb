package dao

import (
	"gorm.io/gorm"
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
	res := global.Db.Limit(page.Limit).Order(page.Sort).Offset(offset).Find(&list)
	count, err = getErrorAndCount(res)
	return list, count, err
}
func FindWithIDs[T any](model T, IDList []uint) (list []T, count int, err error) {
	res := global.Db.Find(&list, IDList)
	count, err = getErrorAndCount(res)
	return list, count, err

}
func getErrorAndCount(res *gorm.DB) (int, error) {
	return int(res.RowsAffected), res.Error

}
