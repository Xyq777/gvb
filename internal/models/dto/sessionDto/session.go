package sessionDto

type EmailCodeSession struct {
	Code      string
	Email     string
	Exp       int64
	FailCount int
}
