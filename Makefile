ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: test coverage format dev clean build help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
BINARY_NAME=isbn
COVERAGE_FILE=coverage.out

# Default target
all: format test coverage build

# Run all tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...
	@if [ -f ./test.sh ]; then \
		echo "Running test script..."; \
		./test.sh; \
	fi

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report generated: coverage.html"
	$(GOCMD) tool cover -func=$(COVERAGE_FILE)

# Format code
format:
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	$(GOCMD) mod tidy

# Run in development mode
dev:
	@echo "Starting development server..."
	$(GOCMD) run main.go

# Build the application
build:
	@echo "Building application..."
	$(GOBUILD) -o $(BINARY_NAME) -v ./

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(COVERAGE_FILE)
	rm -f coverage.html