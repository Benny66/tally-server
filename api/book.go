package api

import (
	"strconv"

	database "github.com/Benny66/tally-server/db"
	"github.com/Benny66/tally-server/models"
	"github.com/Benny66/tally-server/schemas"
	"github.com/Benny66/tally-server/service"
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/language"
	"github.com/gin-gonic/gin"
)

var BookApi *bookApi

func init() {
	BookApi = NewBookApi()
}

func NewBookApi() *bookApi {
	return &bookApi{}
}

type bookApi struct {
}

func (api *bookApi) GetUserBooks(context *gin.Context) {
	userInfo := service.UserService.User(context)

	books, err := models.NewBookDao().FindAllWhere("user_id", userInfo.ID)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	format.NewResponseJson(context).Success(books)
}

func (api *bookApi) GetUserBookInfo(context *gin.Context) {
	idStr := context.Request.PostFormValue("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}

	userInfo := service.UserService.User(context)

	book, err := models.NewBookDao().FindOneWhere("user_id = ? and id = ?", userInfo.ID, id)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	format.NewResponseJson(context).Success(book)
}
func (api *bookApi) SetBookEdit(context *gin.Context) {
	userInfo := service.UserService.User(context)

	var req = schemas.SetBookEditApiReq{
		Name: context.Request.FormValue("name"),
	}
	sortStr := context.Request.FormValue("sort")
	req.Sort, _ = strconv.Atoi(sortStr)
	if req.Name == "" {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}
	idStr := context.Request.FormValue("id")
	req.Id, _ = strconv.Atoi(idStr)
	tx := database.Orm.DB()
	if req.Id == 0 {
		book := models.BookModel{
			UserId: userInfo.ID,
			Name:   req.Name,
			Sort:   req.Sort,
		}
		_, err := models.BookDao.Create(tx, &book)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	} else {
		_, err := models.BookDao.Update(tx, uint(req.Id), map[string]interface{}{
			"name": req.Name,
			"sort": req.Sort,
		})
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	}
	format.NewResponseJson(context).Success("")
}
