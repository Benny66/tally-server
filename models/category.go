package models

import (
	database "github.com/Benny66/tally-server/db"

	"gorm.io/gorm"
)

var CategoryDao *categoryDao

func init() {
	CategoryDao = NewCategoryDao()
}

func NewCategoryDao() *categoryDao {
	return &categoryDao{
		gm: database.Orm.DB(),
	}
}

type categoryDao struct {
	gm *gorm.DB
}

func (dao *categoryDao) Create(tx *gorm.DB, data *CategoryModel) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *categoryDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&CategoryModel{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *categoryDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&CategoryModel{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *categoryDao) FindAll() (list []CategoryModel, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *categoryDao) FindAllWhere(query interface{}, args ...interface{}) (list []CategoryModel, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *categoryDao) FindOneWhere(query interface{}, args ...interface{}) (record CategoryModel, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *categoryDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&CategoryModel{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *categoryDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&CategoryModel{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *categoryDao) Raw(sqlStr string, params ...interface{}) (list []CategoryModel, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}
func (dao *categoryDao) WhereQuery(query interface{}, args ...interface{}) *categoryDao {
	return &categoryDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *categoryDao) WhereUserNameLike(username string) *categoryDao {
	return &categoryDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *categoryDao) WhereDisabled(isDisabled int) *categoryDao {
	return &categoryDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

func (dao *categoryDao) Paginate(offset, limit int) (count int64, list []CategoryModel, err error) {
	db := dao.gm.Model(&CategoryModel{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *categoryDao) Debug() *categoryDao {
	return &categoryDao{
		dao.gm.Debug(),
	}
}

func (dao *categoryDao) Offset(offset int) *categoryDao {
	return &categoryDao{
		dao.gm.Offset(offset),
	}
}

func (dao *categoryDao) Limit(limit int) *categoryDao {
	return &categoryDao{
		dao.gm.Limit(limit),
	}
}

func (dao *categoryDao) OrderBy(sortFlag, sortOrder string) *categoryDao {
	return &categoryDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

func (dao *categoryDao) Joins(query string, args ...interface{}) *categoryDao {
	return &categoryDao{
		dao.gm.Joins(query, args),
	}
}

func (dao *categoryDao) Preloads(query string) *categoryDao {
	return &categoryDao{
		dao.gm.Preload(query),
	}
}
