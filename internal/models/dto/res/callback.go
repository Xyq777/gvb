package res

type Response struct {
	Code ErrorCode `json:"code"`
	Data any       `json:"data"`
	Msg  string    `json:"msg"`
}
type List struct {
	ModelList any `json:"model_list"`
	Count     int `json:"count"`
}

var EmptyData = struct{}{}

func NewResponse(code ErrorCode, data any, msg string) *Response {
	return &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}
