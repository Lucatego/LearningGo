SRC_DIR = src
BUILD_DIR = build
BINARY_NAME_C = client
BINARY_NAME_S = server

BINARY_PATH_C = $(BUILD_DIR)/$(BINARY_NAME_C).exe
BINARY_PATH_S = $(BUILD_DIR)/$(BINARY_NAME_S).exe

build:
	@echo "Compilando $(BINARY_NAME_S)..."
	@go build -o $(BINARY_PATH_S) $(SRC_DIR)/$(BINARY_NAME_S)/server.go
	@echo "Compilando $(BINARY_NAME_C)..."
	@go build -o $(BINARY_PATH_C) $(SRC_DIR)/$(BINARY_NAME_C)/client.go
	@echo "Compilacion exitosa"

clean:
	@echo "Limpiando archivos compilados..."
	@cd $(BUILD_DIR)
	@rm *.exe
	@cd ..
	@echo "Limpieza completa"
	
.PHONY: clean build
