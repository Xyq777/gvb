package system

type Jwt struct {
	Secret    string `json:"secret" `    // 密钥
	RTExpires int    `json:"rt_expires"` // refreshToken过期时间
	ATExpires int    `json:"at_expires"` //accessToken过期时间
	Issuer    string `json:"issuer" `    // 颁发人
}
