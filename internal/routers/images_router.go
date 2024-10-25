package routers

import v1 "gvb/api/v1"

func (r *RouterGroup) ImagesRouter() {
	group := r.Group("/images")
	imagesApi := v1.ApiGroupApp.ImagesApi
	{
		group.POST("", imagesApi.ImagesUploadAPI)
		group.POST("/list", imagesApi.ImageListApi)
	}

}
