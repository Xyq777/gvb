package srv_user

import (
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/tools/encryptor"
	"gvb/tools/jwt"
)

func (s UserSrv) EmailLogin(user *req.UserEmailLoginReq) (*res.Response, error) {
	userModel, exist, err := dao.FindWithUsername(dao.UserModel{}, user.Username)
	if !exist {
		resp := res.NewResponse(res.UserNotExist, res.EmptyData, res.CodeMsg(res.UserNotExist))
		return resp, err
	}
	if err != nil {
		resp := res.NewResponse(res.DatabaseOperateError, res.EmptyData, res.CodeMsg(res.DatabaseOperateError))
		return resp, err
	}
	if !encryptor.CheckMd5([]byte(user.Password), userModel.Password) {
		resp := res.NewResponse(res.PasswordNotMatched, res.EmptyData, res.CodeMsg(res.PasswordNotMatched))
		return resp, err
	}
	payload := jwt.Payload{
		UserID:   userModel.ID,
		Username: userModel.Username,
		Nickname: userModel.Nickname,
		Role:     userModel.Role,
	}
	token, err := jwt.GenerateToken(payload)
	if err != nil {
		resp := res.NewResponse(res.TokenGenerateFailed, res.EmptyData, res.CodeMsg(res.TokenGenerateFailed))
		return resp, err
	}
	resp := res.NewResponse(res.Success, token, res.CodeMsg(res.Success))
	return resp, nil
}
