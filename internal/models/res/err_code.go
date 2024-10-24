package res

type ErrorCode = int

const (
	Success                    ErrorCode = 20000
	InvalidParams              ErrorCode = 40000
	NotFound                   ErrorCode = 40004
	UploadFileSizeExceedsLimit ErrorCode = 40005
	FailedRewriteToml          ErrorCode = 50000
	UploadFileFailed           ErrorCode = 50001
)
