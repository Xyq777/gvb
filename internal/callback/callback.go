package callback

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/internal/models/dto/res"
	"net/http"
)

func Result(code res.ErrorCode, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, res.Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func OK(data any, c *gin.Context) {
	Result(res.Success, data, "操作成功", c)
}
func Redirect(code int, c *gin.Context, url string) {
	c.Redirect(code, url)
}
func FAIL(code res.ErrorCode, msg string, c *gin.Context, option ...any) {
	if len(option) == 0 {
		Result(code, struct{}{}, msg, c)
		return
	}
	if len(option) == 1 {
		f, ok := option[0].(error)
		if ok {
			if isDev() {
				FailWithOrigin(code, msg, f, c)
				return
			}
			Result(code, struct{}{}, msg, c)
			return

		}

		Result(code, option[0], msg, c)
		return

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
			global.Log.Panicf("FAIL函数参数错误")
		}
	}
	global.Log.Panicf("FAIL函数参数错误")
}
func FailWithOrigin(code res.ErrorCode, msg string, err error, c *gin.Context) {
	Result(code, struct{}{}, msg+err.Error(), c)
}
func isDev() bool {
	return global.Config.System.App.LogLevel == "dev"
}
