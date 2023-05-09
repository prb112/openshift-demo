# https://learn.microsoft.com/en-us/azure/developer/ansible/configure-in-docker-container?tabs=azure-cli
# https://github.com/ansible-community/ansible-bender
# https://ansible-community.github.io/ansible-bender/build/html/index.html
FROM quay.io/centos/centos:stream9 AS runner

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

# Package Install List
ARG PACKAGE_LIST="\
  gcc \
  libffi-devel \
  python3 \
  epel-release \
  python3-pip \
  wget \
  tar \
  make \
  git \
  python3-devel \
  nfs-utils \
  nfs4-acl-tools \
"

RUN set -ex \
    && dnf install -y --nodocs ${PACKAGE_LIST}  \
    && dnf clean all -y

RUN pip3 install --upgrade pip \
    && pip3 install requests \
    && python3 -m pip install ansible

WORKDIR /ansible

RUN groupadd -r ansible -g 1001 \
    && useradd -u 1001 -r -g 1001 -m -d /ansible -s /sbin/nologin ansible

RUN mkdir -p /.ansible/tmp \
    && chmod -R 755 /.ansible \
    && mkdir -p /ansible/.ansible \
    && mkdir -p /ansible/.ansible/tmp \
    && chmod -R 755 /ansible \
    && chown -R ansible:ansible /ansible \
    && chown -R ansible:ansible /.ansible

USER 1001

#ENTRYPOINT ["/usr/local/bin/ansible-playbook"]
#CMD ["--version" ]