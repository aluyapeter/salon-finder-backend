
```makefile
.PHONY: run build lint tidy test

run:
 go run ./cmd/api

build:
 @mkdir -p bin
 go build -o bin/api ./cmd/api

lint:
 golangci-lint run

tidy:
 go mod tidy

test:
 go test ./...
