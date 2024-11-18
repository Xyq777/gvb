package system

import "strconv"

type ES struct {
	Host     string `json:"host" `
	Port     int    `json:"port" `
	Username string `json:"username" `
	Password string `json:"password" `
	UseSSL   bool   `json:"use_ssl" `
}

func (e ES) Addr() string {
	if e.UseSSL {
		return "https://" + e.Host + ":" + strconv.Itoa(e.Port)
	}
	return e.Host + ":" + strconv.Itoa(e.Port)
}
