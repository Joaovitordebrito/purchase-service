CREATE DATABASE IF NOT EXISTS wexdb;
USE wexdb;

DROP TABLE IF EXISTS purchase;

CREATE TABLE purchase(
    uuid VARCHAR(36) NOT NULL PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    purchaseAmount FLOAT NOT NULL,
    transactionDate VARCHAR(10)
) ENGINE=INNODB;