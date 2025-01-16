#!/bin/bash

echo "Waiting for MySQL to start..."
while ! mysqladmin ping -h"localhost" --silent; do
    sleep 1
done

echo "MySQL started, initializing database..."
mysql -u root -p"${MYSQL_ROOT_PASSWORD}" << EOF
CREATE DATABASE IF NOT EXISTS currency_db;
USE currency_db;

CREATE TABLE IF NOT EXISTS exchange_rates (
    Cur_ID           INT NOT NULL,
    Date             DATE NOT NULL,  
    Cur_Abbreviation VARCHAR(3) NOT NULL,
    Cur_Scale        INT NOT NULL,
    Cur_Name         VARCHAR(100) NOT NULL,
    Cur_OfficialRate DECIMAL(10,4) NOT NULL,
    PRIMARY KEY (Cur_ID, Date)
);

# Добавляем тестовые данные для проверки
INSERT IGNORE INTO exchange_rates VALUES 
    (431, CURDATE(), 'USD', 1, 'Доллар США', 3.2345),
    (451, CURDATE(), 'EUR', 1, 'Евро', 3.5678);

EOF

echo "Database initialization completed!"
