# @stroeer/tapir-v1

This package includes all tapir models + grpc services for node.js + typescript definitions.
It's based on the official packages `google-protobuf` + `@grpc-js/grpc`.

## Usage

```bash
npm i @stroeer/tapir-v1
# or
yarn add @stroeer/tapir-v1
```

## Development

Make sure you generated the gRPC models and services for node.

```shell
# tapir root folder
make generate TEMPLATE=buf.gen.node.yaml
```

### node

```shell
# https://github.com/nvm-sh/nvm#installing-and-updating
brew install nvm

# add this to your shell/bash profile:
# source /usr/local/opt/nvm/nvm.sh

nvm install 18
nvm use 18
# optional
nvm alias system 18
```

### install

```shell
npm i
```

### Test/Lint

```shell
npm run test
npm run lint
```

### How to update test snapshots

```bash
npm run test -- --u
```

## Helpful resources

Examples for generating TypeScript types of generated services / models.

https://github.com/badsyntax/grpc-js-typescript

## Example Script

Check the `/examples` folder.

### Build examples

```bash
npm run examples:build
```

### Run examples

Make sure that you have set the GRPC_HOST and GRPC_API_TOKEN environment variables

```bash
npm run examples:grpc-js:run
```
