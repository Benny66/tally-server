package models

type ProjectModel struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Git       string    `gorm:"column:git;not null" json:"git"`
	Image     string    `gorm:"column:image;not null" json:"image"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Content   string    `gorm:"column:content;not null" json:"content"`
	Sort      int       `gorm:"column:sort;not null" json:"sort"`
	Views     int       `gorm:"column:views;not null" json:"views"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
}

func (um ProjectModel) TableName() string {
	return "w_project"
}
