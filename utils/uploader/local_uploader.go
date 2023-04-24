package uploader

/*
 * @Description:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2022-03-14 20:00:51
 * @LastEditors: shahao
 * @LastEditTime: 2022-03-16 19:45:08
 */

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-12-03 14:30:48
 * @LastEditors: shahao
 * @LastEditTime: 2022-03-14 19:57:45
 */

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Benny66/tally-server/utils/function"
)

// 本地文件系统
type localUploader struct {
	host string
	path string
}

func (local *localUploader) PutImage(data []byte, contentType string) (string, error) {
	key := generateImageKey(data, contentType)
	return local.PutObject(key, data, contentType)
}

func (local *localUploader) PutObject(key string, data []byte, contentType string) (string, error) {
	if err := os.MkdirAll("/", os.ModeDir); err != nil {
		return "", err
	}
	filename := filepath.Join(local.path, key)
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(filename, data, os.ModePerm); err != nil {
		return "", err
	}
	return function.UrlJoin(local.host, key), nil
}

func (local *localUploader) CopyImage(originUrl string) (string, error) {
	data, contentType, err := download(originUrl)
	if err != nil {
		return "", err
	}
	return local.PutImage(data, contentType)
}

func (local *localUploader) GetToken(expires uint64, key string, persistentOps string) string {
	return ""
}

func (local *localUploader) GetUrl(key string, style string, deadline int64) string {
	return local.host + local.path + "/" + key
}
