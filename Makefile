.PHONY: run build clean wire

APP_NAME = apiserver
GO ?= GO111MODULE=on go
BUILD_DIR = $(PWD)/build
MAIN_FILE = ./application/cmd/main.go
SERVER_FILE = ./application/server.go
MIGRATION_DIR = $(PWD)/infrastructure/migrations

setup:
	go mod tidy
	go install github.com/google/wire/cmd/wire@latest

# remove binary		
clean:
	echo "remove bin exe"
	rm -rf $(BUILD_DIR)

# build binary
build:
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

swag:
	swag init -g $(SERVER_FILE)

wire:
	cd application && wire

# local run
local:
	make swag
	make wire
	make build
	$(BUILD_DIR)/$(APP_NAME)

docker.fiber.build:
	make swag
	make wire
	docker build -t fiber .

docker.fiber.local:
	make docker.fiber.build
	docker run --rm -p 8080:8080 --name $(APP_NAME) --env-file ./.env.local fiber


docker.fiber:
	make docker.fiber.build
	docker run --rm -p 8080:8080 --name $(APP_NAME) --env PHASE=prod fiber

