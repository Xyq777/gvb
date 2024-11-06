package dao

import (
	"errors"
	"gorm.io/gorm"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"strconv"
)

func CreateMenu(menu *req.MenuRequest) (*res.MenuResponse, error) {
	// 创建banner数据入库
	menuModel := models.MenuModel{
		MenuTitle:    menu.MenuTitle,
		MenuTitleEn:  menu.MenuTitleEn,
		Slogan:       menu.Slogan,
		Abstract:     menu.Abstract,
		AbstractTime: menu.AbstractTime,
		BannerTime:   menu.BannerTime,
		Sort:         menu.Sort,
	}

	err := global.Db.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}
	return &res.MenuResponse{
		MenuTitle: menuModel.MenuTitle,
		MenuID:    menuModel.ID,
	}, nil

}
func CreateMenuBanner(menuBannerList []models.MenuBannerModel) error {
	//检查banner是否存在
	for _, menuBanner := range menuBannerList {
		_, count, err := FindWithID(models.BannerModel{}, menuBanner.BannerID)
		if count == 0 {
			return errors.New(strconv.Itoa(int(menuBanner.BannerID)) + "图片不存在")
		}
		if err != nil {
			return err
		}
	}
	err := global.Db.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}
	return nil
}
func ListMenus(limit int) (menuList []models.MenuModel, err error) {
	err = global.Db.Preload("Banners").Order("sort desc").Find(&menuList).Select("id").Limit(limit).Error

	if err != nil {
		return nil, err
	}

	return menuList, nil

}
func UpdateMenu(ID uint, menu *req.MenuRequest) (*res.MenuResponse, error) {
	menuModel := models.MenuModel{
		Model:        gorm.Model{ID: ID},
		MenuTitle:    menu.MenuTitle,
		MenuTitleEn:  menu.MenuTitleEn,
		Slogan:       menu.Slogan,
		Abstract:     menu.Abstract,
		AbstractTime: menu.AbstractTime,
		BannerTime:   menu.BannerTime,
		Sort:         menu.Sort,
	}
	err := global.Db.Model(&menuModel).Where("ID = ?", menuModel.ID).Updates(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}
	return &res.MenuResponse{
		MenuTitle: menuModel.MenuTitle,
		MenuID:    menuModel.ID,
	}, nil

}
func DeleteMenu() {

}
