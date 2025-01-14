# @stroeer/tapir-proto-v1 (deprecated)

This package includes all tapir models without the grpc services.
It's based on the packages `@bufbuild/protobuf`.

## :warning: Deprecated

Use the package directly from the "Buf Schema Registry" (BSR): https://buf.build/stroeer/tapir/sdks

## Usage

```bash
npm i @stroeer/tapir-proto-v1
# or
yarn add @stroeer/tapir-proto-v1
```

## Development

Make sure you generated the gRPC models and services for node.

```shell
# tapir root folder
make generate TEMPLATE=buf.gen.node-proto.yaml
```

### node

```shell
# https://github.com/nvm-sh/nvm#installing-and-updating
brew install nvm

# add this to your shell/bash profile:
# source /usr/local/opt/nvm/nvm.sh

nvm install 20
nvm use 20
# optional
nvm alias system 20
```

### install

```shell
npm i
```

### Lint

```shell
npm run lint
```

## Example Script

Check the `/examples` folder.

### Build examples

```bash
npm run examples:build
```

### Run examples

Make sure that you have set the API_ENDPOINT environment variable

```bash
npm run examples:protobuf-es:run
```
