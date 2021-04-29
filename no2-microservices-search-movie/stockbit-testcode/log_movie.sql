CREATE DATABASE movie;
use movie;

CREATE TABLE `movie`.`log_movie`  (
  `ID` int(0) NOT NULL AUTO_INCREMENT,
  `ImdbID` varchar(50)  NULL DEFAULT NULL,  /*Digunakan untuk menyimpan id dari movie*/
  `CreatedAt` timestamp(0) NULL DEFAULT NULL,
  `UpdatedAt` timestamp(0) NULL DEFAULT NULL,
  `path` varchar(10)  NULL DEFAULT NULL, /*Digunakan untuk menyimpan sumber log ketika seach atau klik detail*/
  PRIMARY KEY (`ID`) USING BTREE
);