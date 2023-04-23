package models

type PhraseModel struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Phrase    string    `gorm:"column:phrase;unique;not null" json:"phrase"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
}

func (um PhraseModel) TableName() string {
	return "tally_phrase"
}
