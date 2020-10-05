FROM node:lts-stretch as node


# node plugins
COPY package.json .
COPY yarn.lock .
RUN yarn install

FROM node:lts-stretch as installer

ARG PROTOC_VERSION

# protoc
RUN mkdir /installer
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -O protoc.zip
RUN mkdir -p /installer/protoc
RUN unzip -o protoc.zip -d /installer/protoc/

# Download Java Lib
FROM gradle:latest as java
COPY build.gradle .
RUN gradle download
RUN mv installer /

# go
FROM golang:1.15 as gopher
ARG GO_GRPC_TOOLS_VERSION
# install the protocol compiler plugin for Go
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go
# installing grpc-go cmd, see https://github.com/grpc/grpc.io/issues/298
RUN git clone -b v$GO_GRPC_TOOLS_VERSION https://github.com/grpc/grpc-go && cd grpc-go/cmd/protoc-gen-go-grpc && go install .

FROM node:lts-stretch-slim
USER root
COPY --from=node /node_modules /node_modules
COPY --from=installer /installer/protoc /usr/bin/protoc
COPY --from=java /installer/protoc-gen-grpc-java /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-go /usr/bin/
COPY --from=gopher /go/bin/protoc-gen-go-grpc /usr/bin/

RUN chmod +x /usr/bin/protoc-gen-grpc-java

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
