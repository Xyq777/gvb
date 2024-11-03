package menus_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/res"
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
		res.FAIL(res.DatabaseFailedCreate, res.ErrorMsg(res.DatabaseFailedCreate), c, err)
		return
	}

	res.OK(menus, c)
}
