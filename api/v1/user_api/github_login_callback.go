package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/ctype"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/res"
	"gvb/internal/service"
	"gvb/internal/service/srv_github"
	"gvb/tools/random"
	"strconv"
)

func (a *UsersApi) UserGithubLoginCallback(c *gin.Context) {
	var userSrv = service.Srv.NewUserSrv(c)
	code := c.Query("code")
	if code == "" {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c)
		return
	}
	stateCallback := c.Query("state")
	if code == "" {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c)
		return
	}
	stateLocal, err := c.Cookie("githubState")
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.StateNotMatched, res.CodeMsg(res.StateNotMatched), c, err)
		return
	}
	if stateCallback != stateLocal {
		callback.FAIL(res.StateNotMatched, res.CodeMsg(res.StateNotMatched), c)
		return
	}
	githubInfo, err := srv_github.GetGithubInfo(code)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.GithubLoginFailed, res.CodeMsg(res.GithubLoginFailed), c, err)
		return
	}

	userModel, exist, err := dao.FindWithUsername(dao.UserModel{}, strconv.Itoa(githubInfo.Id))
	if exist {
		//jwt
		rt, at, err := userSrv.GenJwt(userModel)
		if err != nil {
			callback.FAIL(res.TokenGenerateFailed, res.CodeMsg(res.TokenGenerateFailed), c, err)
			return
		}
		dataResp := res.LoginRes{
			RefreshToken: *rt,
			AccessToken:  *at,
		}
		callback.OK(dataResp, c)
		return
	}
	password, err := random.GenerateString(16)
	if err != nil {
		callback.FAIL(res.SeverError, res.CodeMsg(res.SeverError), c, err)
		return
	}
	userModel.Password = password
	userModel.Username = strconv.Itoa(githubInfo.Id)
	userModel.Nickname = githubInfo.Login
	userModel.Avatar = githubInfo.AvatarUrl
	userModel.SignStatus = ctype.SignGithub
	err = userModel.Create(global.Db)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	//jwt
	rt, at, err := userSrv.GenJwt(userModel)
	if err != nil {
		callback.FAIL(res.TokenGenerateFailed, res.CodeMsg(res.TokenGenerateFailed), c, err)
		return
	}
	dataResp := res.LoginRes{
		RefreshToken: *rt,
		AccessToken:  *at,
	}
	callback.OK(dataResp, c)
}
