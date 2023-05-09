FROM quay.io/centos/centos:stream9-minimal as builder

#################################################################################
# DNF Package Install List
ARG DNF_LIST="\
  jq \
  tar \
  gcc \
  make \
  git \
  gpgme-devel \
  libassuan-devel \
"

RUN set -ex \
     && dnf install -y --nodocs --setopt=install_weak_deps=false ${DNF_LIST}    \
     && dnf clean all -y                                                        \
     && GO_VERSION=go1.19.5                                       \
     && curl -sL https://golang.org/dl/${GO_VERSION}.linux-amd64.tar.gz         \
        | tar xzvf - --directory /usr/local/                                    \
     && /usr/local/go/bin/go version                                            \
     && ln -f /usr/local/go/bin/go /usr/bin/go                                  \
     && make

ENV PATH="/root/platform/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

FROM quay.io/centos/centos:stream9-minimal

COPY --from=builder /build/_output/cli /

LABEL \
  name="numa-machine-reporter"                                                             \
  license=ApacheV2                                                                         \
  distribution-scope="public"                                                              \
  io.openshift.tags="numa-machine-reporter"                                                \
  summary="numa-machine-reporter image"                                                    \
  io.k8s.display-name="numa-machine-reporter"                                              \
  build_date="`date +'%Y%m%d%H%M%S'`"                                                      \
  project="https://github.com/openshift/oc-mirror"                                         \
  description="numa-machine-reporter reports alignment of processes and containers"        \
  io.k8s.description="numa-machine-reporter reports alignment of processes and containers."


WORKDIR /build
ENTRYPOINT ["make"]
CMD []


COPY _output /usr/local/bin
ENTRYPOINT ["/bin/sh"]