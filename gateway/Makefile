APP = gateway

build:
	go build -o build/${APP} ./cmd/server

mod:
	GO111MODULE=on go mod download

.PHONY: build mod