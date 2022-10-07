-- Active: 1665067269490@@127.0.0.1@3306@golangdb
# 创建数据库
CREATE  DATABASE IF NOT EXISTS golangdb;
USE golangdb;
SHOW TABLES;
-- # 创建用户表
-- CREATE TABLE users(

-- )
SET foreign_key_checks=0;  # 关闭外键检查

SHOW CREATE TABLE admins;

SELECT * FROM users;
