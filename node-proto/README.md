# @stroeer/tapir-proto-v1

This package only includes all tapir models. Without the grpc services.
It's based on the packages `@bufbuild/protobuf`.

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

### Lint

```shell
npm run lint
```

## Example Script

Check the `/examples` folder.
