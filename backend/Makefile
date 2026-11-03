.PHONY: wire docs build up down run test fmt clean update

wire:
	@echo "==> Generating Wire..."
	@go generate ./cmd/server

docs:
	@echo "==> Generating Swagger docs..."
	@swag init -g main.go --dir ./cmd/server,./internal --output ./docs

build:
	@echo "==> Building application..."
	@go build -o ./bin/server.exe ./cmd/server

up:
	@echo "==> Building and starting API..."
	@docker compose up -d --build

down:
	@echo "==> Stopping services..."
	@docker compose down

run:
	@echo "==> Starting server locally..."
	@go run ./cmd/server

test:
	@echo "==> Running tests..."
	@go test ./...

fmt:
	@echo "==> Formatting Go code..."
	@gofmt -w ./cmd ./internal

clean:
	@echo "==> Cleaning build artifacts..."
	@if exist bin rmdir /s /q bin

update: docs wire
	@echo "==> Update completed successfully."
