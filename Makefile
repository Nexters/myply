.PHONY: run build clean

GO ?= GO111MODULE=on go
APP_NAME = myply
BIN_DIR = ./bin
BUILD_DIR = ./application/cmd
BUILD_FILE = $(addprefix $(BUILD_DIR)/, main.go)

# local run
local:
	$(GO) run $(BUILD_FILE)

# build binary
build:
	$(GO) build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) $(BUILD_FILE)

# remove binary		
clean:
	echo "remove bin exe"
	rm -f $(addprefix $(BIN_DIR)/, $(APP_NAME))
