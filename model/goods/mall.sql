/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : mall

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-11-06 15:28:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for attribute
-- ----------------------------
DROP TABLE IF EXISTS `attribute`;
CREATE TABLE `attribute` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(100) DEFAULT '' COMMENT '属性名称',
  `spec_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '规格id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='属性表';

-- ----------------------------
-- Records of attribute
-- ----------------------------
INSERT INTO `attribute` VALUES ('1', '亮黑色', '1', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('2', '釉白色', '1', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('3', '秘银色', '1', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('4', '夏日胡杨', '1', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('5', '秋日胡杨', '1', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('6', '5G全网通 8GB+128GB', '2', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('7', '5G全网通 8GB+256GB', '2', '2020-11-05 15:46:01', null);
INSERT INTO `attribute` VALUES ('8', '5G全网通 8GB+512GB', '2', '2020-11-05 15:46:01', null);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键⁯id',
  `name` varchar(100) DEFAULT '' COMMENT '分类名称',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='分类表';

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES ('1', '手机', '2020-11-05 15:43:45', null);

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(100) DEFAULT '' COMMENT '商品名称',
  `category_id` int(10) unsigned DEFAULT '0' COMMENT '分类id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='商品表';

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES ('1', 'HUAWEI Mate40 Pro', '1', '2020-11-05 15:43:58', null);

-- ----------------------------
-- Table structure for goods_attribute
-- ----------------------------
DROP TABLE IF EXISTS `goods_attribute`;
CREATE TABLE `goods_attribute` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `goods_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `attribute_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '属性id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='商品属性表';

-- ----------------------------
-- Records of goods_attribute
-- ----------------------------
INSERT INTO `goods_attribute` VALUES ('1', '1', '1', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('2', '1', '2', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('3', '1', '3', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('4', '1', '4', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('5', '1', '5', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('6', '1', '6', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('7', '1', '7', '2020-11-05 15:53:25', null);
INSERT INTO `goods_attribute` VALUES ('8', '1', '8', '2020-11-05 15:53:25', null);

-- ----------------------------
-- Table structure for goods_spec
-- ----------------------------
DROP TABLE IF EXISTS `goods_spec`;
CREATE TABLE `goods_spec` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `goods_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `spec_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '规格id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='商品规格表';

-- ----------------------------
-- Records of goods_spec
-- ----------------------------
INSERT INTO `goods_spec` VALUES ('1', '1', '1', '2020-11-05 15:53:11', null);
INSERT INTO `goods_spec` VALUES ('2', '1', '2', '2020-11-05 15:53:11', null);

-- ----------------------------
-- Table structure for sku
-- ----------------------------
DROP TABLE IF EXISTS `sku`;
CREATE TABLE `sku` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `goods_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT 'sku名称',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT 'sku价格',
  `stock` int(255) unsigned NOT NULL DEFAULT '0' COMMENT 'sku库存',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='sku表';

-- ----------------------------
-- Records of sku
-- ----------------------------
INSERT INTO `sku` VALUES ('1', '1', '亮黑色·5G全网通 8GB+128GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('2', '1', '亮黑色·5G全网通 8GB+256GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('3', '1', '亮黑色·5G全网通 8GB+512GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('4', '1', '釉白色·5G全网通 8GB+128GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('5', '1', '釉白色·5G全网通 8GB+256GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('6', '1', '釉白色·5G全网通 8GB+512GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('7', '1', '秘银色·5G全网通 8GB+128GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('8', '1', '秘银色·5G全网通 8GB+256GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('9', '1', '秘银色·5G全网通 8GB+512GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('10', '1', '夏日胡杨·5G全网通 8GB+128GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('11', '1', '夏日胡杨·5G全网通 8GB+256GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('12', '1', '夏日胡杨·5G全网通 8GB+512GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('13', '1', '秋日胡杨·5G全网通 8GB+128GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('14', '1', '秋日胡杨·5G全网通 8GB+256GB', '100.00', '10', '2020-11-05 16:02:28', null);
INSERT INTO `sku` VALUES ('15', '1', '秋日胡杨·5G全网通 8GB+512GB', '100.00', '10', '2020-11-05 16:02:28', null);

-- ----------------------------
-- Table structure for sku_attribute
-- ----------------------------
DROP TABLE IF EXISTS `sku_attribute`;
CREATE TABLE `sku_attribute` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `sku_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'sku id',
  `attribute_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '属性id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 COMMENT='sku属性表';

-- ----------------------------
-- Records of sku_attribute
-- ----------------------------
INSERT INTO `sku_attribute` VALUES ('1', '1', '1', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('2', '1', '6', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('3', '2', '1', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('4', '2', '7', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('5', '3', '1', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('6', '3', '8', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('7', '4', '2', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('8', '4', '6', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('9', '5', '2', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('10', '5', '7', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('11', '6', '2', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('12', '6', '8', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('13', '7', '3', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('14', '7', '6', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('15', '8', '3', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('16', '8', '7', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('17', '9', '3', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('18', '9', '8', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('19', '10', '4', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('20', '10', '6', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('21', '11', '4', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('22', '11', '7', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('23', '12', '4', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('24', '12', '8', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('25', '13', '5', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('26', '13', '6', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('27', '14', '5', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('28', '14', '7', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('29', '15', '5', '2020-11-05 16:13:05', null);
INSERT INTO `sku_attribute` VALUES ('30', '15', '8', '2020-11-05 16:13:05', null);

-- ----------------------------
-- Table structure for spec
-- ----------------------------
DROP TABLE IF EXISTS `spec`;
CREATE TABLE `spec` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(100) DEFAULT '' COMMENT '规格名称',
  `category_id` int(10) unsigned DEFAULT '0' COMMENT '分类id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='规格表';

-- ----------------------------
-- Records of spec
-- ----------------------------
INSERT INTO `spec` VALUES ('1', '颜色', '1', '2020-11-05 15:46:01', null);
INSERT INTO `spec` VALUES ('2', '版本', '1', '2020-11-05 15:46:01', null);
