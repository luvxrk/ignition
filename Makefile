# Define variables
BINARY_NAME=ignition
BUILD_DIR=build
GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)
GOBIN=$(shell go env GOPATH)/bin

# Default target
.PHONY: all
all: build

# Build target: detects OS and Architecture
build:
	@echo "Building for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Binary generated for $(GOOS)/$(GOARCH)"

# Install target: builds and installs the binary at $GOBIN or $GOPATH/bin path
install:
	@echo "Installing the binary..."
	@go install
	@echo "Installed the binary at $(GOBIN)"

# Clean target: remove build artifacts
clean:
	@go clean
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned up build artifacts"
