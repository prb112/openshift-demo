# Builder
FROM docker.io/library/golang:1.18 AS builder
RUN mkdir -p src
WORKDIR /src
COPY . .
RUN make build

# We need a package manager so we can run the command
FROM registry.access.redhat.com/ubi8/ubi:latest

COPY --from=builder /src/bin/hugepagesd /

RUN yum update -y && \
    rm -rf /var/cache/yum && \
    rm -f /@System.solv && \
    yum clean all && \
    rm -rf /var/tmp/* && \
    rm -rf /tmp/*

ENTRYPOINT ["/hugepagesd"]