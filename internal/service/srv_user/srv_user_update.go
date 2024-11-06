package srv_user

import (
	"errors"
	"gvb/internal/dao"
	"gvb/internal/global"
	"gvb/internal/models/ctype"
	dao2 "gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/tools/encryptor"
	"gvb/tools/jwt"
)

func (s UserSrv) UpdateNickname(claims *jwt.CustomClaims, userUpdateReq *req.UserUpdateReq) (*res.Response, error) {
	var userDao = dao.NewUserDao()
	//普通用户
	if claims.Role != ctype.PermissionAdmin {
		if claims.UserID != userUpdateReq.UserID {
			resp := res.NewResponse(res.PermissionDenied, res.EmptyData, "权限不足")
			return resp, errors.New("PermissionDenied")
		}

		err := userDao.UpdateUserNickname(userUpdateReq.UserID, userUpdateReq.Nickname)
		if err != nil {
			resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, "更新失败")
			return resp, err
		}
		return res.NewResponse(res.Success, res.EmptyData, res.CodeMsg(res.Success)), nil
	}
	//管理员
	err := userDao.UpdateUserNickname(userUpdateReq.UserID, userUpdateReq.Nickname)
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, "更新失败")
		return resp, err
	}
	return res.NewResponse(res.Success, res.EmptyData, res.CodeMsg(res.Success)), nil
}

func (s UserSrv) UpdatePassword(claims *jwt.CustomClaims, UpdateReq *req.UserUpdatePasswordReq) (*res.Response, error) {
	user, exist, err := dao2.FindWithID(dao2.UserModel{}, claims.UserID)
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	if !exist {
		resp := res.NewResponse(res.NotFound, res.EmptyData, "用户不存在")
		return resp, errors.New("用户不存在")
	}

	if !encryptor.CheckMd5([]byte(UpdateReq.OldPassword), user.Password) {
		resp := res.NewResponse(res.PasswordNotMatch, res.EmptyData, "密码错误")
		return resp, errors.New("密码错误")
	}
	user.Password = encryptor.Md5([]byte(UpdateReq.NewPassword))
	err = user.Update(global.Db)
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	return res.NewResponse(res.Success, res.EmptyData, res.CodeMsg(res.Success)), nil
}
