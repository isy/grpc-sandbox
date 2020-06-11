APP = user-service

build:
	go build -o build/${APP} ./cmd/server

mod:
	GO111MODULE=on go mod download

grpc/cli:
	evans -r -p8080 repl

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.27.0 golangci-lint run -v

.PHONY: build mod lint