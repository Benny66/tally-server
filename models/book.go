package models

import (
	database "github.com/Benny66/tally-server/db"

	"gorm.io/gorm"
)

var BookDao *bookDao

func init() {
	BookDao = NewBookDao()
}

func NewBookDao() *bookDao {
	return &bookDao{
		gm: database.Orm.DB(),
	}
}

type bookDao struct {
	gm *gorm.DB
}

func (dao *bookDao) Create(tx *gorm.DB, data *BookModel) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *bookDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&BookModel{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *bookDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&BookModel{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *bookDao) FindAll() (list []BookModel, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *bookDao) FindAllWhere(query interface{}, args ...interface{}) (list []BookModel, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *bookDao) FindOneWhere(query interface{}, args ...interface{}) (record BookModel, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *bookDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&BookModel{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *bookDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&BookModel{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *bookDao) Raw(sqlStr string, params ...interface{}) (list []BookModel, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}
func (dao *bookDao) WhereQuery(query interface{}, args ...interface{}) *bookDao {
	return &bookDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *bookDao) WhereUserNameLike(username string) *bookDao {
	return &bookDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *bookDao) WhereDisabled(isDisabled int) *bookDao {
	return &bookDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

func (dao *bookDao) Paginate(offset, limit int) (count int64, list []BookModel, err error) {
	db := dao.gm.Model(&BookModel{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *bookDao) Debug() *bookDao {
	return &bookDao{
		dao.gm.Debug(),
	}
}

func (dao *bookDao) Offset(offset int) *bookDao {
	return &bookDao{
		dao.gm.Offset(offset),
	}
}

func (dao *bookDao) Limit(limit int) *bookDao {
	return &bookDao{
		dao.gm.Limit(limit),
	}
}

func (dao *bookDao) OrderBy(sortFlag, sortOrder string) *bookDao {
	return &bookDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

func (dao *bookDao) Joins(query string, args ...interface{}) *bookDao {
	return &bookDao{
		dao.gm.Joins(query, args),
	}
}

func (dao *bookDao) Preloads(query string) *bookDao {
	return &bookDao{
		dao.gm.Preload(query),
	}
}
