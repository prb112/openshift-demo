export GO111MODULE=on

OUTDIR   := $(shell pwd)/bin

# Runs ALL
.PHONY: all
all: clean tidy test-unit build

# Build with golang
.PHONY: build
build: fmt vet build-windows build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-ppc64le

# Format and Vet
.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/...

# Build for Architecture
.PHONEY: build-windows
build-windows:
	cd cmd/oc-test-multiarch-plugin && \
	GOOS=windows GOARCH=amd64 GO_BUILD_PACKAGES=./cmd/oc-test-multiarch-plugin go build -o ../../bin/oc-multiarch-amd64.exe 

.PHONEY: build-darwin-amd64
build-darwin-amd64:
	cd cmd/oc-test-multiarch-plugin && \
	GOOS=darwin GOARCH=amd64 GO_BUILD_PACKAGES=./cmd/oc-test-multiarch-plugin go build -o ../../bin/oc-multiarch-darwin-amd64

.PHONEY: build-darwin-arm64
build-darwin-arm64:
	cd cmd/oc-test-multiarch-plugin && \
	GOOS=darwin GOARCH=arm64 GO_BUILD_PACKAGES=./cmd/oc-test-multiarch-plugin go build -o ../../bin/oc-multiarch-darwin-arm64

.PHONEY: build-linux-amd64
build-linux-amd64:
	cd cmd/oc-test-multiarch-plugin && \
	GOOS=linux GOARCH=amd64 GO_BUILD_PACKAGES=./cmd/oc-test-multiarch-plugin go build -o ../../bin/oc-multiarch-linux-amd64

.PHONEY: build-linux-ppc64le
build-linux-ppc64le:
	cd cmd/oc-test-multiarch-plugin && \
	GOOS=linux GOARCH=ppc64le GO_BUILD_PACKAGES=./cmd/oc-test-multiarch-plugin go build -o ../../bin/oc-multiarch-linux-ppc64le

# Tests the code
.PHONY: test
test:
	go test -timeout 10m ./pkg/... ./cmd/... -coverprofile cover.out

# Downloads the Dependencies
.PHONY: deps
deps:
	go get k8s.io/client-go/plugin/pkg/client/auth
