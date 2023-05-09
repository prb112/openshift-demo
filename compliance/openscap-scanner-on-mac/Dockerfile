FROM fedora:37

RUN dnf -y upgrade --refresh && \
        dnf install -y openscap-scanner && \
        dnf clean all && \
        mkdir -p /content