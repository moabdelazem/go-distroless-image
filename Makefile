include .env

PORT ?= 8080
APP_NAME = "Distroless API"

build:
	@go build -o ./bin/main ./cmd/main.go

docker-build:
	@docker build -t moabdelazem/my-configurable-api .

docker-run:
	@docker run -p $(PORT):$(PORT) -e SERVER_PORT=:$(PORT) -e APP_NAME=$(APP_NAME) moabdelazem/my-configurable-api
