# Targets not related to individual files
.PHONY: all build test cover_func cover_html out clean vet loc run fmt

# Build constants
BUILD_OUT_DIR = bin
BINARY_FILE_NAME = go-ray-tracer
MAIN_PROGRAM_FILE = main.go
TEST_COVERAGE_PROFILE = coverage.out

all: out fmt vet test loc build

out:
	mkdir -p $(BUILD_OUT_DIR)

fmt:
	go fmt ./...

vet:
	go vet ./...

test: out
	go test -race -v ./... -coverprofile=$(BUILD_OUT_DIR)/$(TEST_COVERAGE_PROFILE)

loc:
	find . -type f -not -path "./vendor/*" -name "*.go" | xargs wc -l

build: out
	go build -o $(BUILD_OUT_DIR)/$(BINARY_FILE_NAME) $(MAIN_PROGRAM_FILE)

run:
	go run $(MAIN_PROGRAM_FILE)

cover_func: test
	go tool cover -func=$(BUILD_OUT_DIR)/$(TEST_COVERAGE_PROFILE)

cover_html: test
	go tool cover -html=$(BUILD_OUT_DIR)/$(TEST_COVERAGE_PROFILE)

clean:
	rm -rf $(BUILD_OUT_DIR)
