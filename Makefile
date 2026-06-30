.PHONY: build test lint run-cli

build:
	go build -o bin/server ./cmd/server
	go build -o bin/cli ./cmd/cli

test:
	go test -v -race ./...

lint:
	golangci-lint run ./...

run-cli:
	go run ./cmd/cli $(URL)