package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb/config"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Log    *logrus.Logger
	Redis  *redis.Client
)
