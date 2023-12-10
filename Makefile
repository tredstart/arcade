
# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

migrate: 
	@echo "Migrating..."
	@sqlite3 testing.db < migrations/create.sql

# Clean the binary
clean:
	@echo "Cleaning..."
	@sqlite3 prod.db < migrations/drop.sql
	@rm -f main

.PHONY: all build run test clean
		
