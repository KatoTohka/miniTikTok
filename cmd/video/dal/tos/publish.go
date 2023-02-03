package tos

import (
	"bytes"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
)

func Upload(data []byte, title string) error {
	accessKey := constants.TosAccessKey
	secretKey := constants.TosSecretKey
	bucket := constants.TosBucket
	//资源管理相关的操作首先构建BucketManager对象
	mac := qbox.NewMac(accessKey, secretKey)
	/*storage.Config 配置*/
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	upToken := putPolicy.UploadToken(mac)
	// 可改成断点。
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{}
	// 构建表单上传的对象
	uploadReader := bytes.NewReader(data)
	// TODO:int64??
	err := formUploader.Put(context.Background(), &ret, upToken, title, uploadReader, int64(len(data)), &putExtra)
	if err != nil {
		log.Println(err)
		return errno.UploadErr

	}
	return nil
	//fmt.Println(ret.Key, ret.Hash)
}
