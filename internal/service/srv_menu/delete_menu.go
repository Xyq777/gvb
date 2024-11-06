package srv_menu

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
)

func (s MenuSrv) Delete(deleteReq []req.DeleteReq) (*res.Response, error) {
	menuList, count, err := dao.FindWithIDs(dao.MenuModel{}, req.DeleteReqList(deleteReq))
	if count != len(deleteReq) {
		global.Log.Debugln("未找到菜单")
		response := res.NewResponse(res.NotFound, res.EmptyData, res.CodeMsg(res.NotFound))
		return response, errors.New("未找到菜单")
	}
	if err != nil {
		global.Log.Debugln(err)
		response := res.NewResponse(res.DatabaseMenuFailedDelete, res.EmptyData, res.CodeMsg(res.DatabaseMenuFailedDelete))
		return response, err
	}

	// 事务
	err = global.Db.Transaction(func(tx *gorm.DB) error {
		err = global.Db.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.Db.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		response := res.NewResponse(res.DatabaseMenuFailedDelete, res.EmptyData, res.CodeMsg(res.DatabaseMenuFailedDelete))
		return response, err
	}
	respnse := res.NewResponse(res.Success, fmt.Sprintf("成功删除%d个菜单", count), res.CodeMsg(res.Success))
	return respnse, err
}
