# protoc and plugins
FROM golang:1-bullseye AS build
ARG PROTOC_VERSION
ARG PROTOC_VERSION
ARG GRPC_VERSION
ARG GRPC_JAVA_VERSION
ARG PROTOBUF_JS_VERSION
ARG GO_PROTOC_GEN_GO_GRPC_VERSION
ARG GO_PROTOC_GEN_GO_VERSION
ARG GRPC_GATEWAY_VERSION

RUN set -ex && apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    pkg-config \
    cmake \
    curl \
    git \
    openjdk-11-jre \
    unzip \
    libtool \
    autoconf \
    zlib1g-dev \
    libssl-dev \
    clang \
    python3-pip

# protoc
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -o protoc.zip && \
    unzip protoc.zip && \
    chmod +x bin/protoc

WORKDIR /tmp
RUN git clone --depth 1 --shallow-submodules -b v$GRPC_VERSION --recursive https://github.com/grpc/grpc && \
    git clone --depth 1 --shallow-submodules -b v$GRPC_JAVA_VERSION --recursive https://github.com/grpc/grpc-java && \
    git clone --depth 1 --shallow-submodules -b v$PROTOBUF_JS_VERSION --recursive https://github.com/protocolbuffers/protobuf-javascript

ARG bazel=/tmp/grpc/tools/bazel

# grpc
WORKDIR /tmp/grpc
RUN $bazel build //external:protocol_compiler && \
    $bazel build //src/compiler:all && \
    $bazel build //test/cpp/util:grpc_cli

# grpc-java
WORKDIR /tmp/grpc-java
RUN $bazel build //compiler:grpc_java_plugin

# protoc-gen-js, this is a workaround for https://github.com/protocolbuffers/protobuf-javascript/issues/127
WORKDIR /tmp/protobuf-javascript
RUN $bazel build //generator:protoc-gen-js

# grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@${GRPC_GATEWAY_VERSION}
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${GRPC_GATEWAY_VERSION}

# go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${GO_PROTOC_GEN_GO_GRPC_VERSION}
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${GO_PROTOC_GEN_GO_VERSION}


# Node
FROM node:lts-alpine3.9 as node

COPY node .
RUN npm ci


# Composition
FROM debian:bullseye-slim

RUN set -ex && apt-get update && apt-get install -y --no-install-recommends nodejs && apt-get clean

USER root

# protoc and standard plugins
COPY --from=build /tmp/protoc /usr/bin/protoc
COPY --from=build /tmp/grpc/bazel-bin/src/compiler/ /usr/bin/
COPY --from=build /tmp/protobuf-javascript/bazel-bin/generator/protoc-gen-js /usr/bin/
COPY --from=build /tmp/grpc-java/bazel-bin/compiler/ /usr/bin/

# go and grcp-gateway
COPY --from=build /go/bin/* /usr/bin/

# node
COPY --from=node /node_modules /node_modules

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
