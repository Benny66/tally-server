package models

type BookModel struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	UserId    uint      `gorm:"column:user_id;not null" json:"user_id"`
	CoverUrl  string    `gorm:"column:cover_url;not null" json:"cover_url"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Sort      int       `gorm:"column:sort;not null" json:"sort"`
	IsDel     int       `gorm:"column:is_del;not null" json:"is_del"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
}

func (um BookModel) TableName() string {
	return "tally_book"
}
