# Currency Exchange Rates Service

Simple and reliable service for tracking currency exchange rates.

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

2. Configure Database:
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

3. Configure Environment:
```bash
# Edit .env file
PORT=8080
DATABASE_DSN=root:your_password@tcp(localhost:3306)/currency_db?parseTime=true
```

4. Run Application:
```bash
go run cmd/main.go
```

## 📖 API Documentation

### Endpoints

- `GET /api/rates` - Get current exchange rates
- `GET /api/rates?date=YYYY-MM-DD` - Get rates for specific date

### Example Request
```bash
curl http://localhost:8080/api/rates?date=2024-03-14
```

### Example Response
```json
[
  {
    "Cur_ID": 431,
    "Date": "2024-03-14",
    "Cur_Abbreviation": "USD",
    "Cur_Scale": 1,
    "Cur_Name": "US Dollar",
    "Cur_OfficialRate": 3.2345
  }
]
```

## 🛠 Development

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


