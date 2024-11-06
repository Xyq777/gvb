package dao

import (
	"gorm.io/gorm"
	"gvb/internal/global"
	"gvb/internal/models/serializition/req"
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
func FindWithIDs[T any](model T, IDListProvider req.IDListProvider) (list []T, count int, err error) {
	res := global.Db.Find(&list, IDListProvider.GetIDList())
	count, err = getErrorAndCount(res)
	return list, count, err

}
func FindWithID[T any](model T, ID uint) (resModel *T, count int, err error) {
	res := global.Db.Find(&resModel, ID)
	count, err = getErrorAndCount(res)
	return resModel, count, err

}
func FindWithUsername[T any](model T, name string) (resModel *T, exist bool, err error) {
	res := global.Db.Find(&resModel, "username = ?", name)
	count, err := getErrorAndCount(res)
	exist = count != 0
	return resModel, exist, err

}
func getErrorAndCount(res *gorm.DB) (int, error) {
	return int(res.RowsAffected), res.Error

}
