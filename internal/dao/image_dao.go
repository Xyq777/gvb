package dao

import (
	"errors"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/req"
)

func UpdateImage(req req.UpdateImageNameReq) (*models.BannerModel, error) {
	image, count, err := FindWithID(models.BannerModel{}, req.ID)
	if count == 0 {
		return nil, errors.New("图片不存在")
	}
	if err != nil {
		return nil, err
	}
	err = global.Db.Model(&image).Update("name", req.ImageName).Error
	if err != nil {
		return nil, err
	}
	return image, nil
}
