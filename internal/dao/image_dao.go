package dao

import (
	"errors"
	"gvb/internal/global"
	"gvb/internal/models"
)

func UpdateImage(ID uint, name string) (*models.BannerModel, error) {
	image, count, err := FindWithIDs(models.BannerModel{}, []uint{ID})
	if count == 0 {
		return nil, errors.New("图片不存在")
	}
	if err != nil {
		return nil, err
	}
	err = global.Db.Model(&image).Update("name", name).Error
	if err != nil {
		return nil, err
	}
	return &image[0], nil
}
