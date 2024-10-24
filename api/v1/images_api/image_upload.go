package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/internal/models/res"
	"gvb/tools/validator"
	"io/fs"
	"os"
	"path"
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

// ImagesUploadAPI 上传图片，返回图片url
func (a ImagesApi) ImagesUploadAPI(c *gin.Context) {
	var hasFail = false
	form, err := c.MultipartForm()
	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	files, ok := form.File["images"]
	if !ok {
		res.FAIL(res.InvalidParams, "未找到images", c, err)
	}

	//判断文件路径是否存在
	basePath := global.Config.System.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	//上传文件
	resList := make([]FileUploadResponse, 0)
	for _, fileHeader := range files {
		//检查文件是否合规
		//检查后缀
		ext := path.Ext(fileHeader.Filename)
		if !validator.InList(ext, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  fileHeader.Filename,
				IsSuccess: false,
				Msg:       "图片格式错误",
			})
			hasFail = true
			continue
		}

		//检查大小
		size := float64(fileHeader.Size) / float64(1*MB)
		if size >= float64(global.Config.System.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  fileHeader.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过%.2fMB,设定为%dMB", size, global.Config.System.Upload.Size),
			})
			hasFail = true
			continue
		}
		//上传
		filePath := path.Join(basePath, fileHeader.Filename)
		err = c.SaveUploadedFile(fileHeader, filePath)
		if err != nil {
			resList = append(resList, FileUploadResponse{
				FileName:  fileHeader.Filename,
				IsSuccess: false,
				Msg:       "上传失败",
			})
			hasFail = true
			global.Log.Errorln(err)
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	if hasFail {
		res.FAIL(res.InvalidParams, "图片上传错误", c, resList)
		return
	}
	res.OK(resList, c)
}
