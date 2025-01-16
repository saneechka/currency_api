#!/bin/bash

echo "Выберите режим запуска:"
echo "1. Использовать существующую базу данных"
echo "2. Создать новую базу данных с данными из API"
read -p "Введите номер (1 или 2): " choice

case $choice in
  1)
    echo "Экспорт существующей базы данных..."
    ./scripts/export_database.sh
    echo "Запуск с существующей базой данных..."
    docker-compose -f docker-compose.existing-db.yml up -d
    ;;
  2)
    echo "Запуск с новой базой данных..."
    docker-compose -f docker-compose.new-db.yml up -d
    ;;
  *)
    echo "Неверный выбор"
    exit 1
    ;;
esac
