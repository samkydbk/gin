/*
Navicat MySQL Data Transfer

Source Server         : local_mysql
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : gin

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2021-02-05 11:15:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for z_category
-- ----------------------------
DROP TABLE IF EXISTS `z_category`;
CREATE TABLE `z_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `alias` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类别名',
  `img` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标地址',
  `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of z_category
-- ----------------------------

-- ----------------------------
-- Table structure for z_comment
-- ----------------------------
DROP TABLE IF EXISTS `z_comment`;
CREATE TABLE `z_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论者昵称',
  `website` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称链接',
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '评论内容',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of z_comment
-- ----------------------------

-- ----------------------------
-- Table structure for z_content
-- ----------------------------
DROP TABLE IF EXISTS `z_content`;
CREATE TABLE `z_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL DEFAULT '0',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of z_content
-- ----------------------------

-- ----------------------------
-- Table structure for z_post
-- ----------------------------
DROP TABLE IF EXISTS `z_post`;
CREATE TABLE `z_post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `alias` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文章别名',
  `category` smallint(4) NOT NULL DEFAULT '0' COMMENT '分类',
  `labels` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '外链Url',
  `view_count` int(11) NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `is_local` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否本地文档，否则是外链',
  `is_draft` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否草稿',
  `is_active` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有效',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '修改时间',
  `publish_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of z_post
-- ----------------------------

-- ----------------------------
-- Table structure for z_user
-- ----------------------------
DROP TABLE IF EXISTS `z_user`;
CREATE TABLE `z_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0-男 1-女',
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of z_user
-- ----------------------------
INSERT INTO `z_user` VALUES ('1', 'sam1', '0', '13228016321', '', '2021-01-17 09:21:39', '2021-01-17 09:21:43');
INSERT INTO `z_user` VALUES ('2', 'abc', '1', '13228016321', '', '2021-01-27 22:06:34', '2021-01-27 22:06:34');
INSERT INTO `z_user` VALUES ('3', 'abc', '1', '13228016321', '', '2021-01-27 22:06:45', '2021-01-27 22:06:45');
INSERT INTO `z_user` VALUES ('4', '123', '0', '', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `z_user` VALUES ('5', 'abc', '1', '13288888888', '', '2021-02-05 10:01:30', '2021-02-05 11:06:53');
