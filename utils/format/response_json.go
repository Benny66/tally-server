/*
* FileName: responseJson.go
* Author: shahao
* CreatedOn: 2019-11-19 13:55
* Description:
 */
package format

import (
	"fmt"
	"net/http"

	"github.com/Benny66/tally-server/utils/language"

	"github.com/gin-gonic/gin"
)

func NewResponseJson(ctx *gin.Context) *responseJson {
	return &responseJson{
		context: ctx,
	}
}

type responseJson struct {
	context *gin.Context
}

/*
* description: 设置响应头
* author: shahao
* created on: 19-11-19 下午3:12
* param param_1:
* param param_2:
* return return_1:
 */
func (r *responseJson) SetHeader(key, value string) *responseJson {
	r.context.Writer.Header().Set(key, value)
	return r
}

/*
* description: 成功返回数据构造
* author: shahao
* created on: 19-11-19 下午2:17
* param data: 返回的数据
* param message: 返回提示信息
* return :
 */
func (r *responseJson) Success(data interface{}) {
	r.context.JSON(http.StatusOK, ResultData{
		Code: language.SUCCESS,
		Msg:  language.Lang.Msg(language.SUCCESS),
		Data: data,
	})
}

/*
* description: 错误返回数据构造
* author: shahao
* created on: 19-11-19 下午2:17
* param data: 返回的数据
* param message: 返回提示信息
* return :
 */
func (r *responseJson) Error(errorCode int, params ...interface{}) {
	result := ResultData{
		Code: errorCode,
		Msg:  language.Lang.Msg(errorCode, params...),
		Data: "",
	}
	r.context.Abort()
	httpCode := http.StatusInternalServerError
	switch errorCode {
	case language.TOKEN_EMPTY:
		httpCode = http.StatusUnauthorized
	}
	r.context.JSON(httpCode, result)
}

/*
* description: 文件下载
* author: shahao
* created on: 19-11-19 下午2:17
* param data: 返回的数据
* param message: 返回提示信息
* return :
 */
func (r *responseJson) Download(filename, path string) {
	r.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	r.SetHeader("Content-Type", "application/octet-stream")
	r.context.File(path)
}

type ResultData struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
