package system

type Email struct {
	Host             string `json:"host" `
	Port             int    `json:"port" `
	User             string `json:"user" ` // 发件人邮箱
	Password         string `json:"password" `
	DefaultFromEmail string `json:"default_from_email" ` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" `            // 是否使用ssl
	UserTLS          bool   `json:"user_tls" `           //
}
