package models

// CREATE TABLE IF NOT EXISTS `tally_category` (
// 	`id` int(11) NOT NULL,
// 	`name` varchar(100) DEFAULT NULL COMMENT '分类名称',
// 	`type` int(2) DEFAULT '1' COMMENT '所属类型1支出2收入3不计入收支',
// 	`sort` int(11) DEFAULT NULL,
// 	`user_id` int(11) DEFAULT '0' COMMENT '关联用户空代表系统固定',
// 	`is_del` tinyint(2) DEFAULT '1' COMMENT '1正常2删除',
// 	`create_time` int(11) DEFAULT NULL,
// 	`icon_url` varchar(256) DEFAULT NULL COMMENT '图标地址'
//   ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='记账分类';

type CategoryModel struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	UserId    uint      `gorm:"column:user_id;not null" json:"user_id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	CType     int       `gorm:"column:type;not null" json:"type"`
	Sort      int       `gorm:"column:sort;not null" json:"sort"`
	IconUrl   string    `gorm:"column:icon_url;not null" json:"icon_url"`
	IsDel     int       `gorm:"column:is_del;not null" json:"is_del"`
	CreatedAt ModelTime `gorm:"column:created_at" json:"created_at"`
}

func (um CategoryModel) TableName() string {
	return "tally_category"
}
