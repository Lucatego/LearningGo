MAIN_FILE_SERVER = main\\main_server.go
TEST_FILE_SERVER = main\\main_client.go

BUILD_DIR = build

BIN_NAME_SERVER = server
BIN_NAME_TEST_SERVER = client

BUILD_PATH_SERVER = $(BUILD_DIR)\\$(BIN_NAME_SERVER).exe
BUILD_PATH_TEST_SERVER = $(BUILD_DIR)\\$(BIN_NAME_TEST_SERVER).exe

build_server:
	@echo "Compiling $(BIN_NAME_SERVER)..."
	@go build -o $(BUILD_PATH_SERVER) $(MAIN_FILE_SERVER)
	@echo "Server compiled and build in $(BUILD_PATH_SERVER)"

build_test_server:
	@echo "Compiling $(BIN_NAME_TEST_SERVER)"
	@go build -o $(BUILD_PATH_TEST_SERVER) $(TEST_FILE_SERVER)
	@echo "Client compiled and build in $(BUILD_PATH_TEST_SERVER)"

clean_build:
	@echo "Cleaning binary files..."
	@del $(BUILD_DIR) /S /F
	@echo "Finished cleaning"

test_server:
	@echo "Starting client..."
	@$(BUILD_PATH_TEST_SERVER)

run_server:
	@echo "Starting server..."
	@$(BUILD_PATH_SERVER)

.PHONY: build_server build_test_server clean_build run_server test_server