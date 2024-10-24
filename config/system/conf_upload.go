package system

type Upload struct {
	Size int    `toml:"size"` //上传图片大小 MB
	Path string `toml:"path"` //上传文件路径
}
