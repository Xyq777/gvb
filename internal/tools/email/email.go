package email

import (
	"fmt"
	_email "github.com/jordan-wright/email"
	"gvb/internal/global"
	"net/smtp"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}
func NewNote() Api {
	return Api{
		Subject: Note,
	}
}
func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 邮件发送  发给谁，主题，正文
func send(name, subject, body string) error {
	e := global.Config.System.Email
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := _email.NewEmail()
	header := sendName + "<" + userName + ">"
	m.From = header         // 发件人邮箱，发件人名字
	m.To = []string{mailTo} // 发送给谁
	m.Subject = subject     // 主题
	m.HTML = []byte(body)   // 正文
	err := m.Send(fmt.Sprintf("%s:%d", host, port), smtp.PlainAuth("", userName, authCode, host))
	return err
}
