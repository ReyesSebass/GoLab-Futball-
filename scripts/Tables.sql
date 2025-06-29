CREATE DATABASE football;
USE football;

CREATE TABLE leagues (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name_league VARCHAR(100) NOT NULL,
    logo_league VARCHAR(255)
);