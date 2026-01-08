APP_NAME=portfolio
WEB_DIR=cmd/web
SERVER_DIR=cmd/server
STATIC_DIR=static

.PHONY: all build-web build-server run clean

all: build-web build-server

## Build WASM frontend
build-web:
	GOOS=js GOARCH=wasm go build -o $(STATIC_DIR)/main.wasm ./$(WEB_DIR)

## Build server binary
build-server:
	go build -o bin/$(APP_NAME)-server ./$(SERVER_DIR)

## Run server (serves WASM + HTML)
run: build-web
	go run ./$(SERVER_DIR)

## Clean artifacts
clean:
	rm -f $(STATIC_DIR)/main.wasm
	rm -f bin/$(APP_NAME)-server
