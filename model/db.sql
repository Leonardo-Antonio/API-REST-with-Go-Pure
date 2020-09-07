DROP DATABASE BD_GO;
CREATE DATABASE BD_GO;
USE BD_GO;


CREATE TABLE tb_alumns(
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(20),
    lastName VARCHAR(20),
    age TINYINT,
    dni char(8) not null
);


CREATE TABLE tb_users(
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    email VARCHAR(20) not null ,
    pass VARCHAR(20) not null
);