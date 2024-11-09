package res

type ErrorCode = int

const (
	Success                    ErrorCode = 20000
	InvalidParams              ErrorCode = 40000
	AuthFailed                 ErrorCode = 40001
	NotFound                   ErrorCode = 40004
	UploadFileSizeExceedsLimit ErrorCode = 40005
	NotFoundImages             ErrorCode = 40006
	UserNotExist               ErrorCode = 40007
	PasswordNotMatched         ErrorCode = 40008
	PermissionDenied           ErrorCode = 40009
	PasswordNotMatch           ErrorCode = 40010
	AlreadyLogout              ErrorCode = 40011
	NotFoundSession            ErrorCode = 40012
	CodeNotMatched             ErrorCode = 40013
	NotFoundSessionField       ErrorCode = 40014

	FailedRewriteToml        ErrorCode = 50000
	UploadFileFailed         ErrorCode = 50001
	FailedGetImageList       ErrorCode = 50002
	FailedDeleteImages       ErrorCode = 50003
	FailedCreateDir          ErrorCode = 50004
	DatabaseOperateError     ErrorCode = 50005
	DatabaseFailedCreate     ErrorCode = 50006
	DatabaseMenuFailedDelete ErrorCode = 50007
	TokenGenerateFailed      ErrorCode = 50008

	RedisGetFailed ErrorCode = 50009
	RedisDelFailed ErrorCode = 50010
	RedisSetFailed ErrorCode = 50011

	EmailSendError ErrorCode = 50012
	SessionError   ErrorCode = 50013
)

var codeToMsg = map[ErrorCode]string{
	Success:              "操作成功",
	NotFound:             "资源不存在",
	InvalidParams:        "参数错误",
	UserNotExist:         "用户不存在",
	AuthFailed:           "认证失败",
	PasswordNotMatched:   "密码错误",
	AlreadyLogout:        "已经登出",
	NotFoundSession:      "Session不存在",
	CodeNotMatched:       "验证码错误",
	NotFoundSessionField: "Session字段不存在",

	DatabaseFailedCreate:     "数据库创建操作失败",
	DatabaseOperateError:     "数据库操作失败",
	DatabaseMenuFailedDelete: "数据库菜单删除失败",
	TokenGenerateFailed:      "Token生成失败",
	RedisGetFailed:           "Redis获取数据失败",
	RedisDelFailed:           "Redis删除数据失败",
	RedisSetFailed:           "Redis设置数据失败",

	EmailSendError: "邮件发送失败",
	SessionError:   "Session操作失败",
}

func CodeMsg(code ErrorCode) string {
	return codeToMsg[code]
}
