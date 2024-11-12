package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/service"
	"gvb/internal/tools/claimx"
	"gvb/tools/random"
	"time"
)

// 只使用加密逻辑，不存服务器
var store = sessions.NewCookieStore([]byte("HQBVQKWB@5@"))

// UserBindEmailApi TODO 为了线程安全，采用redis
func (a *UsersApi) UserBindEmailApi(c *gin.Context) {

	claims, err := claimx.GetClaim(c)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c, err)
		return
	}

	var bindEmailReq req.UserBindEmailReq
	err = c.ShouldBind(&bindEmailReq)
	if err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	var userSrv = service.Srv.NewUserSrv(c)
	//通过判断是否传了验证码来决定是第几步操作
	//验证码设置为引用类型，排除用户提交""字符串
	if bindEmailReq.Code == nil {
		//第一步，发送验证码
		code := random.NewCodeSix()
		//go email.NewCode().Send(bindEmailReq.Email, "你的验证码是"+code)
		fmt.Println(code)
		//if err != nil {
		//	global.Log.Error(err)
		//	callback.FAIL(res.EmailSendError, res.CodeMsg(res.EmailSendError), c, err)
		//	return
		//}
		session, err := store.Get(c.Request, "sessionID")
		if err != nil {
			global.Log.Error(err)
			callback.FAIL(res.SessionError, res.CodeMsg(res.SessionError), c, err)
			return
		}
		session.Values["code"] = code
		session.Values["email"] = bindEmailReq.Email
		session.Values["failCount"] = 0
		session.Values["exp"] = time.Now().Add(60 * time.Second).Unix()
		//TODO HTTPS
		session.Options.Secure = false

		err = store.Save(c.Request, c.Writer, session)
		if err != nil {
			global.Log.Error(err)
			callback.FAIL(res.SessionError, res.CodeMsg(res.SessionError), c, err)
			return
		}
		callback.OK(res.EmptyData, c)
		return
	}
	session, err := store.Get(c.Request, "sessionID")
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.NotFoundSession, res.CodeMsg(res.NotFoundSession), c, err)
		return
	}
	//第二步，绑定邮箱
	s, err := userSrv.CheckSession(session)
	if err != nil {
		callback.FAIL(res.SessionError, res.CodeMsg(res.SessionError), c, err)
		return
	}
	global.Log.Debugln(s.Code, s.Email, s.Exp)
	if s.Code != *bindEmailReq.Code || s.Email != bindEmailReq.Email {
		callback.FAIL(res.CodeNotMatched, res.CodeMsg(res.CodeNotMatched), c, err)
		return
	}
	if s.Exp < time.Now().Unix() {
		callback.FAIL(res.SessionExpired, res.CodeMsg(res.SessionExpired), c, err)
		return
	}
	var user dao.UserModel
	user.ID = claims.UserID
	user.Email = bindEmailReq.Email
	err = user.Update(global.Db)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	callback.OK(res.EmptyData, c)

}
