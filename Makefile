.PHONY: all build run clean test lint lint-install docker docker-down

# Переменные сборки
BINARY_NAME=currency-service
GO_FILES=$(shell find . -name '*.go')
BUILD_DIR=bin
DOCKER_COMPOSE=docker-compose

# Основные команды
all: lint test build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/main.go

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)
	go clean
	docker-compose down -v

# Тестирование и линтинг
test:
	go test -v ./...

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	golangci-lint run --config=.golangci.yml

# Docker команды
docker:
	$(DOCKER_COMPOSE) up --build -d

docker-down:
	$(DOCKER_COMPOSE) down

# Вспомогательные команды
swagger:
	docker run -p 8081:8080 -e SWAGGER_JSON=/api/openapi.yaml -v ${PWD}/api:/api swaggerapi/swagger-ui

deps:
	go mod download
	go mod tidy

help:
	@echo "Доступные команды:"
	@echo "  make build      - собрать проект"
	@echo "  make run       - собрать и запустить"
	@echo "  make test      - запустить тесты"
	@echo "  make lint      - запустить линтер"
	@echo "  make docker    - запустить в Docker"
	@echo "  make clean     - очистить сборку"
