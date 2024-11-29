package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dao"
	"gvb/internal/models/dto/req"
	"gvb/internal/models/dto/res"
	"gvb/internal/tools/claimx"
	"gvb/tools/validator"
	"math/rand"
	"time"
)

func (a ArticleApi) CreateArticleApi(c *gin.Context) {
	var articleCreateReq req.ArticleReq
	if err := c.ShouldBindJSON(&articleCreateReq); err != nil {
		callback.FAIL(res.InvalidParams, res.CodeMsg(res.InvalidParams), c, err)
		return
	}
	claims, err := claimx.GetClaim(c)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.AuthFailed, res.CodeMsg(res.AuthFailed), c, err)
		return
	}
	articleCreateReq.Content, err = validator.FilterScriptTag(articleCreateReq.Content)
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.MarkdownTransferFailed, res.CodeMsg(res.MarkdownTransferFailed), c, err)
		return
	}

	//自动提取摘要
	if articleCreateReq.Brief == "" {
		//string按字节切分
		brief := []rune(articleCreateReq.Content)
		articleCreateReq.Brief = string(brief[:100])
	}
	// 不传banner_id,后台就随机去选择一张
	if articleCreateReq.BannerID == 0 {
		var bannerIDList []uint
		global.Db.Model(dao.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			callback.FAIL(res.BannerNotExist, res.CodeMsg(res.BannerNotExist), c)
			return
		}
		articleCreateReq.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}
	//banner
	var bannerUrl string
	bannerModel, exist, err := dao.FindWithID(dao.BannerModel{}, articleCreateReq.BannerID)
	if err != nil {
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	if !exist {
		callback.FAIL(res.BannerNotExist, res.CodeMsg(res.BannerNotExist), c, err)
		return
	}
	bannerUrl = bannerModel.Path
	//avatar
	var userAvatar string
	userModel, exist, err := dao.FindWithID(dao.UserModel{}, claims.UserID)
	if err != nil {
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	if !exist {
		callback.FAIL(res.UserNotExist, res.CodeMsg(res.UserNotExist), c, err)
		return
	}
	userAvatar = userModel.Avatar

	//new model
	articleModel := &dao.ArticleModel{
		Title:        articleCreateReq.Title,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Brief:        articleCreateReq.Brief,
		Content:      articleCreateReq.Content,
		UserID:       claims.UserID,
		UserNickName: claims.Nickname,
		UserAvatar:   userAvatar,
		Category:     articleCreateReq.Category,
		Source:       articleCreateReq.Source,
		Link:         articleCreateReq.Link,
		BannerID:     articleCreateReq.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         articleCreateReq.Tags,
	}

	err = articleModel.CreateInES()
	if err != nil {
		global.Log.Error(err)
		callback.FAIL(res.DatabaseOperateError, res.CodeMsg(res.DatabaseOperateError), c, err)
		return
	}
	callback.OK(nil, c)

}
