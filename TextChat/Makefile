MAIN_FILE_SERVER = main_server.go
MAIN_FILE_CLIENT = main_client.go

BUILD_DIR = build

BIN_NAME_SERVER = server
BIN_NAME_CLIENT = client

BUILD_PATH_SERVER = $(BUILD_DIR)\\$(BIN_NAME_SERVER).exe
BUILD_PATH_CLIENT = $(BUILD_DIR)\\$(BIN_NAME_CLIENT).exe

build:
	@echo "Compiling $(BIN_NAME_SERVER)..."
	@go build -o $(BUILD_PATH_SERVER) $(MAIN_FILE_SERVER)
	@echo "Server compiled and build in $(BUILD_PATH_SERVER)"

	@echo "Compiling $(BIN_NAME_CLIENT)"
	@go build -o $(BUILD_PATH_CLIENT) $(MAIN_FILE_CLIENT)
	@echo "Client compiled and build in $(BUILD_PATH_CLIENT)"

clean:
	@echo "Cleaning binary files..."
	@rm -rf .\\$(BUILD_DIR)\\*.exe
	@echo "Finished cleaning"

.PHONY: build clean