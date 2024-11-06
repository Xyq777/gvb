package srv_user

import (
	"gorm.io/gorm"
)

type UserSrv struct {
	*gorm.DB
}
