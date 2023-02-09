FROM golang:1.19.5-alpine3.17
ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:/opt/bento4/bin"

# FFMPEG
RUN apk add --update ffmpeg bash curl make
RUN apk add --update --upgrade curl python3 unzip bash gcc g++ scons ca-certificates wget libgcc make

# Install Bento
ENV BENTO4_VERSION=1.6.0-637
ENV BENTO4_INSTALL_DIR=/opt/bento4
ENV PATH=/opt/bento4/bin:${PATH}

# Fetch source.
RUN cd /tmp/ \
    && wget -O Bento4-${BENTO4_VERSION}.tar.gz https://github.com/axiomatic-systems/Bento4/archive/v${BENTO4_VERSION}.tar.gz \
    && tar -xzvf Bento4-${BENTO4_VERSION}.tar.gz && rm Bento4-${BENTO4_VERSION}.tar.gz

# Create installation directories.
RUN mkdir -p \
    ${BENTO4_INSTALL_DIR}/bin \
    ${BENTO4_INSTALL_DIR}/scripts \
    ${BENTO4_INSTALL_DIR}/include

# Build.
RUN cd /tmp/Bento4-${BENTO4_VERSION}/Build/Targets/x86-unknown-linux \
    && make AP4_BUILD_CONFIG=Release

# Install.
RUN cd /tmp \
    && cp -r Bento4-${BENTO4_VERSION}/Build/Targets/x86-unknown-linux/Release/. ${BENTO4_INSTALL_DIR}/bin \
    && cp -r Bento4-${BENTO4_VERSION}/Source/Python/utils/. ${BENTO4_INSTALL_DIR}/utils \
    && cp -r Bento4-${BENTO4_VERSION}/Source/Python/wrappers/. ${BENTO4_INSTALL_DIR}/bin \
    && cp -r Bento4-${BENTO4_VERSION}/Source/C++/**/*.h . ${BENTO4_INSTALL_DIR}/include

# Cleanup.
RUN rm -rf /var/cache/apk/* /tmp/*

ENV CGO_CFLAGS="-g -O2 -Wno-return-local-addr"

WORKDIR /go/src

#vamos mudar para o endpoint correto. Usando top apenas para segurar o processo rodando
ENTRYPOINT [ "top" ]