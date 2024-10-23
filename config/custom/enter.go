package custom

type Custom struct {
	SiteInfo SiteInfo `toml:"site_info"`
	QQ       QQ       `toml:"qq"`
	QiNiu    QiNiu    `toml:"qi_niu"`
	Email    Email    `toml:"email"`
}
