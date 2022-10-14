# protoc and plugins
FROM golang:1-bullseye AS build
ARG PROTOC_VERSION
ARG PROTOC_VERSION
ARG GRPC_VERSION
ARG GRPC_JAVA_VERSION
ARG PROTOBUF_JS_VERSION

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

WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -o protoc.zip && \
    unzip protoc.zip && \
    chmod +x bin/protoc

WORKDIR /tmp
RUN git clone --depth 1 --shallow-submodules -b v$GRPC_VERSION --recursive https://github.com/grpc/grpc && \
    git clone --depth 1 --shallow-submodules -b v$GRPC_JAVA_VERSION --recursive https://github.com/grpc/grpc-java && \
    git clone --depth 1 --shallow-submodules -b v$PROTOBUF_JS_VERSION --recursive https://github.com/protocolbuffers/protobuf-javascript

ARG bazel=/tmp/grpc/tools/bazel

WORKDIR /tmp/grpc
RUN $bazel build //external:protocol_compiler && \
    $bazel build //src/compiler:all && \
    $bazel build //test/cpp/util:grpc_cli

WORKDIR /tmp/grpc-java
RUN $bazel build //compiler:grpc_java_plugin

# this is a workaround for https://github.com/protocolbuffers/protobuf-javascript/issues/127
WORKDIR /tmp/protobuf-javascript
RUN $bazel build //generator:protoc-gen-js

# Node
FROM node:lts-alpine as node

COPY node .
RUN npm ci


# Go
FROM golang:1 as gopher
ARG GO_GRPC_TOOLS_VERSION

COPY go /temp
WORKDIR /temp

RUN cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %

# Composition
FROM debian:bullseye-slim

RUN set -ex && apt-get update && apt-get install -y --no-install-recommends nodejs && apt-get clean

USER root

# protoc and plugins
COPY --from=build /tmp/protoc /usr/bin/protoc
COPY --from=build /tmp/grpc/bazel-bin/src/compiler/ /usr/bin/
COPY --from=build /tmp/protobuf-javascript/bazel-bin/generator/protoc-gen-js /usr/bin/
COPY --from=build /tmp/grpc-java/bazel-bin/compiler/ /usr/bin/

# go tools
COPY --from=gopher /go/bin/protoc-gen-go /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-go-grpc /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-grpc-gateway /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-openapiv2 /usr/bin/

# node
COPY --from=node /node_modules /node_modules

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
