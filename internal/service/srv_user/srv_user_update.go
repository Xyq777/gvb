package srv_user

import (
	"errors"
	"gvb/internal/dao"
	"gvb/internal/models/ctype"
	"gvb/internal/models/serializition/req"
	"gvb/internal/models/serializition/res"
	"gvb/tools/jwt"
)

func (s UserSrv) UpdateNickname(claims *jwt.CustomClaims, userUpdateReq req.UserUpdateReq) (*res.Response, error) {
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
