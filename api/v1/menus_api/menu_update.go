package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/serializition/req"
	"gvb/internal/models/serializition/res"
	"strconv"
)

func (a *MenusApi) MenuUpdateAPi(c *gin.Context) {
	var menuReq req.MenuRequest
	err := c.ShouldBindJSON(&menuReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	_id, err := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	//gorm不提供外键的实际实例时，不会操作第三方表
	//所以需要先清空关联表，在更新
	menu, count, err := dao.FindWithID(models.MenuModel{}, uint(id))
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	if count == 0 {
		callback.FAIL(res.NotFound, res.CodeMsg(res.NotFound), c, err)
		return
	}

	//清空关联表
	err = global.Db.Model(menu).Association("Banners").Clear()
	if err != nil {
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
	}
	//更新
	menuBannerList := make([]models.MenuBannerModel, 0)
	for _, sort := range menuReq.ImageSortList {
		// 由于第三方表有额外字段，大抵只能这样手动添加第三方表
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   id,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 给第三张表入库
	err = dao.CreateMenuBanner(menuBannerList)
	if err != nil {
		global.Log.Debugln(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	menuRes, err := dao.UpdateMenu(id, &menuReq)
	if err != nil {
		global.Log.Debugln(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	callback.OK(menuRes, c)

}