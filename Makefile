GIT_TAG_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null)
GIT_COMMIT_HASH := $(shell git rev-parse --short HEAD)
VERSION         := $(if $(GIT_TAG_VERSION),$(GIT_TAG_VERSION)-$(GIT_COMMIT_HASH),$(GIT_COMMIT_HASH))

TARGETOS := linux
TARGETARCH := amd64

MODULE_NAME     := tgbot
APP_VERSION_VAR_PATH := ${MODULE_NAME}/cmd.appVersion

BINARY_NAME := tgbot
MAIN_GO_FILE := ./main.go
REGISTRY := urapikh

LDFLAGS         := -s -w -X '${APP_VERSION_VAR_PATH}=${VERSION}'

# Додаємо build-only сюди
.PHONY: format lint test get build build-only image push clean

format:
	@echo ">>> Formatting Go files..."
	@gofmt -s -w ./

lint:
	@echo ">>> Linting Go files..."
	@golangci-lint run

test:
	@echo ">>> Running Go tests..."
	@go test -v ./... # Краще вказати ./... для всіх тестів

# Ціль для локального оновлення/завантаження залежностей
get:
	@echo ">>> Getting/updating Go dependencies..."
	@go get ./... # Більш явно, що ви хочете оновити/завантажити для поточного модуля

# Ціль тільки для компіляції (для Docker)
build-only:
	@echo ">>> Building binary ${BINARY_NAME} for ${TARGETOS}/${TARGETARCH}..."
	@CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="${LDFLAGS}" -o ${BINARY_NAME} ${MAIN_GO_FILE}

# Ціль для локальної збірки (включає все)
build: format get build-only
	@echo ">>> Build complete: ${BINARY_NAME}"

image:
	@echo ">>> Building Docker image..."
	@docker build --platform linux/${TARGETARCH} . -t ${REGISTRY}/${MODULE_NAME}:${VERSION}-${TARGETARCH}

push:
	@echo ">>> Pushing Docker image ${REGISTRY}/${MODULE_NAME}:${VERSION}-${TARGETARCH}..."
	@docker push ${REGISTRY}/${MODULE_NAME}:${VERSION}-${TARGETARCH}

clean:
	@echo ">>> Cleaning up..."
	@rm -f ${BINARY_NAME}
	@go clean