package api

import (
	"fmt"
	"strconv"

	"github.com/Benny66/tally-server/config"
	database "github.com/Benny66/tally-server/db"
	"github.com/Benny66/tally-server/models"
	"github.com/Benny66/tally-server/schemas"
	"github.com/Benny66/tally-server/service"
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/language"
	"github.com/gin-gonic/gin"
)

var CategoryApi *categoryApi

func init() {
	CategoryApi = NewCategoryApi()
}

func NewCategoryApi() *categoryApi {
	return &categoryApi{}
}

type categoryApi struct {
}

func (api *categoryApi) GetCategoryList(context *gin.Context) {
	userInfo := service.UserService.User(context)
	typeStr := context.Request.PostFormValue("type")
	// type, _ := strconv.Atoi(typeStr)
	categories, err := models.NewCategoryDao().FindAllWhere("type = ? and is_del = ? and (user_id = 0 or user_id = ?)", typeStr, 0, userInfo.ID)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	format.NewResponseJson(context).Success(categories)
}

func (api *categoryApi) SetCategoryEdit(context *gin.Context) {
	userInfo := service.UserService.User(context)

	var req = schemas.SetCategoryApiReq{
		Name: context.Request.FormValue("name"),
	}
	typeStr := context.Request.FormValue("type")
	req.Type, _ = strconv.Atoi(typeStr)
	if req.Name == "" {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}
	idStr := context.Request.FormValue("id")
	req.Id, _ = strconv.Atoi(idStr)
	tx := database.Orm.DB()
	if req.Id == 0 {
		category := models.CategoryModel{
			UserId:  userInfo.ID,
			Name:    req.Name,
			CType:   req.Type,
			IconUrl: fmt.Sprintf("%s%s", config.Config.AppUrl, "/public/static/imgs/mot1.png"),
		}
		_, err := models.CategoryDao.Create(tx, &category)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	} else {
		_, err := models.NewCategoryDao().Update(tx, uint(req.Id), map[string]interface{}{
			"name": req.Name,
			"type": req.Type,
		})
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	}
	format.NewResponseJson(context).Success("")
}
