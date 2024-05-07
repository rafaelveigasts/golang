create database if not exists devbook;

use devbook;

drop table if exists users;



create table users (
    id int auto_increment primary key,
    name varchar(100) not null,
    alias varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null
    created_at datetime default current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;