package models

import (
	database "github.com/Benny66/tally-server/db"

	"gorm.io/gorm"
)

var MainDao *mainDao

func init() {
	MainDao = NewMainDao()
}

func NewMainDao() *mainDao {
	return &mainDao{
		gm: database.Orm.DB(),
	}
}

type mainDao struct {
	gm *gorm.DB
}

func (dao *mainDao) Create(tx *gorm.DB, data *MainModel) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *mainDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&MainModel{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *mainDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&MainModel{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *mainDao) FindAll() (list []MainModel, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *mainDao) FindAllWhere(query interface{}, args ...interface{}) (list []MainModel, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *mainDao) FindOneWhere(query interface{}, args ...interface{}) (record MainModel, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *mainDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&MainModel{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *mainDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&MainModel{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *mainDao) Raw(sqlStr string, params ...interface{}) (list []MainModel, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}
func (dao *mainDao) WhereQuery(query interface{}, args ...interface{}) *mainDao {
	return &mainDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *mainDao) WhereUserNameLike(username string) *mainDao {
	return &mainDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *mainDao) WhereDisabled(isDisabled int) *mainDao {
	return &mainDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

func (dao *mainDao) Paginate(offset, limit int) (count int64, list []MainModel, err error) {
	db := dao.gm.Model(&MainModel{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *mainDao) Debug() *mainDao {
	return &mainDao{
		dao.gm.Debug(),
	}
}

func (dao *mainDao) Offset(offset int) *mainDao {
	return &mainDao{
		dao.gm.Offset(offset),
	}
}

func (dao *mainDao) Limit(limit int) *mainDao {
	return &mainDao{
		dao.gm.Limit(limit),
	}
}

func (dao *mainDao) OrderBy(sortFlag, sortOrder string) *mainDao {
	return &mainDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

func (dao *mainDao) Joins(query string, args ...interface{}) *mainDao {
	return &mainDao{
		dao.gm.Joins(query, args),
	}
}

func (dao *mainDao) Preloads(query string) *mainDao {
	return &mainDao{
		dao.gm.Preload(query),
	}
}
