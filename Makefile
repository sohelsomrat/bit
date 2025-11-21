.PHONY: help build build-all clean test snapshot install

# Directories
BIN_DIR := bin
CMD_DIR := ./cmd/bit

help:
	@echo "Bit - Terminal ANSI Logo Designer"
	@echo ""
	@echo "Available targets:"
	@echo "  make build          - Build bit binary"
	@echo "  make build-all      - Build for all platforms"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make test           - Run tests"
	@echo "  make snapshot       - Test GoReleaser build (no publish)"
	@echo "  make install        - Install binary to /usr/local/bin"

# Build single binary
build:
	@echo "Building bit..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/bit $(CMD_DIR)
	@echo "Done! Binary in ./$(BIN_DIR)/"

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(BIN_DIR)
	GOOS=linux   GOARCH=amd64 go build -o $(BIN_DIR)/bit-linux-amd64 $(CMD_DIR)
	GOOS=linux   GOARCH=arm64 go build -o $(BIN_DIR)/bit-linux-arm64 $(CMD_DIR)
	GOOS=darwin  GOARCH=amd64 go build -o $(BIN_DIR)/bit-darwin-amd64 $(CMD_DIR)
	GOOS=darwin  GOARCH=arm64 go build -o $(BIN_DIR)/bit-darwin-arm64 $(CMD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/bit-windows-amd64.exe $(CMD_DIR)
	GOOS=windows GOARCH=arm64 go build -o $(BIN_DIR)/bit-windows-arm64.exe $(CMD_DIR)
	@echo "Done! Binaries in ./$(BIN_DIR)/"

# Clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BIN_DIR)/
	rm -rf dist/
	@echo "Done!"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Test GoReleaser build locally
snapshot:
	@echo "Running GoReleaser snapshot build..."
	@if ! command -v goreleaser >/dev/null 2>&1; then \
		echo "Error: goreleaser not found. Install with: brew install goreleaser"; \
		exit 1; \
	fi
	goreleaser release --snapshot --clean
	@echo ""
	@echo "Snapshot build complete! Artifacts in ./dist/"
	@echo ""
	@echo "Test the binaries:"
	@echo "  ls dist/bit_*/bit"
	@echo "  ./dist/bit_*/bit -version"
	@echo "  ./dist/bit_*/bit -list"
	@echo "  ./dist/bit_*/bit \"Hello\""

# Install binary globally
install: build
	@echo "Installing binary to /usr/local/bin..."
	@sudo cp $(BIN_DIR)/bit /usr/local/bin/bit
	@echo "Done! Run 'bit' from anywhere."
