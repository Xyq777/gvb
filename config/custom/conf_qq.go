package custom

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" toml:"app_id" binding:"required"`
	Key      string `json:"key" toml:"key" binding:"required"`
	Redirect string `json:"redirect" toml:"redirect" binding:"required"` // 登录之后的回调地址
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
