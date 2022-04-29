BUILD_DIR := bin
CGO_CFLAGS := -Wno-nullability-completeness -Wno-expansion-to-defined
export CGO_CFLAGS
all: cat echo ls

$(BUILD_DIR):
				@echo "Folder $(BUILD_DIR) does not exist, creating it..."
				mkdir -p $@

cat: |$(BUILD_DIR)
				@echo "Building cat..."
				@go build -o bin/cat ./cmd/cat/cat.go

echo: |$(BUILD_DIR)
				@echo "Building echo..."
				@go build -o bin/echo ./cmd/echo/echo.go

ls: |$(BUILD_DIR)
				@echo "Building ls..."
				@go build -o bin/ls ./cmd/ls/ls.go

clean:
				@echo "Cleaning up..."
				@rm -rf bin/*