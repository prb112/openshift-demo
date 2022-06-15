# Builds the hugepages-demo
# GOOS=linux GOARCH=ppc64le 

build: 
	go fmt ./...
	go vet ./...
	go build -o bin/hugepagesd -ldflags "-v ${GO_LDFLAGS}" cmd/main.go
.PHONY: build

build-image: 
	podman build -t 
.PHONY: build-image

push-image: 
	podman