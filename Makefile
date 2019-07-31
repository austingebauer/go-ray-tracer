# Targets not related to individual files
.PHONY: all build test cover_func cover_html out clean vet loc

# Build constants
BUILD_OUT_DIR = bin
BINARY_FILE_NAME = go-ray-tracer
MAIN_PROGRAM_FILE = main.go
TEST_COVERAGE_PROFILE = coverage.out

all: out build test vet cover_func loc

build: out
	go build -o $(BUILD_OUT_DIR)/$(BINARY_FILE_NAME) $(MAIN_PROGRAM_FILE)

test: out
	go test ./... -coverprofile=$(BUILD_OUT_DIR)/$(TEST_COVERAGE_PROFILE)

cover_func: test
	go tool cover -func=$(BUILD_OUT_DIR)/$(TEST_COVERAGE_PROFILE)

cover_html: test
	go tool cover -html=$(BUILD_OUT_DIR)/$(TEST_COVERAGE_PROFILE)

out:
	mkdir -p $(BUILD_OUT_DIR)

loc:
	find . -type f -not -path "vendor" -name "*.go" | xargs wc -l

vet:
	go vet ./...

clean:
	rm -rf $(BUILD_OUT_DIR)