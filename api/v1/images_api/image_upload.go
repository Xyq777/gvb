package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/ctype"
	"gvb/internal/models/res"
	"gvb/internal/tools/qiniu"
	"gvb/tools/Encryptor"
	"gvb/tools/validator"
	"io"
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

var hasFail = false

// ImagesUploadAPI 上传图片，返回图片url
func (a ImagesApi) ImagesUploadAPI(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		res.FAIL(res.InvalidParams, "参数错误", c, err)
		return
	}
	files, ok1 := form.File["images"]
	uploadType, ok2 := form.Value["type"]
	if !ok1 || !ok2 {
		res.FAIL(res.InvalidParams, "请求字段错误", c, err)
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
			appendFailedResList(&resList, fileHeader.Filename, "图片格式不正确")
			continue
		}

		//检查大小
		size := float64(fileHeader.Size) / float64(1*MB)
		if size >= float64(global.Config.System.Upload.Size) {
			appendFailedResList(&resList, fileHeader.Filename,
				fmt.Sprintf("图片大小超过%.2fMB,设定为%dMB", size, global.Config.System.Upload.Size),
			)
			continue
		}
		//判断图片是否入库
		fileObj, err := fileHeader.Open()
		if err != nil {
			global.Log.Errorln(err)
			appendFailedResList(&resList, fileHeader.Filename, "图片打开错误")
			continue
		}
		//检查图片是否已经存在
		fileByte, err := io.ReadAll(fileObj)
		if err != nil {
			global.Log.Errorln(err)
			appendFailedResList(&resList, fileHeader.Filename, "图片读取错误")
			continue
		}
		fileHash := Encryptor.Md5(fileByte)
		result := global.Db.Limit(1).Find(&models.BannerModel{}, "hash = ?", fileHash)
		if result.Error != nil {
			appendFailedResList(&resList, fileHeader.Filename, "数据库查询错误")
			continue
		}
		if result.RowsAffected != 0 {
			appendFailedResList(&resList, fileHeader.Filename, "图片已存在")
			continue
		}

		//上传
		//上传到七牛云
		var filePath string
		var modelImgaeType ctype.ImageType
		switch uploadType[0] {
		case "qiniu":
			modelImgaeType = 2
			filePath, err = qiniu.UploadImage(fileByte, fileHeader.Filename, "gvb")
			if err != nil {
				appendFailedResList(&resList, fileHeader.Filename, "上传七牛失败")
				continue
			}
		case "local":
			modelImgaeType = 1
			filePath = path.Join(basePath, fileHeader.Filename)
			err = c.SaveUploadedFile(fileHeader, filePath)
			if err != nil {
				appendFailedResList(&resList, fileHeader.Filename, "图片保存错误")
				global.Log.Errorln(err)
				continue
			}
		}
		//入库
		err = global.Db.Create(&models.BannerModel{
			Path:      filePath,
			Hash:      fileHash,
			Name:      fileHeader.Filename,
			ImageType: modelImgaeType,
		}).Error
		if err != nil {
			appendFailedResList(&resList, fileHeader.Filename, "图片入库错误")
			continue
		}
		appendSuccessResList(&resList, filePath, "图片上传成功")

	}

	if hasFail {
		res.FAIL(res.InvalidParams, "图片上传错误", c, resList)
		return
	}
	res.OK(resList, c)
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
