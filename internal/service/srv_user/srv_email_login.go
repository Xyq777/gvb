package srv_user

import (
	"errors"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/tools/encryptor"
)

func (s UserSrv) EmailLogin(user *req.UserEmailLoginReq) (*res.Response, error) {
	userModel, exist, err := dao.FindWithUsername(dao.UserModel{}, user.Username)
	if !exist {
		resp := res.NewResponse(res.UserNotExist, res.EmptyData, res.CodeMsg(res.UserNotExist))
		return resp, errors.New(res.CodeMsg(res.UserNotExist))
	}
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	if !encryptor.CheckMd5([]byte(user.Password), userModel.Password) {
		resp := res.NewResponse(res.PasswordNotMatched, res.EmptyData, res.CodeMsg(res.PasswordNotMatched))
		return resp, err
	}
	//jwt

	rt, at, err := s.GenJwt(userModel)
	if err != nil {
		resp := res.NewResponse(res.TokenGenerateFailed, res.EmptyData, res.CodeMsg(res.TokenGenerateFailed))
		return resp, err
	}
	dataResp := res.LoginRes{
		RefreshToken: *rt,
		AccessToken:  *at,
	}
	resp := res.NewResponse(res.Success, dataResp, res.CodeMsg(res.Success))
	return resp, nil
}
