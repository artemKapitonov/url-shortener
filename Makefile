include .env
LOCAL_BIN:=$(CURDIR)/bin


install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

migrate-up:
	migrate -path ./migrations/schema -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./migrations/schema -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

generate:
	make generate-api

generate-api:
	mkdir -p pkg/url-shortner_v1
	protoc --proto_path api \
	--go_out=pkg --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/url-shortner_v1/url-shortner.proto


run:
	go build -o ./bin/url-shortner ./cmd/url-shortner
	./bin/url-shortner