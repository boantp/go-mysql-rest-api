/*
SQLyog Ultimate v10.42 
MySQL - 5.5.5-10.1.28-MariaDB : Database - shopee
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`shopee` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `shopee`;

/*Table structure for table `order_details` */

DROP TABLE IF EXISTS `order_details`;

CREATE TABLE `order_details` (
  `order_details_id` bigint(20) NOT NULL,
  `orders_id` int(11) NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `tax_code_id` int(11) DEFAULT NULL,
  `amount` decimal(15,2) DEFAULT NULL,
  `tax_amount` decimal(15,2) DEFAULT NULL,
  `total_amount` decimal(15,2) DEFAULT NULL,
  `order_status` smallint(1) DEFAULT '0' COMMENT '0:draft order, 1:submit, 2:paid',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `store_id` int(11) NOT NULL,
  PRIMARY KEY (`order_details_id`),
  KEY `fk_order_item_order_header_1` (`orders_id`),
  CONSTRAINT `fk_order_item_order_header_1` FOREIGN KEY (`orders_id`) REFERENCES `orders` (`orders_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='number 1';

/*Data for the table `order_details` */

/*Table structure for table `orders` */

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
  `orders_id` int(11) NOT NULL,
  `store_id` int(11) DEFAULT NULL,
  `store_name` varchar(100) DEFAULT NULL,
  `total_amount` decimal(15,2) DEFAULT NULL,
  `total_tax_amount` decimal(15,2) DEFAULT NULL,
  `grand_total` decimal(15,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`orders_id`),
  KEY `fk_orders_store_1` (`store_id`),
  CONSTRAINT `fk_orders_store_1` FOREIGN KEY (`store_id`) REFERENCES `store` (`store_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `orders` */

/*Table structure for table `store` */

DROP TABLE IF EXISTS `store`;

CREATE TABLE `store` (
  `store_id` int(11) NOT NULL,
  `store_name` varchar(50) DEFAULT NULL,
  `status` smallint(1) DEFAULT NULL,
  PRIMARY KEY (`store_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `store` */

insert  into `store`(`store_id`,`store_name`,`status`) values (1,'Grand Indonesia',1),(2,'Plaza Indonesia',1);

/*Table structure for table `tax_code` */

DROP TABLE IF EXISTS `tax_code`;

CREATE TABLE `tax_code` (
  `tax_code_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`tax_code_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `tax_code` */

insert  into `tax_code`(`tax_code_id`,`name`,`status`,`created_at`,`updated_at`) values (1,'food',1,'2018-09-10 01:36:16','2018-09-10 01:36:16'),(2,'tobacco',1,'2018-09-10 01:36:17','2018-09-10 01:36:17'),(3,'entertainment',1,'2018-09-10 01:36:20','2018-09-10 01:36:20');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
