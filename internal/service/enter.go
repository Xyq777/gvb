package service

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/service/srv_menu"
	"gvb/internal/service/srv_user"
)

type Srv_ struct {
}

var Srv = Srv_{}

func (s Srv_) NewMenuSrv(c *gin.Context) *srv_menu.MenuSrv {
	return &srv_menu.MenuSrv{C: c}
}
func (s Srv_) NewUserSrv(c *gin.Context) *srv_user.UserSrv {
	return &srv_user.UserSrv{}

}
