package models

import (
	database "github.com/Benny66/tally-server/db"

	"gorm.io/gorm"
)

var PhraseDao *phraseDao

func init() {
	PhraseDao = NewPhraseDao()
}

func NewPhraseDao() *phraseDao {
	return &phraseDao{
		gm: database.Orm.DB(),
	}
}

type phraseDao struct {
	gm *gorm.DB
}

func (dao *phraseDao) Create(tx *gorm.DB, data *PhraseModel) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *phraseDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&PhraseModel{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *phraseDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&PhraseModel{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *phraseDao) FindAll() (list []PhraseModel, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *phraseDao) FindAllWhere(query interface{}, args ...interface{}) (list []PhraseModel, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *phraseDao) FindOneWhere(query interface{}, args ...interface{}) (record PhraseModel, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *phraseDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&PhraseModel{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *phraseDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&PhraseModel{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *phraseDao) Raw(sqlStr string, params ...interface{}) (list []PhraseModel, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}
func (dao *phraseDao) WhereQuery(query interface{}, args ...interface{}) *phraseDao {
	return &phraseDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *phraseDao) WhereUserNameLike(username string) *phraseDao {
	return &phraseDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *phraseDao) WhereDisabled(isDisabled int) *phraseDao {
	return &phraseDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

func (dao *phraseDao) Paginate(offset, limit int) (count int64, list []PhraseModel, err error) {
	db := dao.gm.Model(&PhraseModel{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *phraseDao) Debug() *phraseDao {
	return &phraseDao{
		dao.gm.Debug(),
	}
}

func (dao *phraseDao) Offset(offset int) *phraseDao {
	return &phraseDao{
		dao.gm.Offset(offset),
	}
}

func (dao *phraseDao) Limit(limit int) *phraseDao {
	return &phraseDao{
		dao.gm.Limit(limit),
	}
}

func (dao *phraseDao) OrderBy(sortFlag, sortOrder string) *phraseDao {
	return &phraseDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

func (dao *phraseDao) Joins(query string, args ...interface{}) *phraseDao {
	return &phraseDao{
		dao.gm.Joins(query, args),
	}
}

func (dao *phraseDao) Preloads(query string) *phraseDao {
	return &phraseDao{
		dao.gm.Preload(query),
	}
}
