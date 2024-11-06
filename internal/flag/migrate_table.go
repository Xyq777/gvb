package flag

import (
	"gvb/internal/global"
	"gvb/internal/models/dao"
)

func Makemigrations() {
	var err error
	global.Db.SetupJoinTable(&dao.UserModel{}, "CollectsModels", &dao.UserCollectModel{})
	global.Db.SetupJoinTable(&dao.MenuModel{}, "Banners", &dao.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.Db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&dao.BannerModel{},
			&dao.TagModel{},
			&dao.MessageModel{},
			&dao.UserModel{},
			&dao.CommentModel{},
			&dao.ArticleModel{},
			&dao.MenuModel{},
			&dao.MenuBannerModel{},
			&dao.FadeBackModel{},
			&dao.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
