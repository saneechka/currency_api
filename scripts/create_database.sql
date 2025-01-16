CREATE DATABASE IF NOT EXISTS currency_db;

USE currency_db;

DROP TABLE IF EXISTS exchange_rates;

CREATE TABLE IF NOT EXISTS exchange_rates (
    Cur_ID           INT NOT NULL,
    Date             DATE NOT NULL,  
    Cur_Abbreviation VARCHAR(3) NOT NULL,
    Cur_Scale        INT NOT NULL,
    Cur_Name         VARCHAR(100) NOT NULL,
    Cur_OfficialRate DECIMAL(10,4) NOT NULL,
    PRIMARY KEY (Cur_ID, Date)
);

