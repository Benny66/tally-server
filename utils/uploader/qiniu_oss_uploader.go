package uploader

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-10-12 17:23:30
 * @LastEditors: shahao
 * @LastEditTime: 2022-07-12 12:02:44
 */

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 阿里云oss
type qiniuOssUploader struct {
	once         sync.Once
	formUploader *storage.FormUploader
	Config       *qiniuOssUploaderConfig
}

type qiniuOssUploaderConfig struct {
	accessId     string
	accessSecret string
	styleDetail  string
	Bucket       string
	RegionHost   string
	host         string
}

func (qiniu *qiniuOssUploader) PutImage(data []byte, contentType string) (string, error) {
	if contentType == "" {
		contentType = "image/jpeg"
	}
	key := generateImageKey(data, contentType)
	return qiniu.PutObject(key, data, contentType)
}

func (qiniu *qiniuOssUploader) PutObject(key string, data []byte, contentType string) (string, error) {
	formUploader := qiniu.getFormUploader()
	options := map[string]string{}
	if contentType != "" {
		options["Content-Type"] = contentType
	}
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: options,
	}
	token := qiniu.GetToken(0, key, qiniu.Config.styleDetail)
	err := formUploader.Put(context.Background(), &ret, token, key, bytes.NewReader(data), int64(len(data)), &putExtra)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (qiniu *qiniuOssUploader) CopyImage(originUrl string) (string, error) {
	data, contentType, err := download(originUrl)
	if err != nil {
		return "", err
	}
	return qiniu.PutImage(data, contentType)
}

func (qiniu *qiniuOssUploader) getFormUploader() *storage.FormUploader {
	qiniu.once.Do(func() {
		cfg := storage.Config{}
		// 空间对应的机房
		//cfg.Zone = &storage.ZoneHuadong
		// 是否使用https域名
		cfg.UseHTTPS = true
		// 上传是否使用CDN上传加速
		cfg.UseCdnDomains = false
		qiniu.formUploader = storage.NewFormUploader(&cfg)
	})
	return qiniu.formUploader
}

func (qiniu *qiniuOssUploader) GetToken(expires uint64, key string, persistentOps string) string {
	putPolicy := storage.PutPolicy{
		Scope:   qiniu.Config.Bucket, //覆盖文件
		Expires: expires,
	}
	if persistentOps != "" {
		urlbase64 := storage.EncodedEntry(qiniu.Config.Bucket, key)
		putPolicy.Scope = qiniu.Config.Bucket + ":" + key
		putPolicy.PersistentOps = persistentOps + "|saveas/" + urlbase64
	}
	mac := qbox.NewMac(qiniu.Config.accessId, qiniu.Config.accessSecret)
	return putPolicy.UploadToken(mac)
}

func (qiniu *qiniuOssUploader) GetUrl(key string, style string, deadline int64) string {
	mac := qbox.NewMac(qiniu.Config.accessId, qiniu.Config.accessSecret)
	domain := qiniu.Config.host
	if deadline == -1 {
		//99年不过期url
		deadline = time.Now().AddDate(99, 0, 0).Unix() //24小时有效期
	} else if deadline == 0 {
		//默认24小时
		deadline = time.Now().Add(time.Hour * 24).Unix() //24小时有效期
	}
	var privateAccessURL = ""
	if style == "" || style == "/" {
		privateAccessURL = storage.MakePrivateURL(mac, domain, key, deadline)
	} else {
		privateAccessURL = storage.MakePrivateURLv2WithQueryString(mac, domain, key, style, deadline)
	}
	return privateAccessURL
}
