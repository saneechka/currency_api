# Сервис Обмена Валют



## 🚀 Быстрый Старт

### Использование Docker (Рекомендуется)
```bash
# Клонирование репозитория
git clone https://github.com/saneechka/currency-service.git
cd currency-service

# Запуск сервисов
docker-compose up -d
```

### Ручная Установка

1. Установка MySQL:
```bash
# Ubuntu
sudo apt update && sudo apt install mysql-server

# Arch Linux
sudo pacman -Sy && sudo pacman -S mariadb

# macOS
brew install mysql
```

2. Настройка MySQL:
```bash
# Запуск службы MySQL
sudo systemctl start mysql   # Linux
sudo systemctl enable mysql  # Linux
brew services start mysql    # macOS

# Безопасная установка MySQL
sudo mysql_secure_installation

# Следуйте подсказкам:
# - Установите пароль root
# - Удалите анонимных пользователей
# - Запретите удаленный вход root
# - Удалите тестовую базу данных
# - Перезагрузите таблицы привилегий
```

3. Проверка MySQL:
```bash
# Проверка статуса MySQL
systemctl status mysql      # Linux
brew services list         # macOS

# Проверка подключения
mysql -u root -p -e "SELECT VERSION();"
```

4. Настройка базы данных:
```sql
CREATE DATABASE currency_db;
USE currency_db;

CREATE TABLE exchange_rates (
    Cur_ID           INT NOT NULL,
    Date             DATE NOT NULL,  
    Cur_Abbreviation VARCHAR(3) NOT NULL,
    Cur_Scale        INT NOT NULL,
    Cur_Name         VARCHAR(100) NOT NULL,
    Cur_OfficialRate DECIMAL(10,4) NOT NULL,
    PRIMARY KEY (Cur_ID, Date)
);
```

5. Настройка окружения:
```bash
# Редактирование файла .env
PORT=8080
DATABASE_DSN=root:your_password@tcp(localhost:3306)/currency_db?parseTime=true
```

6. Запуск приложения:
```bash
go run cmd/main.go
```

## 📖 API Документация

### Эндпоинты

- `GET /api/rates` - Получить текущие курсы валют
- `GET /api/rates?date=YYYY-MM-DD` - Получить курсы за определённую дату

### Пример запроса
```bash
curl http://localhost:8080/api/rates?date=2025-01-15
```

### Пример ответа
```json
[
  {
    "Cur_ID": 431,
    "Date": "2025-01-15",
    "Cur_Abbreviation": "USD",
    "Cur_Scale": 1,
    "Cur_Name": "Доллар США",
    "Cur_OfficialRate": 3.2345
  }
]
```### Сборка проекта
```bash
# Полная сборка с проверками
make all

# Только сборка
make build

# Сборка и запуск
make run

# Запуск в Docker
make docker

# Очистка
make clean
```
```

### Swagger Документация

Для просмотра интерактивной API документации:

```bash
# Запуск Swagger UI через Docker
docker run -p 8081:8080 -e SWAGGER_JSON=/api/openapi.yaml -v ${PWD}/api:/api swaggerapi/swagger-ui
```

Откройте http://localhost:8081 в браузере для просмотра документации.

Альтернативные способы:

1. Онлайн редактор:
   - Перейдите на https://editor.swagger.io/
   - Импортируйте файл `api/openapi.yaml`

2. Локальная установка:
```bash
# Установка swagger-cli
npm install -g swagger-cli

# Валидация документации
swagger-cli validate api/openapi.yaml

# Сборка документации
swagger-cli bundle api/openapi.yaml > api/swagger.json
```

## 🛠 Разработка

### Линтинг
```bash
# Установка линтера (только первый раз)
make lint-install

# Запуск проверки кода
make lint

# Автоматическое исправление проблем (где возможно)
make lint-fix
```

### Тестирование
```bash
# Запуск всех тестов
go test ./...

# Запуск конкретного теста
go test ./internal/repository -v -run TestGetRates
```

### Генерация API клиента
```bash
# Установка openapi-generator
npm install @openapitools/openapi-generator-cli -g

# Генерация TypeScript клиента
openapi-generator generate -i api/openapi.yaml -g typescript-fetch -o generated/typescript
```



## 🔍 Мониторинг и Устранение Проблем

### Просмотр логов
```bash
# Логи приложения
docker-compose logs -f app

# Логи базы данных
docker-compose logs -f db
```

### Частые проблемы

1. **Сервис не запускается:**
   ```bash
   docker-compose logs app
   docker-compose restart app
   ```

2. **Проблемы с подключением к БД:**
   ```bash
   docker-compose exec db mysqladmin -u root -p status
   ```

## 📚 Дополнительные Ресурсы

- [Документация API](./api/openapi.yaml)
- [Схема базы данных](./scripts/create_database.sql)

## 🛑 Остановка Сервиса
```bash
docker-compose down
```

---

# Currency Exchange Rates Service

## 🚀 Quick Start

### Using Docker (Recommended)
```bash
# Clone repository
git clone https://github.com/saneechka/currency-service.git
cd currency-service

# Start services
docker-compose up -d
```

### Manual Installation

1. Install MySQL:
```bash
# Ubuntu
sudo apt update && sudo apt install mysql-server

# Arch Linux
sudo pacman -Sy && sudo pacman -S mariadb

# macOS
brew install mysql
```

2. Configure MySQL:
```bash
# Start MySQL service
sudo systemctl start mysql   # Linux
sudo systemctl enable mysql  # Linux
brew services start mysql    # macOS

# Secure MySQL installation
sudo mysql_secure_installation

# Follow the prompts:
# - Set root password
# - Remove anonymous users
# - Disallow root login remotely
# - Remove test database
# - Reload privilege tables
```

3. Verify MySQL Setup:
```bash
# Check MySQL status
systemctl status mysql      # Linux
brew services list         # macOS

# Test connection
mysql -u root -p -e "SELECT VERSION();"
```

4. Configure Database:
```sql
CREATE DATABASE currency_db;
USE currency_db;

CREATE TABLE exchange_rates (
    Cur_ID           INT NOT NULL,
    Date             DATE NOT NULL,  
    Cur_Abbreviation VARCHAR(3) NOT NULL,
    Cur_Scale        INT NOT NULL,
    Cur_Name         VARCHAR(100) NOT NULL,
    Cur_OfficialRate DECIMAL(10,4) NOT NULL,
    PRIMARY KEY (Cur_ID, Date)
);
```

5. Configure Environment:
```bash
# Edit .env file
PORT=8080
DATABASE_DSN=root:your_password@tcp(localhost:3306)/currency_db?parseTime=true
```

6. Run Application:
```bash
go run cmd/main.go
```

## 📖 API Documentation

### Endpoints

- `GET /api/rates` - Get current exchange rates
- `GET /api/rates?date=YYYY-MM-DD` - Get rates for specific date

### Example Request
```bash
curl http://localhost:8080/api/rates?date=2025-01-15
```

### Example Response
```json
[
  {
    "Cur_ID": 431,
    "Date": "2025-01-15",
    "Cur_Abbreviation": "USD",
    "Cur_Scale": 1,
    "Cur_Name": "US Dollar",
    "Cur_OfficialRate": 3.2345
  }
]
```

### Swagger Documentation

To view interactive API documentation:

```bash
# Run Swagger UI via Docker
docker run -p 8081:8080 -e SWAGGER_JSON=/api/openapi.yaml -v ${PWD}/api:/api swaggerapi/swagger-ui
```

Open http://localhost:8081 in your browser to view the documentation.

Alternative methods:

1. Online editor:
   - Go to https://editor.swagger.io/
   - Import the `api/openapi.yaml` file

2. Local installation:
```bash
# Install swagger-cli
npm install -g swagger-cli

# Validate documentation
swagger-cli validate api/openapi.yaml

# Bundle documentation
swagger-cli bundle api/openapi.yaml > api/swagger.json
```

## 🛠 Development

### Linting
```bash
# Install linter
make lint-install

# Run lint check
make lint
```

### Testing
```bash
# Run all tests
go test ./...

# Run specific test
go test ./internal/repository -v -run TestGetRates
```

### API Client Generation
```bash
# Install openapi-generator
npm install @openapitools/openapi-generator-cli -g

# Generate TypeScript client
openapi-generator generate -i api/openapi.yaml -g typescript-fetch -o generated/typescript
```

### Project Build
```bash
# Full build with checks
make all

# Build only
make build

# Build and run
make run

# Run in Docker
make docker

# Clean up
make clean
```

## 🔍 Monitoring & Troubleshooting

### View Logs
```bash
# Application logs
docker-compose logs -f app

# Database logs
docker-compose logs -f db
```

### Common Issues

1. **Service won't start:**
   ```bash
   docker-compose logs app
   docker-compose restart app
   ```

2. **Database connection issues:**
   ```bash
   docker-compose exec db mysqladmin -u root -p status
   ```

## 📚 Additional Resources

- [API Documentation](./api/openapi.yaml)
- [Database Schema](./scripts/create_database.sql)

## 🛑 Stopping the Service
```bash
docker-compose down
```


