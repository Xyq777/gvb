package srv_user

import (
	"errors"
	"github.com/gorilla/sessions"
	"gvb/internal/models/dto/sessionDto"
)

func (s UserSrv) CheckSession(session *sessions.Session) (*sessionDto.EmailCodeSession, error) {
	sessionCode, codeOk := session.Values["code"].(string)
	sessionEmail, emailOk := session.Values["email"].(string)
	sessionExp, expOk := session.Values["exp"].(int64)
	sessionFailCount, failCountOk := session.Values["failCount"].(int)
	if !codeOk || !emailOk || !expOk || !failCountOk {
		return nil, errors.New("session无效")
	}
	return &sessionDto.EmailCodeSession{
		Code:      sessionCode,
		Email:     sessionEmail,
		Exp:       sessionExp,
		FailCount: sessionFailCount,
	}, nil
}
