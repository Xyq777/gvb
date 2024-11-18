package flag

import (
	"gvb/internal/global"
	"gvb/internal/models/dao"
)

func Makemigrations() {
	var err error
	err = global.Db.SetupJoinTable(&dao.UserModel{}, "CollectsModels", &dao.UserCollectModel{})
	if err != nil {
		global.Log.Error("[ error ] 创建数据库表关联失败")
		return
	}
	err = global.Db.SetupJoinTable(&dao.MenuModel{}, "Banners", &dao.MenuBannerModel{})
	if err != nil {
		global.Log.Error("[ error ] 创建数据库表关联失败")
		return
	}
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
			&dao.FadeBackModel{},
			&dao.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
