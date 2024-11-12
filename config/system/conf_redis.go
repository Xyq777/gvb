package system

import "fmt"

type Redis struct {
	IP       string `json:"ip" `
	Port     int    `json:"port" `
	Password string `json:"password" config:"omitempty"`
	PoolSize int    `json:"pool_size" config:"omitempty"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
