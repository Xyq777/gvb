package srv_user

import (
	"errors"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/res"
)

func (s UserSrv) DeleteUser(UserID uint) (*res.Response, error) {
	var user dao.UserModel
	user.ID = UserID
	userModel, exist, err := dao.FindWithID(dao.UserModel{}, UserID)
	if !exist {
		resp := res.NewResponse(res.NotFound, res.EmptyData, res.CodeMsg(res.NotFound))
		return resp, errors.New("未找到要删除的user")
	}
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	err = userModel.Delete(global.Db)
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	return res.NewResponse(res.Success, res.EmptyData, res.CodeMsg(res.Success)), nil

}
