RUNTIME_ALIAS ?= podman
REPOOWNER ?= pbastide
IMAGENAME ?= machineinfo_reporter
IMAGETAG ?= latest

.PHONY: test-unit
test-unit:
	go test ./pkg/...

.PHONY: clean
clean:
	rm -rf _output

.PHONY: build
build: gofmt
	GO111MODULE=on GOPROXY=off GOFLAGS=-mod=vendor GOOS=linux GOARCH=ppc64le CGO_ENABLED=0 go build -v -o "${BASEDIR}/_output/machineinfo_reporter-v${TAG}-linux-ppc64le" "${path}"

.PHONY: image
image:
	@echo "building image"
	$(RUNTIME_ALIAS) build -f Dockerfile -t quay.io/$(REPOOWNER)/$(IMAGENAME):$(IMAGETAG) .

.PHONY: push
push: image
	@echo "pushing image"
	$(RUNTIME_ALIAS) push quay.io/$(REPOOWNER)/$(IMAGENAME):$(IMAGETAG)

.PHONY: gofmt
gofmt:
	@echo "Running gofmt"
	gofmt -s -w `find . -path ./vendor -prune -o -type f -name '*.go' -print`