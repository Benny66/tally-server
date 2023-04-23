package schemas

import "github.com/Benny66/tally-server/models"

type UserAuthLoginApiReq struct {
	Code string `form:"code" json:"code"`
}

type UserAuthLoginApiRes struct {
	Token string `form:"token" json:"token"`
}

type UserUpdatePasswordApiReq struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserInfo struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
}

type SetCategoryApiReq struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Type int    `form:"type" json:"type"`
}

type SetTallyMainEditApiReq struct {
	Id         int     `form:"id" json:"id"`
	BookId     int     `form:"book_id" json:"book_id"`
	CategoryId int     `form:"category_id" json:"category_id"`
	Type       int     `form:"type" json:"type"`
	Money      float64 `form:"money" json:"money"`
	Name       string  `form:"name" json:"name"`
	Desc       string  `form:"desc" json:"desc"`
	Date       string  `form:"date" json:"date"`
}

type GetTallyMainListApiReq struct {
	BookId     int    `form:"book_id" json:"book_id"`
	CategoryId int    `form:"category_id" json:"category_id"`
	Type       int    `form:"type" json:"type"`
	StartTime  string `form:"start_time" json:"start_time"`
	EndTime    string `form:"end_time" json:"end_time"`
}

type GetTallyMainListApiRes struct {
	ID         uint                 `gorm:"primaryKey;column:id" json:"id"`
	UserId     uint                 `gorm:"column:user_id;not null" json:"user_id"`
	BookId     uint                 `gorm:"column:book_id;not null" json:"book_id"`
	CategoryId uint                 `gorm:"column:category_id;not null" json:"category_id"`
	MType      int                  `gorm:"column:type;not null" json:"type"`
	Money      float64              `gorm:"column:money;not null" json:"money"`
	Name       string               `gorm:"column:name;not null" json:"name"`
	Desc       string               `gorm:"column:desc;not null" json:"desc"`
	Date       string               `gorm:"column:date" json:"date"`
	IsDel      int                  `gorm:"column:is_del;not null" json:"is_del"`
	CreatedAt  string               `gorm:"column:created_at" json:"created_at"`
	Book       models.BookModel     `gorm:"column:book_info" json:"book_info"`
	Category   models.CategoryModel `gorm:"column:category_info" json:"category_info"`
}
type GetTallyMainTotalApiRes struct {
	Expend    float64 `gorm:"column:expend;not null" json:"expend"`
	Income    float64 `gorm:"column:income;not null" json:"income"`
	Disregard float64 `gorm:"column:disregard;not null" json:"disregard"`
}

type SetBookEditApiReq struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Sort int    `form:"sort" json:"sort"`
}

type SetUserInfoApiReq struct {
	Nickname  string `form:"nick_name" json:"nick_name"`
	Sex       int    `form:"sex" json:"sex"`
	Job       string `form:"job" json:"job"`
	AvatarUrl string `form:"avatar_url" json:"avatar_url"`
}
