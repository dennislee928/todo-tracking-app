.PHONY: run build swagger migrate test

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

swagger:
	swag init -g cmd/server/main.go -o docs --parseInternal --parseDependency

migrate-up:
	migrate -path database/migrations -database "$${DATABASE_URL}" up

migrate-down:
	migrate -path database/migrations -database "$${DATABASE_URL}" down

test:
	go test ./...

tidy:
	go mod tidy
