-- Adminer 4.8.1 MySQL 5.5.5-10.8.3-MariaDB-1:10.8.3+maria~jammy dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `blog`;
CREATE DATABASE `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `blog`;

DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '''''' COMMENT '文章标题',
  `desc` varchar(255) NOT NULL DEFAULT '''''' COMMENT '文章简述',
  `cover_image_url` varchar(255) NOT NULL DEFAULT '''''' COMMENT '封面图片地址',
  `content` longtext NOT NULL COMMENT '文章内容',
  `state` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '状态0为禁用，1为启用',
  `created_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '标签ID',
  `created_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';


DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) NOT NULL DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) NOT NULL DEFAULT '' COMMENT 'Secret',
  `created_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';


DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
  `state` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '状态0为禁用，1为启用',
  `created_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否已删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';
