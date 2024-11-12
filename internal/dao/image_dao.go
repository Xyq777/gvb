package dao

import (
	"errors"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
)

func UpdateImage(req req.UpdateImageNameReq) (*dao.BannerModel, error) {
	image, exist, err := dao.FindWithID(dao.BannerModel{}, req.ID)
	if !exist {
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
