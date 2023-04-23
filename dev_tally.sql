-- phpMyAdmin SQL Dump
-- version 4.4.15.10
-- https://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2022-10-20 18:34:13
-- 服务器版本： 5.6.50-log
-- PHP Version: 5.6.40

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dev_tally`
--

-- --------------------------------------------------------

--
-- 表的结构 `tally_book`
--

CREATE TABLE IF NOT EXISTS `tally_book` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL COMMENT '关联用户',
  `cover_url` varchar(256) DEFAULT NULL COMMENT '账本封面',
  `name` varchar(100) DEFAULT NULL COMMENT '账本名称',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `is_del` tinyint(2) DEFAULT '1' COMMENT '1正常2删除',
  `create_time` int(11) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='记账账本表';

--
-- 转存表中的数据 `tally_book`
--

INSERT INTO `tally_book` (`id`, `user_id`, `cover_url`, `name`, `sort`, `is_del`, `create_time`) VALUES
(1, 2, NULL, '日常记账', 0, 1, 1666172211),
(2, 9, NULL, '哈哈', 1, 1, 1666175815),
(3, 2, NULL, '大额账本', 1, 1, 1666194978),
(4, 2, NULL, '朋友借钱', 0, 1, 1666194999),
(5, 21, NULL, '测试', 0, 1, 1666261024);

-- --------------------------------------------------------

--
-- 表的结构 `tally_category`
--

CREATE TABLE IF NOT EXISTS `tally_category` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL COMMENT '分类名称',
  `type` int(2) DEFAULT '1' COMMENT '所属类型1支出2收入3不计入收支',
  `sort` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT '0' COMMENT '关联用户空代表系统固定',
  `is_del` tinyint(2) DEFAULT '1' COMMENT '1正常2删除',
  `create_time` int(11) DEFAULT NULL,
  `icon_url` varchar(256) DEFAULT NULL COMMENT '图标地址'
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='记账分类';

--
-- 转存表中的数据 `tally_category`
--

INSERT INTO `tally_category` (`id`, `name`, `type`, `sort`, `user_id`, `is_del`, `create_time`, `icon_url`) VALUES
(1, '生活消费', 1, NULL, 2, 1, 1666172262, 'https://tally.noahzhou.com/static/imgs/mot1.png'),
(2, '工资', 2, NULL, 2, 1, 1666172268, 'https://tally.noahzhou.com/static/imgs/mot1.png'),
(3, '理财', 3, NULL, 2, 1, 1666172276, 'https://tally.noahzhou.com/static/imgs/mot1.png');

-- --------------------------------------------------------

--
-- 表的结构 `tally_main`
--

CREATE TABLE IF NOT EXISTS `tally_main` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `book_id` int(11) NOT NULL COMMENT '关联账本',
  `category_id` int(11) NOT NULL COMMENT '关联分类',
  `type` int(2) NOT NULL DEFAULT '1' COMMENT '所属类型1支出2收入3不计入收支',
  `money` decimal(10,2) NOT NULL COMMENT '金额',
  `name` varchar(256) NOT NULL COMMENT '名称',
  `desc` varchar(256) NOT NULL COMMENT '备注',
  `date` date DEFAULT NULL COMMENT '记账日期',
  `is_del` tinyint(2) DEFAULT '1' COMMENT '1正常2删除',
  `create_time` int(11) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='记账主表';

--
-- 转存表中的数据 `tally_main`
--

INSERT INTO `tally_main` (`id`, `user_id`, `book_id`, `category_id`, `type`, `money`, `name`, `desc`, `date`, `is_del`, `create_time`) VALUES
(1, 2, 1, 1, 1, '4.00', '喝水饮料', '', '2022-10-19', 1, 1666172430),
(2, 2, 1, 1, 1, '23.00', '中午吃饭', '', '2022-10-19', 1, 1666172561),
(3, 2, 1, 1, 1, '2.00', '共享单车', '', '2022-10-20', 1, 1666225502),
(4, 2, 1, 1, 1, '1.00', '骑车到公司', '', '2022-10-20', 1, 1666242804),
(5, 2, 1, 1, 1, '28.00', '中午吃饭', '', '2022-10-20', 1, 1666242824);

-- --------------------------------------------------------

--
-- 表的结构 `tally_phrase`
--

CREATE TABLE IF NOT EXISTS `tally_phrase` (
  `id` int(11) NOT NULL,
  `phrase` varchar(128) NOT NULL COMMENT '祝福语',
  `create_time` int(11) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='祝福短语表';

--
-- 转存表中的数据 `tally_phrase`
--

INSERT INTO `tally_phrase` (`id`, `phrase`, `create_time`) VALUES
(1, '金玉其外，败絮其中。', NULL),
(2, '乐而不淫，衰而不伤。', NULL),
(3, '学海无涯，心存高远。', NULL),
(4, '塞翁失马，焉知非福。', NULL),
(5, '从善如登，从恶如崩。', NULL),
(6, '当局者迷，旁观者清。', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `tally_user`
--

CREATE TABLE IF NOT EXISTS `tally_user` (
  `id` int(11) NOT NULL,
  `openid` varchar(256) NOT NULL COMMENT '唯一标识',
  `token` varchar(255) NOT NULL COMMENT 'token',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
  `avatar_url` varchar(256) DEFAULT NULL COMMENT '头像',
  `sex` varchar(64) DEFAULT NULL COMMENT '1男2女',
  `job` varchar(100) DEFAULT NULL COMMENT '职业',
  `create_time` int(11) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 COMMENT='会员表';

--
-- 转存表中的数据 `tally_user`
--

INSERT INTO `tally_user` (`id`, `openid`, `token`, `nick_name`, `avatar_url`, `sex`, `job`, `create_time`) VALUES
(1, 'ouZtk5e_EI8cCF_D1ErdpRkWvk0vw', 'riLW16e6UuQ9k56w3jF5c11UxkR6ihYnI7Hums2RhW5', '飘逸的荔枝', '', '1', '加班狗', 1665630787),
(2, 'ouZtk5e_EI8cCF_DErdpRkWvk0vw', 'aO11V8o66Mj6n21615E0CEcM5aC63ovkqzXEKiP5y3', '复杂的小蘑菇', 'https://tally.noahzhou.com/uploads/20221019/a747d1cf83b1168150eec974d1d11773.jpg', '男', '打工人', 1665650729),
(3, 'ouZtk5TuWGOR45f5wats9RafuNp0', 'KG1n3VsSjcySM6MAz6ff9v61k871Sj6uZXvgLL1lM1', '哭泣的鼠标', 'https://tally.noahzhou.com/uploads/txs/1666171611.png', '未知', '未知', 1666171613),
(4, 'ouZtk5RdOP8qsbd3BZK9Cfx9Ykf4', 'Y12vZtFkYXqBUMS6mK6k6z01q7ginpQ2V5VuQ094F8', '娇气的保温杯', 'https://tally.noahzhou.com/uploads/txs/1666171612.png', '未知', '未知', 1666171613),
(5, 'ouZtk5Y4vxX4ji_UUXC3h8vihbeg', 'blH1wb6lp2vb6l66187Dewv2j2bKmI6pIH1N68GNF6', '贪玩的爆米花', 'https://tally.noahzhou.com/uploads/txs/1666172186.png', '未知', '未知', 1666172188),
(6, 'ouZtk5aCwdCJw91_C_vm7esAM0so', 'Wdown10Vb661M5Wsr6S1TXCCxUEip7k29YO1ooCu86', '飘逸的钻石', 'https://tally.noahzhou.com/uploads/txs/1666172186.png', '未知', '未知', 1666172188),
(7, 'ouZtk5ajWA58NeDj-JScPO7p1erU', '11sEb6g6F6QRhbJzgCx1k57Bv2U13wbgYR5SO1d058', '勤奋的铃铛', 'https://tally.noahzhou.com/uploads/txs/1666172508.png', '未知', '未知', 1666172510),
(8, 'ouZtk5Re9GYsLRbQlCygSQVCdiZw', 'sQTC08Ukq1d6Cc66W2z1c7StdKawr4cSR54F757ls9', '平常的音响', 'https://tally.noahzhou.com/uploads/txs/1666174479.png', '未知', '未知', 1666174481),
(9, 'ouZtk5QbxD2hr3i9b75RsVdG6Xjg', 't166vMpia3AVd695E65S1C7dRZGs5K3at7CX9gC8a8', '端庄的金鱼', 'https://tally.noahzhou.com/uploads/txs/1666175788.png', '未知', '未知', 1666175790),
(10, 'ouZtk5VNx6rYomFnnm54I6GuPwo0', '1rL6oN6sG8QTMwjrZ6K8692gE15Oi1089Oben1BlT6', '香蕉花卷', 'https://tally.noahzhou.com/static/imgs/txs/5.png', '未知', '未知', 1666250916),
(11, 'ouZtk5SIZkjz6ADqk8Efj81QEDWc', 'afvQoITIM1F56xrvyyS6sz9u8lyU6o2Q5Kp0b9r1O6', '虚幻的网络', 'https://tally.noahzhou.com/static/imgs/txs/5.png', '未知', '未知', 1666250916),
(12, 'ouZtk5TzEA0_QgkIHWEjexhHeUlo', 'A1n0CL6HK1mUVvh665263CYS6CsNTy0P8bIW4jgbU6', '魁梧的斑马', 'https://tally.noahzhou.com/static/imgs/txs/3.png', '未知', '未知', 1666260846),
(13, 'ouZtk5aCopAA6h5VxzOwOIszA4jY', '8i164bmcFZ66yFRW2RTja6PnN0B2qBRO58V3P0W5T3', '搞怪的芹菜', 'https://tally.noahzhou.com/static/imgs/txs/5.png', '未知', '未知', 1666260853),
(14, 'ouZtk5TxsU0a5dNIO6CBYV29soX4', '1z666n2u1J5S60YtlizV8cuBhZf0p6B36Y9DwzLmx0', '矮小的书包', 'https://tally.noahzhou.com/static/imgs/txs/1.png', '未知', '未知', 1666260860),
(15, 'ouZtk5WXxDdSJ0iFMvp1uOYxlPSU', '1R6tVYg3glVLi6mVqWvHy0yLK62KhB6Z0BtRYQW868', '香蕉凉面', 'https://tally.noahzhou.com/static/imgs/txs/4.png', '未知', '未知', 1666260867),
(16, 'ouZtk5YvDOtU4yakJmRcmKitfKd0', '16626UZ2a7xXrz2ot6AeYJnkL08198hYHtJ2i3pIV2', '深情的航空', 'https://tally.noahzhou.com/static/imgs/txs/4.png', '未知', '未知', 1666260882),
(17, 'ouZtk5ThLNfD-zRxKiwuZhQhuxxY', 'rVs1TgH6606bv2JtKz6eUSzxEQuS6Wq195rki8vuh1', '冷静的石头', 'https://tally.noahzhou.com/static/imgs/txs/6.png', '未知', '未知', 1666260943),
(18, 'ouZtk5TGiyWycpGimr3p_aa0vE1M', 'UYT10Jfi0snPeN65Nbl6gRE62kzE4Ezq360a9W45O8', '老实的蜜粉', 'https://tally.noahzhou.com/static/imgs/txs/5.png', '未知', '未知', 1666260947),
(19, 'ouZtk5dhJt6z3781bLoG0n7UELFU', '1Geh2Vwi6Cez6Xdk662eszRkVmKmS6lUN0LRd9Ik56', '简单的花卷', 'https://tally.noahzhou.com/static/imgs/txs/1.png', '未知', '未知', 1666260956),
(20, 'ouZtk5dhJt6z3781bLoG0n7UELFU', '1Geh2Vwi6Cez6Xdk662eszRkVmKmS6lUN0LRd9Ik56', '大力的奇异果', 'https://tally.noahzhou.com/static/imgs/txs/1.png', '未知', '未知', 1666260956),
(21, 'ouZtk5Zdmse67tcGBRiLQb0ihAaw', 'ZF7ERIkOb4416NB46X6C2UWrv6AQP09GC2a9NDVOz3', '复杂的季节', 'https://tally.noahzhou.com/static/imgs/txs/4.png', '未知', '未知', 1666260992),
(22, 'ouZtk5a1jQwUpTU0ST0kZGgHWH04', 'OLjNbeW16Z6f64V326MXo2nhtC12tTxWOtMP142vk9', '甜蜜的小蝴蝶', 'https://tally.noahzhou.com/static/imgs/txs/2.png', '未知', '未知', 1666260994),
(23, 'ouZtk5WH7OBHbOGEz7uxxh9CnYtY', '8jqAai8E1t6TW6odyp6mxZp26uDQtfJ1c11L4gjlz7', '美满的指甲油', 'https://tally.noahzhou.com/static/imgs/txs/6.png', '未知', '未知', 1666261104),
(24, 'ouZtk5WH7OBHbOGEz7uxxh9CnYtY', '8jqAai8E1t6TW6odyp6mxZp26uDQtfJ1c11L4gjlz7', '传统的月亮', 'https://tally.noahzhou.com/static/imgs/txs/1.png', '未知', '未知', 1666261104),
(25, 'ouZtk5UP2W7WA_p7JJ2cOWsGz6jQ', '1s0r662rx6Da5Xz3yQ2G61xwN19iKML1dKI2xg5NT9', '想人陪的绿茶', 'https://tally.noahzhou.com/static/imgs/txs/6.png', '未知', '未知', 1666261129),
(26, 'ouZtk5RimZPfKvlVGP8a8tiZBCEA', 'W1CX6nya6N6wkp2k20t6UNgVVObo128t1oQBqau945', '粗心的百褶裙', 'https://tally.noahzhou.com/static/imgs/txs/1.png', '未知', '未知', 1666261195),
(27, 'ouZtk5RimZPfKvlVGP8a8tiZBCEA', 'W1CX6nya6N6wkp2k20t6UNgVVObo128t1oQBqau945', '故意的羊', 'https://tally.noahzhou.com/static/imgs/txs/6.png', '未知', '未知', 1666261195),
(28, 'ouZtk5Q04xTbHHnQaR-etXTb9c6Y', 'y5cr16GqVc65WU6ZxME2uIEq6Ui1244y83XU9rmp18', '沉静的微笑', 'https://tally.noahzhou.com/static/imgs/txs/1.png', '未知', '未知', 1666261234),
(29, 'ouZtk5bTLQP0HKz_OcNqGliY39HQ', '1uVMrly6GA6FRh68Gz662gA4tZKXb2j613KN734oP6', '高挑的蜗牛', 'https://tally.noahzhou.com/static/imgs/txs/3.png', '未知', '未知', 1666261336),
(30, 'ouZtk5UKp0y764u1--hYbm6gD9bE', 'aM1RP6t6RDhC3e6i2Ru8ab6MT2114tMrIU23MNdEO0', '感性的抽屉', 'https://tally.noahzhou.com/static/imgs/txs/6.png', '未知', '未知', 1666261358);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tally_book`
--
ALTER TABLE `tally_book`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tally_category`
--
ALTER TABLE `tally_category`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tally_main`
--
ALTER TABLE `tally_main`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tally_phrase`
--
ALTER TABLE `tally_phrase`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tally_user`
--
ALTER TABLE `tally_user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tally_book`
--
ALTER TABLE `tally_book`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `tally_category`
--
ALTER TABLE `tally_category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=4;
--
-- AUTO_INCREMENT for table `tally_main`
--
ALTER TABLE `tally_main`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `tally_phrase`
--
ALTER TABLE `tally_phrase`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=7;
--
-- AUTO_INCREMENT for table `tally_user`
--
ALTER TABLE `tally_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=31;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
