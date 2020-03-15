
CREATE DATABASE IF NOT EXISTS `ptt`;
USE `ptt`;

CREATE TABLE IF NOT EXISTS `articles` (
  `url` varchar(100) NOT NULL,
  `board` varchar(50) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `author` varchar(50) DEFAULT NULL,
  `contains` longtext DEFAULT NULL,
  `reply` longtext DEFAULT NULL,
  `date` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`url`),
  UNIQUE KEY `url` (`url`),
  KEY `board` (`board`),
  KEY `author` (`author`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `articles` (`url`, `board`, `title`, `author`, `contains`, `reply`, `date`, `updated_at`) VALUES
	('https://www.ptt.cc/bbs/Gossiping/M.1566347622.A.9C7.html', 'Gossiping', '[公告] 八卦板板規(2019.08.21)', 'arsonlolita (蘿莉塔)', 'respect ', 'respect ', '2019-08-21 08:33:39', '2020-03-15 21:49:10'),
	('https://www.ptt.cc/bbs/Soft_Job/M.1501827536.A.DF2.html', 'Soft_Job', '[公告] 2016年1月1日 新板規上路! 新年快樂', 'MOONY135 (談無慾)', 'respect ', 'respect ', '2017-08-04 14:18:53', '2020-03-15 21:49:10'),
	('https://www.ptt.cc/bbs/TaichungBun/M.1576314766.A.41A.html', 'TaichungBun', '[公告] 板規/ Q&A / 檢舉代碼+ID+板規+日期', 'nicely (你在哪裡?)', 'respect ', 'respect ', '2019-12-14 17:12:40', '2020-03-15 21:49:10'),
	('https://www.ptt.cc/bbs/Tech_Job/M.1410062073.A.1B6.html', 'Tech_Job', '[公告] 置底 檢舉/推薦 文章', 'mmkntust (老王廚房)', 'respect ', 'respect ', '2014-09-07 11:54:29', '2020-03-15 21:49:10');
