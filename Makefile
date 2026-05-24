APP_NAME := shuffle
MAIN := ./cmd/shuffle
BIN_DIR := ./bin
BIN := $(BIN_DIR)/$(APP_NAME)

.PHONY: help build run test test-e2e test-github fmt vet tidy clean

help:
	@echo "Available commands:"
	@echo "  make build        Build binary"
	@echo "  make run          Run CLI help"
	@echo "  make test         Run unit and fast e2e tests"
	@echo "  make test-e2e     Run e2e tests without external systems"
	@echo "  make test-github  Run GitHub e2e tests"
	@echo "  make fmt          Format code"
	@echo "  make vet          Run go vet"
	@echo "  make tidy         Run go mod tidy"
	@echo "  make clean        Remove build artifacts"

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN) $(MAIN)

run:
	go run $(MAIN) --help

test:
	go test ./...

test-e2e:
	go test ./tests/e2e -v

test-github:
	@if [ -z "$$GITHUB_TOKEN" ]; then \
		read -s -p "GitHub token: " GITHUB_TOKEN; \
		echo ""; \
		E2E_GITHUB=1 GITHUB_TOKEN=$$GITHUB_TOKEN go test ./tests/e2e -v; \
	else \
		E2E_GITHUB=1 go test ./tests/e2e -v; \
	fi

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
