-- Adminer 4.8.1 MySQL 5.5.5-10.8.3-MariaDB-1:10.8.3+maria~jammy dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `blog`;
CREATE DATABASE `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `blog`;

DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已删除 0 为未删除、1 为已删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 2022-05-28 16:11:03