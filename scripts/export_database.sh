#!/bin/bash

# Создаем директорию для дампа если её нет
mkdir -p scripts/init-data


mysqldump -h localhost -u root -p currency_db > scripts/init-data/dump.sql

echo "База данных экспортирована в scripts/init-data/dump.sql"
