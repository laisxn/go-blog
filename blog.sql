/*
Navicat MySQL Data Transfer

Source Server         : 博客
Source Server Version : 50735
Source Host           : 101.34.99.204:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50735
File Encoding         : 65001

Date: 2022-08-29 21:55:39
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `short_content` varchar(255) NOT NULL DEFAULT '' COMMENT '短内容',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
  `click_num` int(11) NOT NULL DEFAULT '0' COMMENT '浏览量',
  `content` text COMMENT '内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `title` (`title`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='博客-文章';

-- ----------------------------
-- Table structure for chat_records
-- ----------------------------
DROP TABLE IF EXISTS `chat_records`;
CREATE TABLE `chat_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'pid',
  `ip` varchar(32) NOT NULL DEFAULT '' COMMENT 'ip',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `user_nickname` varchar(100) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `content` text COMMENT '内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='记录';

-- ----------------------------
-- Table structure for click_records
-- ----------------------------
DROP TABLE IF EXISTS `click_records`;
CREATE TABLE `click_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ip` varchar(32) NOT NULL DEFAULT '' COMMENT 'ip',
  `article_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='博客-文章-浏览记录';

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'pid',
  `ip` varchar(32) NOT NULL DEFAULT '' COMMENT 'ip',
  `article_id` int(11) NOT NULL DEFAULT '0' COMMENT '文章id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `user_nickname` varchar(100) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `content` text COMMENT '内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `article_id` (`article_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论';