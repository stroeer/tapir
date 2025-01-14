# @stroeer/tapir-buf-v1

This package includes all tapir models + grpc services for node.js + typescript definitions in esm and cjs. It's based on the package `@bufbuild/protobuf` from Buf generated with `bufbuild/es`.

:information_source: If esm only is fine for you, use the SDK from https://buf.build/stroeer/tapir.

## Why this package?

The generated packages of the "Buf Schema Registry" (BSR) includes only **esm** syntax. While this is totally fine, we need **cjs** syntax for our node.js projects. This package provides both syntaxes. 

:warning: We will deprecate this package when our node.js projects are fully migrated to esm.

## Usage

```bash
npm i @stroeer/tapir-buf-v1
# or
yarn add @stroeer/tapir-buf-v1
```

## Development

Make sure you generated the gRPC models and services for node.

```shell
# from tapir root folder
make generate TEMPLATE=buf.gen.node-buf.yaml
```

### node

```shell
# https://github.com/nvm-sh/nvm#installing-and-updating
brew install nvm

# add this to your shell/bash profile:
# source /usr/local/opt/nvm/nvm.sh

nvm install 22
nvm use 22
# optional
nvm alias system 22
```

### install

```shell
npm i
```

## Helpful resources

- tapir [Buf Schema Registry (BSR)](https://buf.build/stroeer/tapir)
- packages: protobuf, protoc-gen-es [@bufbuild/protobuf-es](https://github.com/bufbuild/protobuf-es)
- Manual/Docs [@bufbuild/protobuf-es](https://github.com/bufbuild/protobuf-es/blob/main/MANUAL.md)
- packages: connect, connect-web, connect-fastify [@connectrpc/connect-es](https://github.com/connectrpc/connect-es)
- connectrpc-es examples [@connectrpc/examples-es](https://github.com/connectrpc/examples-es)
- [connectrpc docs](https://connectrpc.com/docs/introduction)
