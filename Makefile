PROJECT_NAME=grpc-clean
MODULE_NAME=grpc-clean

.DEFAULT_GOAL := build

.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: run
run:
	@go run cmd/grpc-clean/main.go

.PHONY: build
build:
	@go build ./cmd/grpc-clean/

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html=coverage.out

.PHONY: get
get:
	@go mod download