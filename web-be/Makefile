.PHONY: run build swagger proto migrate test

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go -o docs --parseInternal --parseDependency

proto:
	@which protoc >/dev/null || (echo "Install protoc: brew install protobuf" && exit 1)
	@which protoc-gen-go >/dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc >/dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc --go_out=api/grpc/proto --go_opt=paths=source_relative \
		--go-grpc_out=api/grpc/proto --go-grpc_opt=paths=source_relative \
		api/grpc/proto/todo.proto

migrate-up:
	migrate -path database/migrations -database "$${DATABASE_URL}" up

migrate-down:
	migrate -path database/migrations -database "$${DATABASE_URL}" down

test:
	go test ./...

tidy:
	go mod tidy
