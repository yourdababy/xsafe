/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : localhost:3306
 Source Schema         : xsafe

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 28/08/2021 18:05:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` char(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '标题',
  `content` text COMMENT '内容',
  `create_at` int(10) unsigned NOT NULL COMMENT '创建时间时间戳',
  `update_at` int(10) unsigned NOT NULL COMMENT '更新时间',
  `classify_id` int(10) unsigned DEFAULT '0' COMMENT '分类ID',
  `author_id` int(10) unsigned DEFAULT '0' COMMENT '作者ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38722 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for author
-- ----------------------------
DROP TABLE IF EXISTS `author`;
CREATE TABLE `author` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '作者名称',
  `tmpId` int(11) unsigned DEFAULT NULL COMMENT '临时ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for classify
-- ----------------------------
DROP TABLE IF EXISTS `classify`;
CREATE TABLE `classify` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `tmpId` int(11) unsigned DEFAULT NULL COMMENT '临时ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
