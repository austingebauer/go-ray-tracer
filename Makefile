# Targets not related to individual files
.PHONY: all build test cover out clean

# Build constants
BUILD_OUT_DIR = bin
BINARY_FILE_NAME = go-ray-tracer
MAIN_PROGRAM_FILE = main.go
TEST_COVERAGE_PROFILE = coverage.out

all: out build test cover_func

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

clean:
	rm -rf $(BUILD_OUT_DIR)