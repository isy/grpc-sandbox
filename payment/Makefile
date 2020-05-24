APP = user-service

build:
	go build -o build/${APP} ./cmd/server

mod:
	GO111MODULE=on go mod download

grpc/cli:
	evans -r -p8080 repl


.PHONY: build mod