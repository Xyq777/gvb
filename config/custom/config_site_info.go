package custom

type SiteInfo struct {
	CreatedAt   string `toml:"created_at" json:"created_at" binding:"required"`
	BeiAn       string `toml:"bei_an" json:"bei_an" binding:"required"`
	Title       string `toml:"title" json:"title" binding:"required"`
	QqImage     string `toml:"qq_image" json:"qq_image" binding:"required"`
	Version     string `toml:"version" json:"version" binding:"required"`
	Email       string `toml:"email" json:"email" binding:"required"`
	WechatImage string `toml:"wechat_image" json:"wechat_image" binding:"required"`
	Name        string `toml:"name" json:"name" binding:"required"`
	Job         string `toml:"job" json:"job" binding:"required"`
	Addr        string `toml:"addr" json:"addr" binding:"required"`
	Slogan      string `toml:"slogan" json:"slogan" binding:"required"`
	SloganEn    string `toml:"slogan_en" json:"slogan_en" binding:"required"`
	Web         string `toml:"web" json:"web" binding:"required"`
	BilibiliUrl string `toml:"bilibili_url" json:"bilibili_url" binding:"required"`
	GiteeUrl    string `toml:"gitee_url" json:"gitee_url" binding:"required"`
	GithubUrl   string `toml:"github_url" json:"github_url" binding:"required"`
}
