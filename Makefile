SILENT:
PHONY:

include .env.prod
export


build:
	go build -o ./.bin/main ./cmd/main/main.go
run: build
	ENV_FILE=".env.prod" ./.bin/main

docker:
	docker-compose  up -d

docker-app: docker-build docker


compose-prod:
	docker compose -f docker-compose.prod.yaml up --build -d