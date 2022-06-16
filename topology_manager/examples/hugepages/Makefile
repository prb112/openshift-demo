# Builds the hugepages-demo and is really only supported on linux/ppc64le
# GOOS=linux GOARCH=ppc64le 

build: 
	go fmt ./...
	go vet ./...
	go build -o bin/hugepagesd -ldflags "-v ${GO_LDFLAGS}" cmd/main.go
.PHONY: build

build-image: 
	podman build -t quay.io/pbastide_rh/openshift-demo:hugepages-demo-$$(date +%Y%m%d%H%M%S) .
.PHONY: build-image
