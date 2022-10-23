CREATE DATABASE IF NOT EXISTS gruposbf;
USE gruposbf;

DROP TABLE IF EXISTS currencies;

CREATE TABLE currencies(
    id int auto_increment primary key,
    symbol varchar(10) not null,
    country varchar(100) not null
)ENGINE=INNODB;