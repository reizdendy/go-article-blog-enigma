CREATE DATABASE  IF NOT EXISTS `db_article_blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_article_blog`;
-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: db_article_blog
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `m_article`
--

DROP TABLE IF EXISTS `m_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_article` (
  `article_id` varchar(36) NOT NULL,
  `article_name` varchar(45) DEFAULT NULL,
  `article_content` text,
  `user_creator` varchar(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_article`
--

LOCK TABLES `m_article` WRITE;
/*!40000 ALTER TABLE `m_article` DISABLE KEYS */;
INSERT  IGNORE INTO `m_article` VALUES ('5294730c-e20d-471f-8eba-040777eb7257','Artikel keempat saya','Disini saya akan mencoba mengedit sebuah artikel dengan judul ...........','aden','2020-07-07 12:48:22'),('52a7b56d-df62-4346-a582-ab09bc10790d','Artikel ketiga saya','Disini saya akan membuat sebuah artikel dengan judul ...........','aden','2020-07-07 11:55:29'),('5e3e03f6-c002-45de-9fdf-982f9f777793','Artikel Pertama Saya','Disini saya akan membuat sebuah artikel ...........','aden','2020-07-07 11:00:34'),('6c37cfa3-1305-4a64-84cb-dafa32644784','Artikel kedua saya','Disini saya akan membuat sebuah artikel dengan judul ...........','aden','2020-07-07 11:53:26'),('f5cdd371-b10c-497b-b8ff-f1513ee628f6','Artikel kelima saya','Ini adalah artikel kelima saya dengan judul ...........','aden','2020-07-07 13:02:05'),('f872b633-72d7-4428-8c1d-035abfc5e948','Artikel keenam saya','Ini adalah artikel keenam saya dengan judul ...........','aden','2020-07-07 13:10:56');
/*!40000 ALTER TABLE `m_article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_user`
--

DROP TABLE IF EXISTS `m_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_user` (
  `user_name` varchar(20) NOT NULL,
  `password` varchar(60) NOT NULL,
  `email` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`user_name`,`password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_user`
--

LOCK TABLES `m_user` WRITE;
/*!40000 ALTER TABLE `m_user` DISABLE KEYS */;
INSERT  IGNORE INTO `m_user` VALUES ('aden','rahasia','aden@gmail.com'),('guest','$2a$10$xQ9SfVT/stnnbUn4c2pfjO2mSZ.5Qx8I4Hdjm2e8OOaEUf9/FLrxq','guest@gmail.com'),('userguest','$2a$10$3na82h.IaYhvnBGEulYmpOF393HQg9Xc9aHClbKwzd62lhhSf66fm','guest@gmail.com');
/*!40000 ALTER TABLE `m_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-07 13:11:22
