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
"

RUN set -ex \
    && dnf install -y --nodocs ${PACKAGE_LIST}  \
    && dnf clean all -y \
    && python3 -m pip install --upgrade --user ansible 

RUN pip3 install --upgrade pip \
    && pip3 install --upgrade virtualenv \
    && pip3 install requests \
    && python3 -m pip install ansible

RUN find / -iname '*ansible*' && mkdir /ansible && \
    mkdir -p /etc/ansible

WORKDIR /ansible

ENTRYPOINT ["/root/.local/bin/ansible-playbook"]
CMD ["--version" ]

#COPY --chown=1001:0 /opt/ol/wlp/usr /opt/ol/wlp/usr

# https://learn.microsoft.com/en-us/azure/developer/ansible/configure-in-docker-container?tabs=azure-cli