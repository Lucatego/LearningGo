SRC_DIR = src
BUILD_DIR = build
BINARY_NAME_C = client
BINARY_NAME_S = server

BINARY_PATH_C = $(BUILD_DIR)/$(BINARY_NAME_C).exe
BINARY_PATH_S = $(BUILD_DIR)\\$(BINARY_NAME_S).exe

build:
	@echo "Compilando $(BINARY_NAME_S)..."
	@go build -o $(BINARY_PATH_S) $(SRC_DIR)/$(BINARY_NAME_S)/server.go
	@echo "Compilando $(BINARY_NAME_C)..."
	@go build -o $(BINARY_PATH_C) $(SRC_DIR)/$(BINARY_NAME_C)/client.go
	@echo "Compilacion exitosa"

clean:
	@echo "Limpiando archivos compilados..."
	@rm -rf $(BUILD_DIR)\*
	@echo "Limpieza completa"
	
run_server:
	@echo "Iniciando servidor..."
	@echo "Ruta: $(BINARY_PATH_S)"
	@start "" "$(BINARY_PATH_S)"
run_client:
	@echo "Iniciando cliente..."
	@echo "Ruta: $(BINARY_PATH_C)"
	@start "" "$(BINARY_PATH_C)"
	
.PHONY: clean build run_server run_client
