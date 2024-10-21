package ctype // Package ctype custom type

import "encoding/json"

type SignStatus int

const (
	SignQQ     SignStatus = 1
	SignGithub SignStatus = 2
	SignEmail  SignStatus = 3
)

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "qq"
	case SignGithub:
		str = "github"
	case SignEmail:
		str = "邮箱"

	default:
		str = "其他"
	}
	return str
}
func (s SignStatus) MarshaJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
