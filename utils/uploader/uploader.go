package uploader

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-10-12 17:23:30
 * @LastEditors: shahao
 * @LastEditTime: 2022-07-12 12:02:00
 */

import (
	"strings"
	"sync"

	"github.com/Benny66/tally-server/config"
	"github.com/Benny66/tally-server/utils/function"
)

type uploader interface {
	PutImage(data []byte, contentType string) (string, error)
	PutObject(key string, data []byte, contentType string) (string, error)
	CopyImage(originUrl string) (string, error)
	GetToken(expires uint64, key string, persistentOps string) string
	GetUrl(key string, style string, deadline int64) string
}

var (
	local = &localUploader{
		host: config.Config.AppUrl,
		path: config.Config.AppLocalPath,
	}
	Qiniu = &qiniuOssUploader{
		once:         sync.Once{},
		formUploader: nil,
		Config: &qiniuOssUploaderConfig{
			host:         config.Config.QiNiuHost,
			RegionHost:   config.Config.QiNiuRegionHost,
			accessId:     config.Config.QiNiuAccessId,
			accessSecret: config.Config.QiNiuAccessSecret,
			Bucket:       config.Config.QiNiuBucket,
		},
	}
)

func PutImage(data []byte, contentType string) (string, error) {
	return getUploader().PutImage(data, contentType)
}

func PutObject(key string, data []byte, contentType string) (string, error) {
	return getUploader().PutObject(key, data, contentType)
}

func CopyImage(url string) (string, error) {
	u1 := function.ParseUrl(url).GetURL()
	u2 := function.ParseUrl(config.Config.AppUrl).GetURL()
	// 本站host，不下载
	if u1.Host == u2.Host {
		return url, nil
	}
	return getUploader().CopyImage(url)
}

func GetToken(expires uint64, key string, persistentOps string) string {
	return getUploader().GetToken(expires, key, persistentOps)
}
func GetUrl(key string, style string, deadline int64) string {
	return getUploader().GetUrl(key, style, deadline)
}

func getUploader() uploader {
	switch config.Config.AppUploadType {
	case UPLOADER_OSS_QINIU:
		return Qiniu
	}
	return local
}

const (
	UPLOADER_OSS_QINIU = "qiniu"
	UPLOADER_LOCAL     = "local"
)

// IsEnabledOss 是否启用oss，默认七牛云
func IsEnabledOss() bool {
	switch config.Config.AppUploadType {
	case UPLOADER_OSS_QINIU:
		return true
	}
	return false
}

// IsOssImageUrl 是否是存放在阿里云oss中的图片
func IsOssImageUrl(url string) bool {
	switch config.Config.AppUploadType {
	case UPLOADER_OSS_QINIU:
		host := function.ParseUrl(config.Config.QiNiuHost).GetURL().Host
		return strings.Contains(url, host)
	}
	return false

}
