package ctype // Package ctype custom type

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 // 本地
	QiNiu ImageType = 2 // 七牛云
)

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "其他"
	}
	return str
}

func (s ImageType) MarshaJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
