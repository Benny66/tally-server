
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tally_book
-- ----------------------------
DROP TABLE IF EXISTS `tally_book`;
CREATE TABLE `tally_book` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '关联用户',
  `cover_url` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '账本封面',
  `name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '账本名称',
  `sort` int DEFAULT NULL COMMENT '排序',
  `is_del` tinyint DEFAULT '1' COMMENT '1正常2删除',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='记账账本表';

-- ----------------------------
-- Table structure for tally_category
-- ----------------------------
DROP TABLE IF EXISTS `tally_category`;
CREATE TABLE `tally_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '分类名称',
  `type` int DEFAULT '1' COMMENT '所属类型1支出2收入3不计入收支',
  `sort` int DEFAULT NULL,
  `user_id` int DEFAULT '0' COMMENT '关联用户空代表系统固定',
  `is_del` tinyint DEFAULT '1' COMMENT '1正常2删除',
  `icon_url` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '图标地址',
  `created_at` datetime DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='记账分类';

-- ----------------------------
-- Table structure for tally_main
-- ----------------------------
DROP TABLE IF EXISTS `tally_main`;
CREATE TABLE `tally_main` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `book_id` int NOT NULL COMMENT '关联账本',
  `category_id` int NOT NULL COMMENT '关联分类',
  `type` int NOT NULL DEFAULT '1' COMMENT '所属类型1支出2收入3不计入收支',
  `money` decimal(10,2) NOT NULL COMMENT '金额',
  `name` varchar(256) COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `desc` varchar(256) COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `date` datetime DEFAULT NULL COMMENT '记账日期',
  `is_del` tinyint DEFAULT '1' COMMENT '1正常2删除',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='记账主表';

-- ----------------------------
-- Table structure for tally_phrase
-- ----------------------------
DROP TABLE IF EXISTS `tally_phrase`;
CREATE TABLE `tally_phrase` (
  `id` int NOT NULL AUTO_INCREMENT,
  `phrase` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '祝福语',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='祝福短语表';

-- ----------------------------
-- Table structure for tally_user
-- ----------------------------
DROP TABLE IF EXISTS `tally_user`;
CREATE TABLE `tally_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `openid` varchar(256) COLLATE utf8mb4_general_ci NOT NULL COMMENT '唯一标识',
  `token` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'token',
  `nick_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `avatar_url` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像',
  `sex` int DEFAULT NULL COMMENT '1男2女',
  `job` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '职业',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='会员表';

-- ----------------------------
-- Table structure for w_project
-- ----------------------------
DROP TABLE IF EXISTS `w_project`;
CREATE TABLE `w_project` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `git` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '',
  `content` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `image` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `sort` int NOT NULL,
  `views` int NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
