MAIN_FILE = main.go

SRC_SERVER = src/server
BUILD_DIR = build

BIN_NAME_SERVER = server
BIN_NAME_CLIENT = client

BUILD_PATH_SERVER = $(BUILD_DIR)\\$(BIN_NAME_SERVER).exe
BUILD_PATH_CLIENT = $(BUILD_DIR)\\$(BIN_NAME_CLIENT).exe

build:
	@echo "Compiling $(BIN_NAME_SERVER)..."
	@go build -o $(BUILD_PATH_SERVER) $(SRC_SERVER)\\$(MAIN_FILE)
	@echo "Server compiled and build in $(BUILD_PATH_SERVER)"

clean:
	@echo "Cleaning binary files..."
	@rm -rf $(BUILD_DIR)\\*
	@echo "Finished cleaning"

.PHONY: build clean