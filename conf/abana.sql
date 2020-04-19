-- MySQL dump 10.13  Distrib 8.0.15, for macos10.14 (x86_64)
--
-- Host: 47.98.135.130    Database: abana
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `activity`
--

DROP TABLE IF EXISTS `activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `activity` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(60) NOT NULL DEFAULT '' COMMENT '标题',
  `content` varchar(250) NOT NULL DEFAULT '' COMMENT '内容',
  `like_num` int(6) NOT NULL DEFAULT '0' COMMENT '喜欢 收藏的量',
  `share_num` int(6) NOT NULL DEFAULT '0' COMMENT '转发的量',
  `evaluate_num` int(6) NOT NULL DEFAULT '0' COMMENT '评论的量',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `c_t` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `u_t` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `type` int(5) NOT NULL COMMENT '类别',
  `status` int(2) NOT NULL DEFAULT '1' COMMENT '状态1 0',
  `avatar` varchar(140) NOT NULL DEFAULT '' COMMENT '用户头像',
  `nick_name` varchar(45) NOT NULL DEFAULT '' COMMENT '用户昵称',
  PRIMARY KEY (`id`,`title`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity`
--

LOCK TABLES `activity` WRITE;
/*!40000 ALTER TABLE `activity` DISABLE KEYS */;
INSERT INTO `activity` VALUES (18,'2222222','33333333333',0,0,0,4315836759645566,1587280074,1587280074,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥'),(19,'ddddddddd','ffffffffffff',0,0,0,4315836759645566,1587280750,1587280750,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥'),(20,'vvvvvv','eeeeeeeeee',0,0,0,4315836759645566,1587280888,1587280888,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥'),(21,'rrrrr','ffffff',0,0,0,4315836759645566,1587280999,1587280999,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥');
/*!40000 ALTER TABLE `activity` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(60) NOT NULL DEFAULT '' COMMENT '标题',
  `content` varchar(250) NOT NULL DEFAULT '' COMMENT '内容',
  `like_num` int(6) NOT NULL DEFAULT '0' COMMENT '喜欢 收藏的量',
  `share_num` int(6) NOT NULL DEFAULT '0' COMMENT '转发的量',
  `evaluate_num` int(6) NOT NULL DEFAULT '0' COMMENT '评论的量',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `c_t` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `u_t` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `type` int(5) NOT NULL COMMENT '类别',
  `status` int(2) NOT NULL DEFAULT '1' COMMENT '状态1 0',
  `avatar` varchar(140) NOT NULL DEFAULT '' COMMENT '用户头像',
  `nick_name` varchar(45) NOT NULL DEFAULT '' COMMENT '用户昵称',
  PRIMARY KEY (`id`,`title`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (1,'hhhhh哈哈哈哈','少时诵诗书所所所所',1,0,0,61583579141,1583666113,1583666113,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(2,'女人是','啊啊啊啊啊啊啊啊啊啊啊啊啊',1,0,0,61583579141,1583666130,1583666130,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(3,'你是真的傻逼啊你','反反复复付',0,0,0,61583579141,1583666209,1583666209,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(4,'反反复复付付付付付付付','哈哈哈哈哈哈哈哈哈',0,0,0,61583579141,1583666218,1583666218,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(5,'222222222222222','咔咔咔咔咔咔扩扩扩扩',1,0,0,61583579141261982,1583666228,1583666228,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(6,'222222','反反复复付',0,0,0,61583579141261982,1583666248,1583666248,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(7,'888888','22222222',0,0,0,61583579141261982,1583666255,1583666255,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(8,'11111','66666666',0,0,0,61583579141261982,1583666263,1583666263,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(9,'999999999','1111111111',0,0,0,61583579141261982,1583666271,1583666271,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(10,'9999999999','1111',0,0,0,61583579141261982,1583666279,1583666279,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(11,'3333333','4444444',0,0,0,61583579141261982,1583666287,1583666287,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(12,'666666','11111111',1,0,0,61583579141261982,1583666295,1583666295,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(13,'rrrrr','eeeeee',1,0,0,61583579141261982,1583669444,1583669444,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132',''),(14,'反反复复付付付付','反反复复付付付付付付付付付付付付付付付付付付付付付付付付付付付付付付',0,0,0,4315836759645566,1583758992,1583758992,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥'),(15,'顶顶顶顶','的点点滴滴',0,0,2,4315836759645566,1583760044,1583760044,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥'),(16,'qqqqq','11111111',1,0,0,4315836759645566,1583852358,1583852358,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥'),(17,'fff','33333',1,0,0,4315836759645566,1583852382,1583852382,0,1,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','乌蝇哥');
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `evaluate`
--

DROP TABLE IF EXISTS `evaluate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `evaluate` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `detail` varchar(200) NOT NULL DEFAULT '' COMMENT '详情内容',
  `c_t` int(11) NOT NULL DEFAULT '0' COMMENT '评论时间',
  `status` int(1) NOT NULL DEFAULT '1' COMMENT '默认是1展示',
  PRIMARY KEY (`id`,`article_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `evaluate`
--

LOCK TABLES `evaluate` WRITE;
/*!40000 ALTER TABLE `evaluate` DISABLE KEYS */;
INSERT INTO `evaluate` VALUES (1,15,4315836759645566,'44',1583763145,1),(2,15,4315836759645566,'kkkkkkkkkk',1583763164,1),(3,15,4315836759645566,'333333',1583763856,1),(4,15,4315836759645566,'ddddddd',1583763918,1),(5,14,4315836759645566,'dddddddddd',1583763943,1),(6,13,4315836759645566,'的点点滴滴',1583764045,1),(7,13,4315836759645566,'顶顶顶顶',1583764060,1),(8,14,4315836759645566,'3',1583765266,1),(9,15,4315836759645566,'jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj',1583766219,1),(10,15,4315836759645566,'rrrrrrr',1583848498,1),(11,15,4315836759645566,'3',1583849040,1),(12,15,4315836759645566,'ggggggggg',1583849058,1);
/*!40000 ALTER TABLE `evaluate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `my_like`
--

DROP TABLE IF EXISTS `my_like`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `my_like` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `article_id` int(11) NOT NULL,
  `c_t` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`user_id`,`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `my_like`
--

LOCK TABLES `my_like` WRITE;
/*!40000 ALTER TABLE `my_like` DISABLE KEYS */;
INSERT INTO `my_like` VALUES (1,4315836759645566,15,1583851351),(2,4315836759645566,14,1583851551),(3,4315836759645566,13,1583851860),(4,4315836759645566,12,1583851873),(5,4315836759645566,5,1583852175),(6,4315836759645566,2,1583852228),(7,4315836759645566,16,1583852674),(8,4315836759645566,17,1583852895),(9,4315836759645566,1,1583852904);
/*!40000 ALTER TABLE `my_like` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_mall_info`
--

DROP TABLE IF EXISTS `user_mall_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `user_mall_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户唯一id',
  `nick_name` varchar(45) DEFAULT NULL,
  `gender` int(2) DEFAULT '0',
  `phone` varchar(11) NOT NULL DEFAULT '0',
  `register_time` varchar(45) NOT NULL COMMENT '注册时间',
  `integral` int(10) DEFAULT '0' COMMENT '用户积分',
  `open_id` varchar(80) NOT NULL DEFAULT '0' COMMENT '用户微信唯一id',
  `if_delete` int(2) NOT NULL DEFAULT '0' COMMENT '是否被删除1是删除，0是正常用户',
  `vip_level` int(3) DEFAULT '0' COMMENT '用户会员级别，达到某种程度划分级别',
  `avatar_url` varchar(200) DEFAULT NULL COMMENT '头像',
  `city` varchar(15) DEFAULT NULL,
  `province` varchar(15) DEFAULT NULL,
  `true_name` varchar(45) NOT NULL DEFAULT '',
  `like_article_num` int(5) NOT NULL DEFAULT '0' COMMENT '收藏的文章数',
  `article_num` int(5) NOT NULL DEFAULT '0' COMMENT '发布的数量',
  `fans_num` int(8) NOT NULL DEFAULT '0' COMMENT '粉丝数量',
  PRIMARY KEY (`id`,`user_id`,`open_id`,`phone`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_mall_info`
--

LOCK TABLES `user_mall_info` WRITE;
/*!40000 ALTER TABLE `user_mall_info` DISABLE KEYS */;
INSERT INTO `user_mall_info` VALUES (21,4315836759645566,'乌蝇哥',1,'','1583675964',0,'oH0FI41fuEa1JxI0bOCq8rq9chUs',0,0,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','Haidian','Beijing','',4,6,0),(22,61583579141,'yy',0,'','1583675964',0,'oH0FI41fuEa1JxI0bOCq8rq9chUs',0,0,'https://wx.qlogo.cn/mmopen/vi_32/V3ZJ9F6NibTfyBcYbBNnztq65woAPhiaehWIdQnpTcib9bIN1icyEcydh4t6y7oAztL2icxyiawRnhpL8uqd9fvjFq0g/132','','','',0,0,0);
/*!40000 ALTER TABLE `user_mall_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-19 16:09:56
