package image_api

import (
	"github.com/gin-gonic/gin"
	"gvb/internal/callback"
	"gvb/internal/global"
	"gvb/internal/models/dto/res"
	"gvb/internal/service/srv_image"
	"io/fs"
	"os"
)

type FileUploadResponse struct {
	FileName  string
	IsSuccess bool
	Msg       string
}

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		".jpg",
		".png",
		".jpeg",
		".ico",
		".tiff",
		".gif",
		".svg",
		".webp",
	}
)

const KB = 1024
const MB = KB * KB

var hasFail = false

// ImagesUploadAPI 上传图片，返回图片url
func (a ImagesApi) ImagesUploadAPI(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		callback.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	files, ok1 := form.File["images"]
	uploadType, ok2 := form.Value["type"]
	if !ok1 || !ok2 {
		callback.FAIL(res.InvalidParams, "请求字段错误", c, err)
		return
	}
	//判断文件路径是否存在
	basePath := global.Config.System.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
			callback.FAIL(res.FailedCreateDir, "创建目录失败", c, err)
		}
	}
	//封装
	var resList []FileUploadResponse
	var imageSrv srv_image.ImageSrv
	for _, fileHeader := range files {
		err = imageSrv.ImageUploadService(c, fileHeader, uploadType[0])
		if err != nil {
			hasFail = true
			appendFailedResList(&resList, fileHeader.Filename, err.Error())
			continue
		}
		appendSuccessResList(&resList, fileHeader.Filename, "上传成功")
	}

	if hasFail {
		callback.FAIL(res.InvalidParams, "图片上传错误", c, resList)
		return
	}
	callback.OK(resList, c)
}
func appendFailedResList(resList *[]FileUploadResponse, fileName string, msg string) {
	hasFail = true
	*resList = append(*resList, FileUploadResponse{
		FileName:  fileName,
		IsSuccess: false,
		Msg:       msg,
	})
}
func appendSuccessResList(resList *[]FileUploadResponse, fileName string, msg string) {
	*resList = append(*resList, FileUploadResponse{
		FileName:  fileName,
		IsSuccess: true,
		Msg:       msg,
	})
}
