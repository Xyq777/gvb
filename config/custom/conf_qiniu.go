package custom

type QiNiu struct {
	Enabled   bool    `json:"enabled" toml:"enabled" binding:"required"`
	AccessKey string  `json:"access_key" toml:"access_key" binding:"required"`
	SecretKey string  `json:"secret_key" toml:"secret_key" binding:"required"`
	Bucket    string  `json:"bucket" toml:"bucket" binding:"required"` // 存储桶的名字
	CDN       string  `json:"cdn" toml:"cdn" binding:"required"`       // 访问图片的地址的前缀
	Zone      string  `json:"zone" toml:"zone" binding:"required"`     // 存储的地区
	Size      float64 `json:"size" toml:"size" binding:"required"`     // 存储的大小限制，单位是MB
}
