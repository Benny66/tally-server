package models

type UserModel struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	OpenId    string    `gorm:"column:openid;unique;not null" json:"openid"`
	Token     string    `gorm:"column:token;not null" json:"token"`
	NickName  string    `gorm:"column:nick_name;not null" json:"nick_name"`
	AvatarUrl string    `gorm:"column:avatar_url;not null" json:"avatar_url"`
	Sex       int       `gorm:"column:sex;not null" json:"sex"`
	Job       string    `gorm:"column:job;not null" json:"job"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt ModelTime `gorm:"column:updated_at" json:"updated_at"`
}

func (um UserModel) TableName() string {
	return "tally_user"
}
