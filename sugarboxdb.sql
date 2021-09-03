CREATE DATABASE  IF NOT EXISTS `sugarboxdb` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `sugarboxdb`;
-- MySQL dump 10.13  Distrib 5.7.12, for Win32 (AMD64)
--
-- Host: localhost    Database: sugarboxdb
-- ------------------------------------------------------
-- Server version	5.7.35-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tblmovieall`
--

DROP TABLE IF EXISTS `tblmovieall`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tblmovieall` (
  `idx` int(11) NOT NULL AUTO_INCREMENT,
  `typeName` varchar(255) NOT NULL,
  `seltype` varchar(255) DEFAULT NULL,
  `directorName` varchar(255) DEFAULT NULL,
  `artist` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idx`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tblmovieall`
--

LOCK TABLES `tblmovieall` WRITE;
/*!40000 ALTER TABLE `tblmovieall` DISABLE KEYS */;
INSERT INTO `tblmovieall` VALUES (1,'Love','movie','Suresh','Salman'),(2,'Maachis','movie','Gulzar','Chandrachur'),(3,'Cheer Haran','documentary','Kuldeep','Ruhil'),(4,'Coolie','movie','Manmohan Desai','Amit'),(5,'Ghulami','movie','J. P. Dutta','Dada'),(6,'Kaagaz','movie','Satish','Pankaj Tripathi'),(7,'Jamun','movie','Gaurav ','Raghubir'),(8,'Machaan','movie','Nitesh','Pawan'),(9,'Sakshi','movie','Anay','Vikram'),(10,'Roohi','movie','Hardik','Rajkummar');
/*!40000 ALTER TABLE `tblmovieall` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tblmoviealltrans`
--

DROP TABLE IF EXISTS `tblmoviealltrans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tblmoviealltrans` (
  `idx` int(11) NOT NULL AUTO_INCREMENT,
  `typeId` int(11) NOT NULL,
  `userId` int(11) DEFAULT NULL,
  `rating` int(11) DEFAULT NULL,
  `userComment` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`idx`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tblmoviealltrans`
--

LOCK TABLES `tblmoviealltrans` WRITE;
/*!40000 ALTER TABLE `tblmoviealltrans` DISABLE KEYS */;
INSERT INTO `tblmoviealltrans` VALUES (1,2,2,4,'Given 4'),(2,2,1,2,'Given 1'),(15,4,1,3,''),(16,2,3,4,'testing......');
/*!40000 ALTER TABLE `tblmoviealltrans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbluser`
--

DROP TABLE IF EXISTS `tbluser`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tbluser` (
  `idx` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(70) NOT NULL,
  `upass` varchar(100) DEFAULT NULL,
  `urole` int(2) unsigned zerofill DEFAULT NULL,
  `ufname` varchar(30) DEFAULT NULL,
  `ulname` varchar(30) DEFAULT NULL,
  `ustatus` bit(1) DEFAULT b'0',
  PRIMARY KEY (`idx`,`uid`),
  UNIQUE KEY `uid_UNIQUE` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbluser`
--

LOCK TABLES `tbluser` WRITE;
/*!40000 ALTER TABLE `tbluser` DISABLE KEYS */;
INSERT INTO `tbluser` VALUES (1,'test1','123456',NULL,'Test1',NULL,'\0'),(2,'test2','123456',NULL,'Test2',NULL,'\0'),(3,'test3','123456',NULL,'Test3',NULL,'\0'),(4,'admin','admin',NULL,'Admin',NULL,'\0');
/*!40000 ALTER TABLE `tbluser` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-03 23:48:16
