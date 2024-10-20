package res

type ErrorCode = int

const (
	Success       ErrorCode = 20000
	InvalidParams ErrorCode = 40000
	NotFound      ErrorCode = 40004
)
