.PHONY: run build clean

APP_NAME = apiserver
GO ?= GO111MODULE=on go
BUILD_DIR = $(PWD)/build
MAIN_FILE = ./application/cmd/main.go
MIGRATION_DIR = $(PWD)/infrastructure/migrations


# remove binary		
clean:
	echo "remove bin exe"
	rm -rf $(BUILD_DIR)

# build binary
build:
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

swag:
	swag init -g $(MAIN_FILE)


# local run
local:
	make swag
	make build
	$(BUILD_DIR)/$(APP_NAME)


docker.fiber.build:
	docker build -t fiber .

docker.fiber:
	make docker.fiber.build
	docker run --rm -p 8080:8080 --name $(APP_NAME) --env fiber

