package service

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/service/srv_menu"
)

type Srv_ struct {
}

var Srv = Srv_{}

func (s Srv_) NewMenuSrv(c *gin.Context) *srv_menu.MenuSrv {
	return &srv_menu.MenuSrv{C: c}
}
