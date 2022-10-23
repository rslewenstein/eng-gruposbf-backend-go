CREATE DATABASE IF NOT EXISTS gruposbf;
USE gruposbf;

DROP TABLE IF EXISTS currencies;

CREATE TABLE currencies(
    id int auto_increment primary key,
    symbol varchar(10) not null,
    country varchar(100) not null
)ENGINE=INNODB;

CREATE TABLE currencylogs(
    id int auto_increment primary key,
    typelog varchar(20) not null,
    method varchar(200) not null,
    logdescription varchar(500) not null,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;