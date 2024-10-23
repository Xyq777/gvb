package custom

type SiteInfo struct {
	CreatedAt   string `toml:"created_at" json:"created_at"`
	BeiAn       string `toml:"bei_an" json:"bei_an"`
	Title       string `toml:"title" json:"title"`
	QqImage     string `toml:"qq_image" json:"qq_image"`
	Version     string `toml:"version" json:"version"`
	Email       string `toml:"email" json:"email"`
	WechatImage string `toml:"wechat_image" json:"wechat_image"`
	Name        string `toml:"name" json:"name"`
	Job         string `toml:"job" json:"job"`
	Addr        string `toml:"addr" json:"addr"`
	Slogan      string `toml:"slogan" json:"slogan"`
	SloganEn    string `toml:"slogan_en" json:"slogan_en"`
	Web         string `toml:"web" json:"web"`
	BilibiliUrl string `toml:"bilibili_url" json:"bilibili_url"`
	GiteeUrl    string `toml:"gitee_url" json:"gitee_url"`
	GithubUrl   string `toml:"github_url" json:"github_url"`
}
