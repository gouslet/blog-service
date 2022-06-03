-- Adminer 4.8.1 MySQL 5.5.5-10.8.3-MariaDB-1:10.8.3+maria~jammy dump
USE `blog`;

DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '''''' COMMENT '文章标题',
  `desc` varchar(255) NOT NULL DEFAULT '''''' COMMENT '文章简述',
  `cover_image_url` varchar(255) NOT NULL DEFAULT '''''' COMMENT '封面图片地址',
  `content` longtext NOT NULL COMMENT '文章内容',
  `state` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '状态0为禁用，1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';

-- 2022-06-03 04:40:46