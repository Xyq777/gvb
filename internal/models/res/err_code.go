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
)
