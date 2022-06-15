# Builder
FROM golang:1.18 AS builder
WORKDIR /
COPY . .
RUN make build

# We need a package manager so we can run the command
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

RUN dnf install libhugetlbfs-utils -y

COPY --from=builder /bin/hugepages-demo/cmd /

ENTRYPOINT ["hugectl /cmd"]