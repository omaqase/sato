# Makefile для Windows (CMD/PowerShell)
# Использование: make [цель] [ПАРАМЕТРЫ]

ifneq (,$(wildcard .env))
    include .env
    export
endif

MIGRATE_CMD = migrate
MIGRATIONS_DIR = ./db/postgres/migrations
MIGRATE_URL ?= $(POSTGRES_CONNECTION_URL)

.PHONY: help migrate-up migrate-down migrate-down-all migrate-force migrate-create migrate-status migrate-version

help:
	@echo Управление миграциями:
	@echo   migrate-up [steps=N]     - Применить миграции (N шагов или all)
	@echo   migrate-down [steps=N]   - Откатить миграции (по умолч. 1)
	@echo   migrate-down-all         - Откатить ВСЕ миграции
	@echo   migrate-force version=N  - Принудительно установить версию
	@echo   migrate-create name=NAME - Создать новую миграцию
	@echo   migrate-status           - Статус миграций
	@echo   migrate-version          - Текущая версия

migrate-up: _check-url
	@$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(MIGRATE_URL)" up $(or $(steps),all)

migrate-down: _check-url
	@$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(MIGRATE_URL)" down $(or $(steps),1)

migrate-down-all: _check-url
	@$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(MIGRATE_URL)" down all

migrate-force: _check-url
ifndef version
	$(error version is required)
endif
	@$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(MIGRATE_URL)" force $(version)

migrate-create:
ifndef name
	$(error name is required. Usage: make migrate-create name=NAME)
endif
	@$(MIGRATE_CMD) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-status: _check-url
	@$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(MIGRATE_URL)" status

migrate-version: _check-url
	@$(MIGRATE_CMD) -path $(MIGRATIONS_DIR) -database "$(MIGRATE_URL)" version

_check-url:
ifndef MIGRATE_URL
	$(error MIGRATE_URL is not set. Define it in .env file)
endif

.PHONY: test test-coverage

test:
	go test -v ./... -race

test-coverage:
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out