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
func FAIL(code ErrorCode, msg string, err error, c *gin.Context) {
	if global.Config.System.App.LogLevel == "dev" {
		FailWithOrigin(code, msg, err, c)
		return
	}
	Result(code, struct{}{}, msg, c)
}
func FailWithOrigin(code ErrorCode, msg string, err error, c *gin.Context) {
	Result(code, struct{}{}, msg+err.Error(), c)
}
