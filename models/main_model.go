package models

// CREATE TABLE IF NOT EXISTS `tally_main` (
// 	`id` int(11) NOT NULL,
// 	`user_id` int(11) DEFAULT NULL,
// 	`book_id` int(11) NOT NULL COMMENT '关联账本',
// 	`category_id` int(11) NOT NULL COMMENT '关联分类',
// 	`type` int(2) NOT NULL DEFAULT '1' COMMENT '所属类型1支出2收入3不计入收支',
// 	`money` decimal(10,2) NOT NULL COMMENT '金额',
// 	`name` varchar(256) NOT NULL COMMENT '名称',
// 	`desc` varchar(256) NOT NULL COMMENT '备注',
// 	`date` date DEFAULT NULL COMMENT '记账日期',
// 	`is_del` tinyint(2) DEFAULT '1' COMMENT '1正常2删除',
// 	`create_time` int(11) DEFAULT NULL
//   ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='记账主表';

type MainModel struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserId     uint      `gorm:"column:user_id;not null" json:"user_id"`
	BookId     uint      `gorm:"column:book_id;not null" json:"book_id"`
	CategoryId uint      `gorm:"column:category_id;not null" json:"category_id"`
	MType      int       `gorm:"column:type;not null" json:"type"`
	Money      float64   `gorm:"column:money;not null" json:"money"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Desc       string    `gorm:"column:desc;not null" json:"desc"`
	Date       ModelTime `gorm:"column:date" json:"date"`
	IsDel      int       `gorm:"column:is_del;not null" json:"is_del"`
	CreatedAt  ModelTime `gorm:"column:created_at" json:"created_at"`
}

func (um MainModel) TableName() string {
	return "tally_main"
}

type MainStaModel struct {
	NickName  string    `gorm:"column:nick_name;not null" json:"nick_name"`
	AvatarUrl string    `gorm:"column:avatar_url;not null" json:"avatar_url"`
	Pay       float64   `gorm:"column:pay;not null" json:"pay"`
	Time      ModelTime `gorm:"column:time" json:"time"`
}

func (um MainStaModel) TableName() string {
	return "tally_main"
}
