
# Protoc binary
FROM node:lts-alpine as protoc-bin
ARG PROTOC_VERSION
RUN mkdir /installer
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -O protoc.zip
RUN mkdir -p /installer/protoc
RUN unzip -o protoc.zip -d /installer/protoc/


# Node
FROM node:lts-alpine as node

COPY node .
RUN npm ci


# Java
FROM gradle:latest as java

COPY java .

RUN gradle download
RUN mv installer /


# Go
FROM golang:1.15 as gopher
ARG GO_GRPC_TOOLS_VERSION

COPY go /temp
WORKDIR /temp

RUN cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %

# Composition
FROM node:lts-stretch-slim

USER root
COPY --from=gopher /go/bin/protoc-gen-go /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-go-grpc /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-grpc-gateway /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-openapiv2 /usr/bin/
COPY --from=java /installer/protoc-gen-grpc-java /usr/bin/
COPY --from=node /node_modules /node_modules
COPY --from=protoc-bin /installer/protoc /usr/bin/protoc

RUN chmod +x /usr/bin/protoc-gen-grpc-java

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
