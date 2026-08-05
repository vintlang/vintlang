.PHONY: build build-all build-linux build-darwin build-windows build-android deps clean test checksums release

# ── Variables ──────────────────────────────────────────────────────────────────
APP_NAME    := vint
MODULE      := github.com/vintlang/vintlang
VERSION     ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT      := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE  := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS     = -s -w \
	-X '$(MODULE)/config.VINT_VERSION=$(VERSION)' \
	-X '$(MODULE)/config.Commit=$(COMMIT)' \
	-X '$(MODULE)/config.BuildDate=$(BUILD_DATE)'
BUILD_DIR   := ./build
CGO_ENABLED ?= 0

# ── Development ───────────────────────────────────────────────────────────────
build: deps
	go run counter.go > internal/toolkit/count.txt
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(APP_NAME) .

# ── Cross-platform builds (all) ──────────────────────────────────────────────
build-all: deps
	go run counter.go > internal/toolkit/count.txt
	@echo "==> Building for all platforms ($(VERSION))..."
	mkdir -p $(BUILD_DIR)
	# macOS Apple Silicon (M1/M2/M3/M4)
	GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64      .
	# macOS Intel
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64      .
	# Linux amd64
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64       .
	# Linux arm64
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64       .
	# Windows amd64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .
	# Windows arm64
	GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-arm64.exe .
	# Android arm64
	GOOS=android GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-android-arm64     .
	@echo "==> Packaging archives..."
	cd $(BUILD_DIR) && cp $(APP_NAME)-darwin-arm64 $(APP_NAME) && tar czf $(APP_NAME)-darwin-arm64.tar.gz  $(APP_NAME) && rm $(APP_NAME)
	cd $(BUILD_DIR) && cp $(APP_NAME)-darwin-amd64 $(APP_NAME) && tar czf $(APP_NAME)-darwin-amd64.tar.gz  $(APP_NAME) && rm $(APP_NAME)
	cd $(BUILD_DIR) && cp $(APP_NAME)-linux-amd64  $(APP_NAME) && tar czf $(APP_NAME)-linux-amd64.tar.gz   $(APP_NAME) && rm $(APP_NAME)
	cd $(BUILD_DIR) && cp $(APP_NAME)-linux-arm64  $(APP_NAME) && tar czf $(APP_NAME)-linux-arm64.tar.gz   $(APP_NAME) && rm $(APP_NAME)
	cd $(BUILD_DIR) && cp $(APP_NAME)-windows-amd64.exe $(APP_NAME).exe && zip $(APP_NAME)-windows-amd64.zip $(APP_NAME).exe && rm $(APP_NAME).exe
	cd $(BUILD_DIR) && cp $(APP_NAME)-windows-arm64.exe $(APP_NAME).exe && zip $(APP_NAME)-windows-arm64.zip $(APP_NAME).exe && rm $(APP_NAME).exe
	cd $(BUILD_DIR) && cp $(APP_NAME)-android-arm64 $(APP_NAME) && tar czf $(APP_NAME)-android-arm64.tar.gz $(APP_NAME) && rm $(APP_NAME)
	@echo "==> Done! Binaries in $(BUILD_DIR)/"

# ── Individual platform builds ───────────────────────────────────────────────
build-darwin: deps
	@echo "==> Building for macOS..."
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 .
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 .
	cd $(BUILD_DIR) && cp $(APP_NAME)-darwin-arm64 $(APP_NAME) && tar czf $(APP_NAME)-darwin-arm64.tar.gz $(APP_NAME) && rm $(APP_NAME)
	cd $(BUILD_DIR) && cp $(APP_NAME)-darwin-amd64 $(APP_NAME) && tar czf $(APP_NAME)-darwin-amd64.tar.gz $(APP_NAME) && rm $(APP_NAME)

build-linux: deps
	@echo "==> Building for Linux..."
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 .
	cd $(BUILD_DIR) && cp $(APP_NAME)-linux-amd64 $(APP_NAME) && tar czf $(APP_NAME)-linux-amd64.tar.gz $(APP_NAME) && rm $(APP_NAME)
	cd $(BUILD_DIR) && cp $(APP_NAME)-linux-arm64 $(APP_NAME) && tar czf $(APP_NAME)-linux-arm64.tar.gz $(APP_NAME) && rm $(APP_NAME)

build-windows: deps
	@echo "==> Building for Windows..."
	mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-arm64.exe .
	cd $(BUILD_DIR) && cp $(APP_NAME)-windows-amd64.exe $(APP_NAME).exe && zip $(APP_NAME)-windows-amd64.zip $(APP_NAME).exe && rm $(APP_NAME).exe
	cd $(BUILD_DIR) && cp $(APP_NAME)-windows-arm64.exe $(APP_NAME).exe && zip $(APP_NAME)-windows-arm64.zip $(APP_NAME).exe && rm $(APP_NAME).exe

build-android: deps
	@echo "==> Building for Android..."
	mkdir -p $(BUILD_DIR)
	GOOS=android GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-android-arm64 .
	cd $(BUILD_DIR) && cp $(APP_NAME)-android-arm64 $(APP_NAME) && tar czf $(APP_NAME)-android-arm64.tar.gz $(APP_NAME) && rm $(APP_NAME)

# ── SHA256 checksums ─────────────────────────────────────────────────────────
checksums:
	@echo "==> Generating checksums..."
	cd $(BUILD_DIR) && shasum -a 256 *.tar.gz *.zip > checksums.txt
	@cat $(BUILD_DIR)/checksums.txt

# ── GitHub Release (requires gh CLI) ─────────────────────────────────────────
# The release version is computed by bumping the most recent vX.Y.Z tag,
# then the tag is created and pushed via git before the release is published.
LAST_TAG    := $(shell git tag --sort=-v:refname | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$$' | head -n1)
RELEASE_TAG := $(shell printf '%s' "$(LAST_TAG)" | awk -F. '{printf "v%d.%d.%d", substr($$1,2)+0, $$2, $$3+1}')

release: VERSION := $(RELEASE_TAG)
release: clean build-all checksums
	@echo "==> Creating tag $(RELEASE_TAG) and GitHub release..."
	git tag "$(RELEASE_TAG)"
	git push origin "$(RELEASE_TAG)"
	gh release create "$(RELEASE_TAG)" \
		--title "VintLang $(RELEASE_TAG)" \
		--generate-notes \
		$(BUILD_DIR)/*.tar.gz \
		$(BUILD_DIR)/*.zip \
		$(BUILD_DIR)/checksums.txt

# ── Utilities ─────────────────────────────────────────────────────────────────
deps:
	@echo "==> Checking dependencies..."
	go mod tidy

test:
	@echo -e '\nTesting Lexer...'
	@./gotest --format testname ./lexer/
	@echo -e '\nTesting Parser...'
	@./gotest --format testname ./parser/
	@echo -e '\nTesting AST...'
	@./gotest --format testname ./ast/
	@echo -e '\nTesting Object...'
	@./gotest --format testname ./object/
	@echo -e '\nTesting Evaluator...'
	@./gotest --format testname ./evaluator/

clean:
	rm -f $(APP_NAME)
	rm -rf $(BUILD_DIR)
