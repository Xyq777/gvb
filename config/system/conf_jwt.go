package system

type Jwt struct {
	Secret  string `json:"secret" `  // 密钥
	Expires int    `json:"expires" ` // 过期时间
	Issuer  string `json:"issuer" `  // 颁发人
}
