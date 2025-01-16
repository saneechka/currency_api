# Currency Exchange Rates Service

## Варианты запуска



## Запуск на новом устройстве

###  Установка Docker
```bash
# Ubuntu
sudo apt update
sudo apt install docker.io docker-compose

# Arch-based
sudo pacman -Sy
sudo pacman -S docker docker-compose

# MacOS
brew install docker docker-compose

```

###  Клонирование репозитория
```bash

git clone https://github.com/saneechka/currency-service.git
cd currency-service
```

### Запуск в Docker

```bash
# Запуск всего API в Docker
## Дополнительная информация

- [Документация API](./api/openapi.yaml)
- [Схема базы данных](./scripts/create_database.sql)



docker-compose up -d
```
```


## Локальная установка MySQL

## Дополнительная информация

- [Документация API](./api/openapi.yaml)
- [Схема базы данных](./scripts/create_database.sql)



### Linux(Arch-based)
```bash
# Установка MySQL
sudo pacman -Sy
sudo pacman -S mariadb

# Запуск MySQL
## Дополнительная информация

- [Документация API](./api/openapi.yaml)
- [Схема базы данных](./scripts/create_database.sql)


```bash
sudo systemctl start mariadb
sudo systemctl enable mariadb
```
# Настройка безопасности
sudo mysql_secure_installation
```

### Настройка базы данных

1. Подключение к MySQL:
```bash
# Linux/MacOS
mysql -u root -p
```

2. Создание базы данных:
```sql
CREATE DATABASE currency_db;
```

3. Создание таблицы:
```sql
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

### Настройка приложения для локальной БД

1. Измените .env файл:
```
PORT=8080
DATABASE_DSN=root:your_password@tcp(localhost:3306)/currency_db?parseTime=true
```

2. Запуск приложения:
```bash
go run cmd/main.go
```

###  Проверка работы

1. Проверьте статус контейнеров:
```bash
docker-compose ps
# Должны быть два контейнера в статусе "Up"
```

2. Проверьте работу API:
```bash
# Получение текущих курсов
curl http://localhost:8080/api/rates

# Получение курсов за дату
curl http://localhost:8080/api/rates?date=2024-03-14
```

### Устранение проблем

Если что-то пошло не так:

1. Проверьте логи:
```bash
# Логи приложения
docker-compose logs app

# Логи базы данных
docker-compose logs db
```

2. Перезапустите сервисы:
```bash
docker-compose down
docker-compose up -d
```

3. Очистите всё и начните заново:
```bash
# Полная очистка
docker-compose down -v
docker system prune -a

# Повторный запуск
docker-compose up -d
```

### Остановка сервиса
```bash
docker-compose down
```

## Дополнительная информация

- [Документация API](./api/openapi.yaml)
- [Схема базы данных](./scripts/create_database.sql)



## Проверка работы

```bash
# Статус сервисов
docker-compose ps


```

## Остановка сервиса

```bash
docker-compose down
```

## Дополнительная информация

### Логи

```bash
# Просмотр логов приложения
docker-compose logs -f app

# Просмотр логов базы данных
docker-compose logs -f db
```


### Возможные проблемы

1. Если порт 3306 занят:
   - Измените порт в docker-compose.yml
   - Или остановите локальный MySQL

2. Если порт 8080 занят:
   - Измените порт в docker-compose.yml для сервиса 


### Проверка работы

1. Проверка статуса MySQL:
```bash
# Linux/MacOS
systemctl status mysql

# MacOS 
brew services list
```

2. Проверка подключения:
```bash
mysql -u root -p -e "SELECT VERSION();"
```

3. Проверка базы данных:
```bash
mysql -u root -p currency_db -e "SHOW TABLES;"
```

## API Endpoints

- `GET /api/rates` - текущие курсы валют
- `GET /api/rates?date=YYYY-MM-DD` - курсы за дату

## API Documentation

API документация доступна в формате OpenAPI (Swagger) спецификации.

### Просмотр документации

1. Онлайн просмотр (рекомендуется):
```bash
# Установка Swagger UI
docker run -p 8081:8080 -e SWAGGER_JSON=/api/openapi.yaml -v ${PWD}/api:/api swaggerapi/swagger-ui
```
Откройте http://localhost:8081 в браузере.

2. Через Swagger Editor:
- Откройте https://editor.swagger.io/
- Импортируйте файл api/openapi.yaml

### Доступные endpoint'ы

#### GET /api/rates
Получение курсов валют

```bash
# Текущие курсы
curl http://localhost:8080/api/rates

# Курсы за дату
curl http://localhost:8080/api/rates?date=2024-03-14
```

Пример ответа:
```json
[
  {
    "Cur_ID": 431,
    "Date": "2024-03-14",
    "Cur_Abbreviation": "USD",
    "Cur_Scale": 1,
    "Cur_Name": "Доллар США",
    "Cur_OfficialRate": 3.2345
  }
]
```

### Генерация клиентов

Вы можете сгенерировать клиенты для различных языков используя openapi-generator:

```bash
# Установка openapi-generator
npm install @openapitools/openapi-generator-cli -g

# Генерация клиента для TypeScript
openapi-generator generate -i api/openapi.yaml -g typescript-fetch -o generated/typescript



## Мониторинг

### Логи

```bash
# Просмотр логов приложения
docker-compose logs -f app

# Просмотр логов базы данных
docker-compose logs -f db
```

### Статус сервисов

```bash
# Проверка статуса всех сервисов
docker-compose ps

# Проверка использования ресурсов
docker stats
```

## Устранение неполадок

### Распространенные проблемы

1. **Сервис не запускается:**
```bash
# Проверка логов
docker-compose logs app

# Перезапуск сервиса
docker-compose restart app
```

2. **Проблемы с базой данных:**
```bash
# Проверка статуса MySQL
docker-compose exec db mysqladmin -u root -p status

# Пересоздание базы данных
docker-compose down -v
docker-compose up -d
```

## Разработка

### Локальное окружение

```bash
# Сборка без кэша
docker-compose build --no-cache

# Запуск с отладкой
## Дополнительная информация

- [Документация API](./api/openapi.yaml)
- [Схема базы данных](./scripts/create_database.sql)



docker-compose --profile debug up
```

### Тестирование

```bash
# Запуск всех тестов
go test ./...

# Запуск конкретного теста
go test ./internal/repository -v -run TestGetRates
```


