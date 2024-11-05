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
	DatabaseOperateError       ErrorCode = 50005
	DatabaseFailedCreate       ErrorCode = 50006
	DatabaseMenuFailedDelete   ErrorCode = 50007
)

var codeToMsg = map[ErrorCode]string{
	Success:                  "操作成功",
	NotFound:                 "资源不存在",
	InvalidParams:            "参数错误",
	DatabaseFailedCreate:     "数据库创建操作失败",
	DatabaseOperateError:     "数据库操作失败",
	DatabaseMenuFailedDelete: "数据库菜单删除失败",
}

func CodeMsg(code ErrorCode) string {
	return codeToMsg[code]
}
