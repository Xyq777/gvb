package res

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code ErrorCode, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func OK(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}
func FailWithData(code ErrorCode, msg string, err error, c *gin.Context) {

}
func FAIL(code ErrorCode, msg string, c *gin.Context, option ...any) {
	if len(option) == 0 {
		Result(code, struct{}{}, msg, c)
		return
	}
	if len(option) == 1 {
		f, ok := option[0].(error)
		if isDev() && ok {
			FailWithOrigin(code, msg, f, c)
			return
		}
		if !ok {
			Result(code, option[0], msg, c)
			return
		}
	}
	if len(option) == 2 {
		f0, ok0 := option[0].(error)
		f1, ok1 := option[1].(error)
		if ok0 {
			if isDev() {
				Result(code, f1, msg+f0.Error(), c)
			} else {
				Result(code, f1, msg, c)
			}
			return
		}
		if ok1 {
			if isDev() {
				Result(code, f0, msg+f1.Error(), c)
			} else {
				Result(code, f0, msg, c)
			}
			return
		}
		if ok0 && ok1 {
			global.Log.Warnln("FAIL函数参数错误")
		}
	}
	global.Log.Warnln("FAIL函数参数错误")
}
func FailWithOrigin(code ErrorCode, msg string, err error, c *gin.Context) {
	Result(code, struct{}{}, msg+err.Error(), c)
}
func isDev() bool {
	return global.Config.System.App.LogLevel == "dev"
}
