# Builds the hugepages-demo
build: 
	go fmt ./...
	go vet ./...
	GOOS=linux GOARCH=ppc64le go build -o bin/hugepagesd -ldflags "${GO_LDFLAGS}" cmd/main.go
.PHONY: build

build-image: 
	podman build -t 
.PHONY: build-image

push-image: 
	podman