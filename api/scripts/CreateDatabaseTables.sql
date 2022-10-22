CREATE DATABASE IF NOT EXISTS guposbf;
USE guposbf;

DROP TABLE IF EXISTS products;

CREATE TABLE products(
    id int auto_increment primary key,
    product_name varchar(100) not null,
    product_price decimal(10,2)
)ENGINE=INNODB;

DROP TABLE IF EXISTS currencies;

CREATE TABLE currencies(
    id int auto_increment primary key,
    symbol varchar(10) not null,
    country varchar(100) not null
)ENGINE=INNODB;