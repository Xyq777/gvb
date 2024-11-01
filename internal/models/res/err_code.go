package res

type ErrorCode = int

const (
	Success                    ErrorCode = 20000
	InvalidParams              ErrorCode = 40000
	NotFound                   ErrorCode = 40004
	UploadFileSizeExceedsLimit ErrorCode = 40005
	NotFoundImages             ErrorCode = 40006
	FailedRewriteToml          ErrorCode = 50000
	UploadFileFailed           ErrorCode = 50001
	FailedGetImageList         ErrorCode = 50002
	FailedDeleteImages         ErrorCode = 50003
	FailedCreateDir            ErrorCode = 50004
	DatabaseFailedCreate       ErrorCode = iota
)

var codeToMsg = map[ErrorCode]string{
	InvalidParams:        "参数错误",
	DatabaseFailedCreate: "数据库创建操作失败",
}

func ErrorMsg(code ErrorCode) string {
	return codeToMsg[code]
}
func CodeAndMsg(code ErrorCode) (ErrorCode, string) {
	return code, codeToMsg[code]
}
