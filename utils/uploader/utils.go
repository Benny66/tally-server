package uploader

/*
 * @Description:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2022-03-14 19:47:12
 * @LastEditors: shahao
 * @LastEditTime: 2022-03-16 19:45:15
 */

import (
	"crypto/md5"
	"encoding/hex"
	"mime"
	"time"

	"github.com/Benny66/tally-server/utils/function"
	"github.com/go-resty/resty/v2"
)

// generateKey 生成图片Key
func generateImageKey(data []byte, contentType string) string {
	h := md5.New()
	h.Write(data)
	md5 := hex.EncodeToString(h.Sum(nil))
	ext := ""
	if contentType != "" {
		exts, err := mime.ExtensionsByType(contentType)
		if err == nil || len(exts) > 0 {
			ext = exts[0]
		}
	}
	return "images/" + function.Format(time.Now(), "2006/01/02/") + md5 + ext
}

func download(url string) ([]byte, string, error) {
	rsp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, "", err
	}
	return rsp.Body(), rsp.Header().Get("Content-Type"), nil
}
