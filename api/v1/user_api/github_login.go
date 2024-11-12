package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/res"
	"gvb/tools/random"
	"net/http"
	"net/url"
)

func (a *UsersApi) UserGithubLoginApi(c *gin.Context) {
	state, err := random.GenerateString(16)
	if err != nil {
		callback.FAIL(res.SeverError, res.CodeMsg(res.SeverError), c)
		return
	}

	c.SetCookie(
		"githubState",
		state,
		1200,
		"/api/user/github",
		"127.0.0.1",
		false,
		true,
	)
	redirectUrl, err := url.Parse("https://github.com/login/oauth/authorize")
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.SeverError, res.CodeMsg(res.SeverError), c, err)
		return
	}
	query := redirectUrl.Query()
	query.Set("state", state)
	query.Set("client_id", "Ov23liu23ffMYOdtPjCS")
	redirectUrl.RawQuery = query.Encode()
	fmt.Println(redirectUrl.String())
	callback.Redirect(http.StatusTemporaryRedirect, c, redirectUrl.String())

}
