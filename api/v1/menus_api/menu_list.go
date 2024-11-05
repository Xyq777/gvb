package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/serializition/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (a *MenusApi) MenuListView(c *gin.Context) {
	// 先查菜单
	menus, err := dao.ListMenus(50)

	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseFailedCreate, res.CodeMsg(res.DatabaseFailedCreate), c, err)
		return
	}

	callback.OK(menus, c)
}
