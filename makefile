# Makefile

# Load environment variables from .env file
include .env
export

# Paths
MIGRATE_CMD = $(shell which migrate)
MIGRATION_DIR = ./cmd/migrate/migrations

# Print database connection details for debugging
print-db:
	@echo "DB_USER: $(DB_USER_NAME)"
	@echo "DB_PASS: $(DB_PASS)"
	@echo "DB_HOST: $(DB_HOST)"
	@echo "DB_PORT: $(DB_PORT)"
	@echo "DB_NAME: $(DB_NAME)"

# Build and run commands
run:
	@go run cmd/main.go

build:
	@go build -o bin/auth cmd/main.go

buildrun: build
	@./bin/auth

test:
	@go test -v ./...

# Migration commands
migration:
	@echo "Creating new migration..."
	@$(MIGRATE_CMD) create -ext sql -dir $(MIGRATION_DIR) -seq $(name)
	@echo "Migration created successfully."

migration-up:
	@echo "Running migrations up..."
	@$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "mysql://$(DB_USER_NAME):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up
	@echo "Migrations applied successfully."

migration-down:
	@echo "Running migrations down..."
	@$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "mysql://$(DB_USER_NAME):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down
	@echo "Migrations rolled back successfully."

migration-force:
	@echo "Forcing migration version..."
	@$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database "mysql://$(DB_USER_NAME):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" force 1
	@echo "Migration version forced successfully."
