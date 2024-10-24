package custom

type Email struct {
	Host             string `json:"host" toml:"host" binding:"required"`
	Port             int    `json:"port" toml:"port" binding:"required"`
	User             string `json:"user" toml:"user" binding:"required"` // 发件人邮箱
	Password         string `json:"password" toml:"password" binding:"required"`
	DefaultFromEmail string `json:"default_from_email" toml:"default_from_email" binding:"required"` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" toml:"use_ssl" binding:"required"`                       // 是否使用ssl
	UserTls          bool   `json:"user_tls" toml:"user_tls" binding:"required"`                     //
}
