package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/objects"
	"gvb/config/custom"
	"gvb/internal/global"
)

// 获取上传的token
func getToken(q custom.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
func getBucket(q custom.QiNiu) *objects.Bucket {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	mac := credentials.NewCredentials(accessKey, secretKey)
	objectsManager := objects.NewObjectsManager(&objects.ObjectsManagerOptions{
		Options: http_client.Options{Credentials: mac},
	})
	bucketName := "gvb-image"
	bucket := objectsManager.Bucket(bucketName)
	return bucket
}

// 获取上传的配置
func getCfg(q custom.QiNiu) storage.Config {
	cfg := storage.Config{}
	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return cfg

}
func DeleteImage(key string) error {
	bucket := getBucket(global.Config.Custom.QiNiu)
	err := bucket.Object(key).Delete().Call(context.Background())
	if err != nil {
		global.Log.Error(err)
		return err
	}
	return nil
}

// UploadImage 上传图片  文件数组，前缀
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.Custom.QiNiu.Enabled {
		return "", errors.New("请启用七牛云上传")
	}
	q := global.Config.Custom.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey及secretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过设定大小")
	}
	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	// 获取当前时间

	key := fmt.Sprintf("%s/%s", prefix, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil

}
