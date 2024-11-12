package srv_github

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"gvb/internal/global"
	"gvb/internal/models/dto/githubInfo"
	"io"
)

func GetGithubInfo(code string) (*githubInfo.GithubResp, error) {
	var githubOAuthConfig = &oauth2.Config{
		ClientID:     global.Config.System.Github.ClientID,
		ClientSecret: global.Config.System.Github.ClientSecret,
		RedirectURL:  global.Config.System.Github.RedirectURI,
		Endpoint:     github.Endpoint,
	}
	token, err := githubOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}
	client := githubOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}
	githubResp := new(githubInfo.GithubResp)
	err = json.Unmarshal(respData, githubResp)
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}
	return githubResp, nil

}
