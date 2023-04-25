package api

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-20 16:26:34
 */

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/Benny66/tally-server/config"
	database "github.com/Benny66/tally-server/db"
	"github.com/Benny66/tally-server/models"
	"github.com/Benny66/tally-server/schemas"
	"github.com/Benny66/tally-server/service"
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/function"
	"github.com/Benny66/tally-server/utils/language"
	"github.com/Benny66/tally-server/utils/uploader"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/util"

	"github.com/gin-gonic/gin"
)

var UserApi *userApi

func init() {
	UserApi = NewUserApi()
}

func NewUserApi() *userApi {
	return &userApi{}
}

type userApi struct {
}

func (api *userApi) AuthLogin(context *gin.Context) {
	var req = schemas.UserAuthLoginApiReq{
		Code: context.Request.FormValue("code"),
	}

	if req.Code == "" {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}

	wx := wechat.NewWechat()
	memory := cache.NewMemcache()
	cfg := &miniConfig.Config{
		AppID:     config.Config.WxAPPID,
		AppSecret: config.Config.WxSecret,
		Cache:     memory,
	}
	mini := wx.GetMiniProgram(cfg)
	resCode2Session, err := mini.GetAuth().Code2Session(req.Code)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	token := util.RandomStr(32)
	count, err := models.NewUserDao().FindCountWhere("openid", resCode2Session.OpenID)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	tx := database.Orm.DB()
	if count == 0 {
		var userInfo models.UserModel
		userInfo.OpenId = resCode2Session.OpenID
		userInfo.Token = token
		userInfo.NickName = function.RandomNickname()
		userInfo.AvatarUrl = fmt.Sprintf("%s%s%d%s", config.Config.AppUrl, "/public/static/imgs/txs/", rand.Intn(4)+1, ".png")
		userInfo.Sex = 0
		userInfo.Job = ""
		_, err := models.NewUserDao().Create(tx, &userInfo)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	} else {
		userInfo, err := models.NewUserDao().FindOneWhere("openid", resCode2Session.OpenID)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
		_, err = models.NewUserDao().Update(tx, userInfo.ID, map[string]interface{}{
			"token": token,
		})
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	}
	format.NewResponseJson(context).Success(schemas.UserAuthLoginApiRes{
		Token: token,
	})
	return
}

func (api *userApi) GetUserInfo(context *gin.Context) {
	userInfo := service.UserService.User(context)
	format.NewResponseJson(context).Success(userInfo)
}

func (api *userApi) SetUserInfo(context *gin.Context) {
	userInfo := service.UserService.User(context)

	var req = schemas.SetUserInfoApiReq{
		Nickname:  context.Request.FormValue("nick_name"),
		AvatarUrl: context.Request.FormValue("avatar_url"),
		Job:       context.Request.FormValue("job"),
	}
	sexStr := context.Request.FormValue("sex")
	req.Sex, _ = strconv.Atoi(sexStr)
	tx := database.Orm.DB()

	_, err := models.UserDao.Update(tx, userInfo.ID, map[string]interface{}{
		"nick_name":  req.Nickname,
		"avatar_url": req.AvatarUrl,
		"job":        req.Job,
		"sex":        req.Sex,
	})
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	userInfo.NickName = req.Nickname
	userInfo.AvatarUrl = req.AvatarUrl
	userInfo.Job = req.Job
	userInfo.Sex = req.Sex
	format.NewResponseJson(context).Success(userInfo)
}
func (api *userApi) Benediction(context *gin.Context) {
	id := rand.Intn(5) + 1
	phraseInfo, err := models.NewPhraseDao().FindOneWhere("id", id)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	format.NewResponseJson(context).Success(phraseInfo)
}

func (api *userApi) UploadFile(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	//限制10m最大
	if file.Size > 10*1024*1024 {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}
	src, err := file.Open()
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	defer src.Close()
	// 读取文件内容并转换为 []byte
	bytesData, err := ioutil.ReadAll(src)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	url, err := uploader.PutImage(bytesData, "")
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	format.NewResponseJson(context).Success(url)
}

func (api *userApi) GetWeather(context *gin.Context) {
	format.NewResponseJson(context).Success(uploader.GetUrl("/static/imgs/clear_day.png", "", 0))
}

// @Summary 退出登录
// @Description 退出登录
// @Tags 用户
// @Security ApiKeyAuth
// @accept x-www-form-urlencoded
// @Produce  json
// @Success 200 {string} string {"company":"BL","device_name":"Audio Matrix","result":"0","result_message":"成功","version":"1.0", "db_version":"202103101750","language":"zh-cn","data":""}
// @Router /v1/user/logout [get]
func (api *userApi) Logout(context *gin.Context) {
	format.NewResponseJson(context).Success("")
}

// func (api *userApi) UpdatePassword(context *gin.Context) {
// 	var req schemas.UserUpdatePasswordApiReq

// 	if err := context.BindJSON(&req); err != nil {
// 		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
// 		return
// 	}
// 	userInfo := service.UserService.User(context)
// 	data, err := service.UserService.UpdatePassword(&req, userInfo)
// 	if err != nil {
// 		format.NewResponseJson(context).Error(err.GetErrorCode(), err.GetParams()...)
// 		return
// 	}
// 	format.NewResponseJson(context).Success(data)
// }
