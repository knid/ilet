all: test build run

build:
	@go build -o builds/main cmd/api.go

run: build
	@./builds/main

test:
	go test -v ./...
