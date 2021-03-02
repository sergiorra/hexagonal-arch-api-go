BUILD_DIR = ./build

.PHONY: golangci-lint
## golangci-lint: Runs golangci-lint
golangci-lint:
	golangci-lint run


.PHONY: build
## build: Builds the binary
build: clean golangci-lint
	go build -o $(BUILD_DIR)/api ./cmd/api


.PHONY: clean
## clean: Cleans build folder
clean:
	go clean
	go clean -testcache
	- rm -rf $(BUILD_DIR)


.PHONY: help
## help: Prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
