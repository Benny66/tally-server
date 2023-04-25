package models

import (
	database "github.com/Benny66/tally-server/db"

	"gorm.io/gorm"
)

var ProjectDao *projectDao

func init() {
	ProjectDao = NewProjectDao()
}

func NewProjectDao() *projectDao {
	return &projectDao{
		gm: database.Orm.DB(),
	}
}

type projectDao struct {
	gm *gorm.DB
}

func (dao *projectDao) Create(tx *gorm.DB, data *ProjectModel) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *projectDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&ProjectModel{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *projectDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&ProjectModel{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *projectDao) FindAll() (list []ProjectModel, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *projectDao) FindAllWhere(query interface{}, args ...interface{}) (list []ProjectModel, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *projectDao) FindOneWhere(query interface{}, args ...interface{}) (record ProjectModel, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *projectDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&ProjectModel{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *projectDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&ProjectModel{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *projectDao) Raw(sqlStr string, params ...interface{}) (list []ProjectModel, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}
func (dao *projectDao) WhereQuery(query interface{}, args ...interface{}) *projectDao {
	return &projectDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *projectDao) WhereUserNameLike(username string) *projectDao {
	return &projectDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *projectDao) WhereDisabled(isDisabled int) *projectDao {
	return &projectDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

func (dao *projectDao) Paginate(offset, limit int) (count int64, list []ProjectModel, err error) {
	db := dao.gm.Model(&ProjectModel{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *projectDao) Debug() *projectDao {
	return &projectDao{
		dao.gm.Debug(),
	}
}

func (dao *projectDao) Offset(offset int) *projectDao {
	return &projectDao{
		dao.gm.Offset(offset),
	}
}

func (dao *projectDao) Limit(limit int) *projectDao {
	return &projectDao{
		dao.gm.Limit(limit),
	}
}

func (dao *projectDao) OrderBy(sortFlag, sortOrder string) *projectDao {
	return &projectDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

func (dao *projectDao) Joins(query string, args ...interface{}) *projectDao {
	return &projectDao{
		dao.gm.Joins(query, args),
	}
}

func (dao *projectDao) Preloads(query string) *projectDao {
	return &projectDao{
		dao.gm.Preload(query),
	}
}
