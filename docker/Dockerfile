
# Protoc binary
FROM node:lts-stretch as protoc-bin
ARG PROTOC_VERSION
RUN mkdir /installer
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -O protoc.zip
RUN mkdir -p /installer/protoc
RUN unzip -o protoc.zip -d /installer/protoc/


# Node
FROM node:lts-stretch as node

COPY package.json .
COPY package-lock.json .
RUN npm install


# Java
FROM gradle:latest as java
COPY build.gradle .
RUN gradle download
RUN mv installer /


# Go
FROM golang:1.15 as gopher
ARG GO_GRPC_TOOLS_VERSION

RUN mkdir /temp
RUN cd /temp
WORKDIR /temp

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go install /go/pkg/mod/google.golang.org/protobuf@*/cmd/protoc-gen-go
RUN go install /go/pkg/mod/google.golang.org/grpc/cmd/protoc-gen-go-grpc@*


# Composition
FROM node:lts-stretch-slim

USER root
COPY --from=gopher /go/bin/protoc-gen-go /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-go-grpc /usr/bin/
COPY --from=java /installer/protoc-gen-grpc-java /usr/bin/
COPY --from=node /node_modules /node_modules
COPY --from=protoc-bin /installer/protoc /usr/bin/protoc

RUN chmod +x /usr/bin/protoc-gen-grpc-java

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
