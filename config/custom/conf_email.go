package custom

type Email struct {
	Host             string `json:"host" toml:"host"`
	Port             int    `json:"port" toml:"port"`
	User             string `json:"user" toml:"user"` // 发件人邮箱
	Password         string `json:"password" toml:"password"`
	DefaultFromEmail string `json:"default_from_email" toml:"default_from_email"` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" toml:"use_ssl"`                       // 是否使用ssl
	UserTls          bool   `json:"user_tls" toml:"user_tls"`                     //
}
