BUILD_DIR := bin
CGO_CFLAGS := -Wno-nullability-completeness -Wno-expansion-to-defined
export CGO_CFLAGS
all: cat echo head mkdir ls

$(BUILD_DIR):
				@echo "Folder $(BUILD_DIR) does not exist, creating it..."
				mkdir -p $@

cat: |$(BUILD_DIR)
				@echo "Building cat..."
				@go build -o bin/cat ./cmd/cat/cat.go

echo: |$(BUILD_DIR)
				@echo "Building echo..."
				@go build -o bin/echo ./cmd/echo/echo.go

head: |$(BUILD_DIR)
				@echo "Building head..."
				@go build -o bin/head ./cmd/head/head.go

mkdir: |$(BUILD_DIR)
				@echo "Building mkdir..."
				@go build -o bin/mkdir ./cmd/mkdir/mkdir.go

ls: |$(BUILD_DIR)
				@echo "Building ls..."
				@go build -o bin/ls ./cmd/ls/ls.go

clean:
				@echo "Cleaning up..."
				@rm -rf bin/*