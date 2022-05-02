BUILD_DIR := bin
CGO_CFLAGS := -Wno-nullability-completeness -Wno-expansion-to-defined
export CGO_CFLAGS
all: textutils shellutils fileutils

textutils: cat head

shellutils: echo false true yes whoami

fileutils: cp ls mkdir mv sync

$(BUILD_DIR):
				@echo "Folder $(BUILD_DIR) does not exist, creating it..."
				mkdir -p $@

# Text utilities

cat: |$(BUILD_DIR)
				@echo "Building cat..."
				@go build -o bin/cat ./cmd/text/cat/cat.go

head: |$(BUILD_DIR)
				@echo "Building head..."
				@go build -o bin/cat ./cmd/text/head/head.go

# Shell utilities

echo: |$(BUILD_DIR)
				@echo "Building echo..."
				@go build -o bin/echo ./cmd/shell/echo/echo.go

false: |$(BUILD_DIR)
				@echo "Building false..."
				@go build -o bin/false ./cmd/shell/false/false.go

true: |$(BUILD_DIR)
				@echo "Building true..."
				@go build -o bin/true ./cmd/shell/true/true.go

yes: |$(BUILD_DIR)
				@echo "Building yes..."
				@go build -o bin/yes ./cmd/shell/yes/yes.go

whoami: |$(BUILD_DIR)
				@echo "Building whoami..."
				@go build -o bin/whoami ./cmd/shell/whoami/whoami.go

# File utilities

cp: |$(BUILD_DIR)
				@echo "Building cp..."
				@go build -o bin/cp ./cmd/file/cp/cp.go

mv: |$(BUILD_DIR)
				@echo "Building mv..."
				@go build -o bin/mv ./cmd/file/mv/mv.go

mkdir: |$(BUILD_DIR)
				@echo "Building mkdir..."
				@go build -o bin/mkdir ./cmd/file/mkdir/mkdir.go

ls: |$(BUILD_DIR)
				@echo "Building ls..."
				@go build -o bin/ls ./cmd/file/ls/ls.go

sync: |$(BUILD_DIR)
				@echo "Building sync..."
				@go build -o bin/sync ./cmd/file/sync/sync.go

# Clean

clean:
				@echo "Cleaning up..."
				@rm -rf bin/*