package custom

type QiNiu struct {
	AccessKey string  `json:"access_key" toml:"access_key"`
	SecretKey string  `json:"secret_key" toml:"secret_key"`
	Bucket    string  `json:"bucket" toml:"bucket"` // 存储桶的名字
	CDN       string  `json:"cdn" toml:"cdn"`       // 访问图片的地址的前缀
	Zone      string  `json:"zone" toml:"zone"`     // 存储的地区
	Size      float64 `json:"size" toml:"size"`     // 存储的大小限制，单位是MB
}
