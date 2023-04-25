package api

import (
	"fmt"
	"strconv"
	"time"

	database "github.com/Benny66/tally-server/db"
	"github.com/Benny66/tally-server/models"
	"github.com/Benny66/tally-server/schemas"
	"github.com/Benny66/tally-server/service"
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/function"
	"github.com/Benny66/tally-server/utils/language"
	"github.com/gin-gonic/gin"
)

var TallyApi *tallyApi

func init() {
	TallyApi = NewTallyApi()
}

func NewTallyApi() *tallyApi {
	return &tallyApi{}
}

type tallyApi struct {
}

func (api *tallyApi) GetTallyMainTotal(context *gin.Context) {
	userInfo := service.UserService.User(context)
	var req = schemas.GetTallyMainListApiReq{
		Date: context.Request.FormValue("date"),
	}

	bookIdStr := context.Request.FormValue("book_id")
	req.BookId, _ = strconv.Atoi(bookIdStr)
	var query string = fmt.Sprintf("user_id = %d ", userInfo.ID)
	if req.BookId != 0 {
		query += fmt.Sprintf(" and book_id = %d ", req.BookId)
	}
	if req.Date != "" {
		dateStartTime, err := function.Parse(fmt.Sprintf("%s 00:00:00", req.Date), function.FmtDateTime)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
		dateStartTime = time.Date(dateStartTime.Year(), dateStartTime.Month(), 1, 0, 0, 0, 0, dateStartTime.Location())
		dateEndTime := dateStartTime.AddDate(0, 1, 0)
		query += fmt.Sprintf("and date >= '%s' and date < '%s'", dateStartTime.String(), dateEndTime.String())
	}
	if req.Date == "" && req.StartTime != "" && req.EndTime != "" {
		req.StartTime = fmt.Sprintf("%s 00:00:00", req.StartTime)
		req.EndTime = fmt.Sprintf("%s 23:59:59", req.EndTime)
		query += fmt.Sprintf("and date >= '%s' and date <= '%s'", req.StartTime, req.EndTime)
	}
	tallies, err := models.MainDao.FindAllWhere(query)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	var res = schemas.GetTallyMainTotalApiRes{}
	for _, v := range tallies {
		switch v.MType {
		case 1:
			res.Expend += v.Money
		case 2:
			res.Income += v.Money
		case 3:
			res.Disregard += v.Money
		}
	}
	format.NewResponseJson(context).Success(res)
}
func (api *tallyApi) GetTallyMainList(context *gin.Context) {
	userInfo := service.UserService.User(context)
	var req = schemas.GetTallyMainListApiReq{}

	bookIdStr := context.Request.FormValue("book_id")
	req.BookId, _ = strconv.Atoi(bookIdStr)
	categoryIdStr := context.Request.FormValue("category_id")
	req.CategoryId, _ = strconv.Atoi(categoryIdStr)
	typeStr := context.Request.FormValue("type")
	req.Type, _ = strconv.Atoi(typeStr)

	req.StartTime = context.Request.PostFormValue("start_time")
	req.EndTime = context.Request.PostFormValue("end_time")
	var query string = fmt.Sprintf("user_id = %d ", userInfo.ID)
	if req.BookId != 0 {
		query += fmt.Sprintf(" and book_id = %d ", req.BookId)
	}
	if req.CategoryId != 0 {
		query += fmt.Sprintf("and category_id = %d ", req.CategoryId)
	}
	if req.Type != 0 {
		query += fmt.Sprintf("and type = %d ", req.Type)
	}
	if req.StartTime != "" && req.EndTime != "" {
		req.StartTime = fmt.Sprintf("%s 00:00:00", req.StartTime)
		req.EndTime = fmt.Sprintf("%s 23:59:59", req.EndTime)
		query += fmt.Sprintf("and date >= '%s' and date <= '%s'", req.StartTime, req.EndTime)
	}

	tallies, err := models.MainDao.FindAllWhere(query)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	bookIds := []uint{}
	bookIdMap := make(map[uint]bool, 0)
	categoryIds := []uint{}
	categoryIdMap := make(map[uint]bool, 0)
	for _, tally := range tallies {
		if _, ok := bookIdMap[tally.BookId]; !ok {
			bookIds = append(bookIds, tally.BookId)
		}
		if _, ok := categoryIdMap[tally.BookId]; !ok {
			categoryIds = append(categoryIds, tally.CategoryId)
		}
	}
	books, err := models.BookDao.FindAllWhere("id in (?)", bookIds)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	bookMap := make(map[uint]models.BookModel, 0)

	for _, book := range books {
		bookMap[book.ID] = book
	}
	categories, err := models.CategoryDao.FindAllWhere("id in (?)", categoryIds)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	categoryMap := make(map[uint]models.CategoryModel, 0)
	for _, category := range categories {
		categoryMap[category.ID] = category
	}
	var tallyList []schemas.GetTallyMainListApiRes
	for _, tally := range tallies {
		tallyRes := schemas.GetTallyMainListApiRes{
			ID:         tally.ID,
			UserId:     tally.UserId,
			BookId:     tally.BookId,
			CategoryId: tally.CategoryId,
			MType:      tally.MType,
			Money:      tally.Money,
			Name:       tally.Name,
			Desc:       tally.Desc,
			Date:       tally.Date.String(),
			IsDel:      tally.IsDel,
			CreatedAt:  tally.CreatedAt.String(),
		}
		tallyRes.Book = bookMap[tallyRes.BookId]
		tallyRes.Category = categoryMap[tallyRes.CategoryId]
		tallyList = append(tallyList, tallyRes)
	}
	format.NewResponseJson(context).Success(tallyList)
}

func (api *tallyApi) SetTallyMainEdit(context *gin.Context) {
	userInfo := service.UserService.User(context)

	var req = schemas.SetTallyMainEditApiReq{
		Name: context.Request.FormValue("name"),
		Desc: context.Request.FormValue("desc"),
		Date: context.Request.FormValue("date"),
	}
	idStr := context.Request.FormValue("id")
	req.Id, _ = strconv.Atoi(idStr)
	bookIdStr := context.Request.FormValue("book_id")
	req.BookId, _ = strconv.Atoi(bookIdStr)
	categoryIdStr := context.Request.FormValue("category_id")
	req.CategoryId, _ = strconv.Atoi(categoryIdStr)
	typeStr := context.Request.FormValue("type")
	req.Type, _ = strconv.Atoi(typeStr)
	moneySrr := context.Request.FormValue("money")
	req.Money, _ = strconv.ParseFloat(moneySrr, 64)
	if req.BookId == 0 || req.CategoryId == 0 || req.Type == 0 || req.Money == 0 {
		format.NewResponseJson(context).Error(language.INVALID_PARMAS)
		return
	}
	tx := database.Orm.DB()
	var main models.MainModel
	if req.Id == 0 {
		main = models.MainModel{
			UserId:     userInfo.ID,
			Name:       req.Name,
			BookId:     uint(req.BookId),
			CategoryId: uint(req.CategoryId),
			MType:      req.Type,
			Money:      req.Money,
			Desc:       req.Desc,
		}
		dateTime, err := function.Parse(req.Date, function.FmtDate)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
		main.Date = models.ModelTime(dateTime)
		_, err = models.MainDao.Create(tx, &main)
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	} else {
		_, err := models.MainDao.Update(tx, uint(req.Id), map[string]interface{}{
			"name":  req.Name,
			"type":  req.Type,
			"money": req.Money,
			"desc":  req.Desc,
			"date":  req.Date,
		})
		if err != nil {
			format.NewResponseJson(context).Error(1, err.Error())
			return
		}
	}
	format.NewResponseJson(context).Success(main)

}

func (api *tallyApi) GetTallyMainInfo(context *gin.Context) {
	userInfo := service.UserService.User(context)
	id := context.Request.PostFormValue("id")
	tally, err := models.MainDao.FindOneWhere("user_id = ? and  id = ? ", userInfo.ID, id)
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	format.NewResponseJson(context).Success(tally)
}
