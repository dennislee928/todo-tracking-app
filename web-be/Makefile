.PHONY: run build swagger proto migrate test

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go -o docs --parseInternal --parseDependency

proto:
	docker run --rm -v $(PWD):/app -w /app bufbuild/buf:latest generate

migrate-up:
	migrate -path database/migrations -database "$${DATABASE_URL}" up

migrate-down:
	migrate -path database/migrations -database "$${DATABASE_URL}" down

test:
	go test ./...

tidy:
	go mod tidy
