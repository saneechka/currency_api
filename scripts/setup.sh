#!/bin/bash

# Получаем конфигурацию MySQL
./scripts/get_mysql_config.sh

# Создаем .env файл с правильными настройками
cat > .env << EOF
PORT=8080
DATABASE_DSN=${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(db:3306)/currency_db?parseTime=true
EOF

# Делаем скрипты исполняемыми
chmod +x scripts/*.sh

# Запускаем docker-compose
docker-compose up -d

echo "Waiting for database to initialize..."
sleep 10

echo "Setup complete! Your application is running with your MySQL configuration."
