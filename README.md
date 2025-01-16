# Сервис Обмена Валют | Currency Exchange Rates Service

Простой и надёжный сервис для отслеживания курсов валют.

## 🚀 Быстрый Старт | Quick Start

### Использование Docker (Рекомендуется) | Using Docker (Recommended)
```bash
# Клонирование репозитория | Clone repository
git clone https://github.com/saneechka/currency-service.git
cd currency-service

# Запуск сервисов | Start services
docker-compose up -d
```

### Ручная Установка | Manual Installation

1. Установка MySQL | Install MySQL:
```bash
# Debian
sudo apt update && sudo apt install mysql-server

# Arch based distro
sudo pacman -Sy && sudo pacman -S mariadb

# macOS
brew install mysql
```

2. Настройка MySQL | Configure MySQL:
```bash
# Запуск службы Linux| Start service
sudo systemctl start mysql   # Linux
sudo systemctl enable mysql  # Linux
brew services start mysql    # macOS

# Безопасная установка | Secure installation
sudo mysql_secure_installation
```

3. Настройка базы данных | Configure Database:
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

4. Настройка окружения | Configure Environment:
```bash
# Редактирование файла .env | Edit .env file
PORT=8080
DATABASE_DSN=root:your_password@tcp(localhost:3306)/currency_db?parseTime=true
```

## 🛠 Разработка | Development

### Сборка проекта | Project Build
```bash
# Полная сборка с проверками | Full build with checks
make all

# Сборка и запуск | Build and run
make run

# Запуск в Docker | Run in Docker
make docker

# Очистка | Clean up
make clean
```

### Линтинг | Linting
```bash
# Установка линтера | Install linter
make lint-install

# Запуск проверки | Run check
make lint

# Автоисправление | Auto-fix
make lint-fix
```

### Тестирование | Testing
```bash
# Запуск всех тестов | Run all tests
go test ./...

# Запуск конкретного теста | Run specific test
go test ./internal/repository -v -run TestGetRates
```

## 📖 API Documentation

### Endpoints
- `GET /api/rates` - Получить все курсы валют | Get all exchange rates
- `GET /api/rates?date=YYYY-MM-DD` - Получить курсы за дату | Get rates for date

### Swagger Documentation
```bash
# Запуск Swagger UI | Run Swagger UI
docker run -p 8081:8080 -e SWAGGER_JSON=/api/openapi.yaml -v ${PWD}/api:/api swaggerapi/swagger-ui
```

## 🔍 Мониторинг | Monitoring

### Просмотр логов | View Logs
```bash
# Логи приложения | Application logs
docker-compose logs -f app

# Логи базы данных | Database logs
docker-compose logs -f db
```

## 📚 Дополнительные Ресурсы | Additional Resources
- [API Documentation](./api/openapi.yaml)
- [Database Schema](./scripts/create_database.sql)

## 🛑 Остановка Сервиса | Stop Service
```bash
docker-compose down
```


