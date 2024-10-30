package image_srv

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_const "gvb/const"
	"gvb/internal/global"
	"gvb/internal/models"
	"gvb/internal/models/ctype"
	"gvb/internal/tools/qiniu"
	"gvb/tools/Encryptor"
	"gvb/tools/validator"
	"io"
	"mime/multipart"
	"path"
)

var WhiteImageList = []string{
	".jpg",
	".png",
	".jpeg",
	".ico",
	".tiff",
	".gif",
	".svg",
	".webp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadService 文件上传的方法
func (ImageSrv) ImageUploadService(c *gin.Context, fh *multipart.FileHeader, uploadType string) (err error) {

	//上传文件
	//检查文件是否合规
	//检查后缀
	ext := path.Ext(fh.Filename)
	if !validator.InList(ext, WhiteImageList) {
		return errors.New("图片格式不正确")
	}

	//检查大小
	size := float64(fh.Size) / float64(1*_const.MB)
	if size >= float64(global.Config.System.Upload.Size) {
		return errors.New(
			fmt.Sprintf("图片大小超过%.2fMB,设定为%dMB", size, global.Config.System.Upload.Size))
	}
	//判断图片是否入库
	fileObj, err := fh.Open()
	if err != nil {
		global.Log.Errorln(err)
		return errors.New("图片打开错误")
	}
	//检查图片是否已经存在
	fileByte, err := io.ReadAll(fileObj)
	if err != nil {
		global.Log.Errorln(err)
		return errors.New("图片读取错误")
	}
	fileHash := Encryptor.Md5(fileByte)
	result := global.Db.Limit(1).Find(&models.BannerModel{}, "hash = ?", fileHash)
	if result.Error != nil {
		return errors.New("数据库查询错误")
	}
	if result.RowsAffected != 0 {
		return errors.New("图片已存在")
	}

	//上传
	//上传到七牛云
	var filePath string
	var modelImgaeType ctype.ImageType
	switch uploadType {
	case "qiniu":
		modelImgaeType = 2
		filePath, err = qiniu.UploadImage(fileByte, fh.Filename, "gvb")
		if err != nil {
			return errors.New("上传七牛失败")
		}
	case "local":
		modelImgaeType = 1
		filePath = path.Join(global.Config.System.Upload.Path, fh.Filename)
		err = c.SaveUploadedFile(fh, filePath)
		if err != nil {
			return errors.New("图片保存错误")
		}
	}
	//入库
	err = global.Db.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      fileHash,
		Name:      fh.Filename,
		ImageType: modelImgaeType,
	}).Error
	if err != nil {
		return errors.New("图片入库错误")
	}
	return nil
}
