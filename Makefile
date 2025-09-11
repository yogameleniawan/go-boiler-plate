# Name of the binary output
BINARY_NAME=main

# Docker Compose file
COMPOSE_FILE=docker-compose.yml
DEV_COMPOSE_FILE=docker-compose.dev.yaml

# migration dir
MIGRATE_DIR=db/migrations

# install dipendencies
init:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### up container
up:
	docker-compose -f $(COMPOSE_FILE) up --build

### shutdown container
down:
	docker-compose -f $(COMPOSE_FILE) down

### Run Only This App
run:
	go run ./cmd/api/main.go

### Build
build:
	go build -o $(BINARY_NAME) ./cmd/api/main.go

### create migration file
create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create-migration name=create_users_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)

### Dev with environment variables support
dev:
	docker-compose -f $(DEV_COMPOSE_FILE) up --build

### Dev down
dev-down:
	docker-compose -f $(DEV_COMPOSE_FILE) down

### Dev with clean volumes
dev-clean:
	docker-compose -f $(DEV_COMPOSE_FILE) down -v
	docker-compose -f $(DEV_COMPOSE_FILE) up --build

### Help
help:
	@echo "Available commands:"
	@echo "  init                     - Install dependencies"
	@echo "  up                       - Run production containers"
	@echo "  down                     - Stop production containers"
	@echo "  run                      - Run app locally"
	@echo "  build                    - Build binary"
	@echo "  create-migration name=X  - Create migration file"
	@echo "  dev                      - Run development containers"
	@echo "  dev-down                 - Stop development containers"
	@echo "  dev-clean                - Clean dev containers and volumes"
	@echo "  help                     - Show this help"

.PHONY: init up down run build create-migration dev dev-down dev-clean show-config help