CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE IF NOT EXISTS customers (id serial, name varchar(255), email varchar(255));

CALL DOLT_COMMIT('-Am', 'create table customers');

INSERT INTO customers(name, email) VALUES ('John', 'john@gmail.com');

CALL DOLT_COMMIT('-Am', 'insert john into customers');
