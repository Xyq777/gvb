package routers

import (
	v1 "gvb/api/v1"
	"gvb/internal/middleware"
)

func (r *RouterGroup) ArticleRouter() {
	g := r.Group("article")
	g.Use(middleware.JwtAuth())
	articleApi := v1.ApiGroupApp.ArticleApi
	g.POST("", articleApi.CreateArticleApi)
}