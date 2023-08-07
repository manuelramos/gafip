# Makefile

# Set the default target (development)
.DEFAULT_GOAL := dev

# Variables
APP_NAME := gafip
BUILD_DIR := build
ENV_FILE := .env

# Target to compile for development
dev:
	@echo "Building for development..."
	@make build BUILD_ENV=development

# Target to compile for production
prod:
	@echo "Building for production..."
	@make build BUILD_ENV=production

# Target to compile the binary
build:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Building for $(BUILD_ENV)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) -ldflags="-s -w" ./cmd/main.go

# Target to run the application
run:
	@echo "Running the application..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Target to setup environment variables
setup:
	@echo "Setting up environment variables from $(ENV_FILE)..."
	@export $(shell sed 's/=.*//' $(ENV_FILE))

# Target to clean the build directory
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

# Target to show help
help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  dev     Compile for development"
	@echo "  prod    Compile for production"
	@echo "  run     Run the application"
	@echo "  setup   Set up environment variables from .env"
	@echo "  clean   Clean the build directory"
	@echo "  help    Show this help"
