DROP DATABASE BD_GO;
CREATE DATABASE BD_GO;
USE BD_GO;


CREATE TABLE tb_alumnos(
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(20),
    lastName VARCHAR(20),
    age TINYINT,
    dni char(8) not null
);