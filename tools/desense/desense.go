package desense

import "strings"

func EmailDesensitization(email string) string {
	emailBlock := strings.Split(email, "@")
	if len(emailBlock) != 2 {
		return ""
	}
	return emailBlock[0][:1] + "****" + emailBlock[1]
}
func TelDesensitization(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
