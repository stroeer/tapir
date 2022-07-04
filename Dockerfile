# Protoc binary
FROM node:lts-alpine as protoc-bin
ARG PROTOC_VERSION
RUN mkdir /installer
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -O protoc.zip
RUN mkdir -p /installer/protoc
RUN unzip -o protoc.zip -d /installer/protoc/

# protoc-gen-js
# this is a workaround for https://github.com/protocolbuffers/protobuf-javascript/issues/127 thanks to https://github.com/namely/docker-protoc/pull/337
FROM golang:1-bullseye AS protoc-gen-js
ARG PROTOC_VERSION
ARG PROTOBUF_JS_VERSION=v3.$PROTOC_VERSION

RUN set -ex && apt-get update && apt-get install -y --no-install-recommends git

WORKDIR /tmp
RUN git clone --depth 1 --shallow-submodules -b $PROTOBUF_JS_VERSION --recursive https://github.com/protocolbuffers/protobuf-javascript && \
    git clone --depth 1 --shallow-submodules --recursive https://github.com/grpc/grpc


WORKDIR /tmp/protobuf-javascript
RUN /tmp/grpc/tools/bazel build //generator:protoc-gen-js


# Node
FROM node:lts-alpine as node

COPY node .
RUN npm ci


# Java
FROM gradle:7 as java

COPY java .

RUN gradle download
RUN mv installer /


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
COPY --from=gopher /go/bin/protoc-gen-go /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-go-grpc /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-grpc-gateway /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-openapiv2 /usr/bin/
COPY --from=java /installer/protoc-gen-grpc-java /usr/bin/
COPY --from=node /node_modules /node_modules
COPY --from=protoc-bin /installer/protoc /usr/bin/protoc
COPY --from=protoc-gen-js /tmp/protobuf-javascript/bazel-bin/generator/protoc-gen-js /usr/bin/

RUN chmod +x /usr/bin/protoc-gen-grpc-java

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
