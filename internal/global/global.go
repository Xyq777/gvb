package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb/config"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Log    *logrus.Logger
)
