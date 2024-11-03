package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/req"
	"gvb/internal/models/res"
)

func (a *MenusApi) MenuCreateView(c *gin.Context) {
	var menuReq req.MenuRequest
	err := c.ShouldBindJSON(&menuReq)
	if err != nil {
		res.FAIL(res.InvalidParams, res.ErrorMsg(res.InvalidParams), c, err)

		return
	}

	// 重复值判断
	menuRes, err := dao.CreateMenu(&menuReq)

	if err != nil {
		global.Log.Error(err)
		res.FAIL(res.DatabaseFailedCreate, res.ErrorMsg(res.DatabaseFailedCreate), c, err)
		return
	}
	if len(menuReq.ImageSortList) == 0 {
		res.OK(struct{}{}, c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range menuReq.ImageSortList {
		// 这里也得判断image_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuRes.MenuID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 给第三张表入库
	err = dao.CreateMenuBanner(menuBannerList)
	if err != nil {
		global.Log.Error(err)
		res.FAIL(res.DatabaseFailedCreate, "菜单图片关联失败", c)
		return
	}
	res.OK(menuRes, c)
}
