# @stroeer/tapir-v1

## Usage

```bash
npm i @stroeer/tapir-v1
# or
yarn add @stroeer/tapir-v1
```

## Development

Make sure you generated the gRPC models and services for node.

### node

```bash
# https://github.com/nvm-sh/nvm#installing-and-updating
brew install nvm

# add this to your shell/bash profile:
# source /usr/local/opt/nvm/nvm.sh

nvm install 14
nvm use 14
# optional
nvm alias system 14
```

### install

```bash
npm i
```

### Test/Lint

```bash
npm run test
npm run lint
```

## How to update jest snapshots

```bash
npm run test -- --u
```

## Helpful resources

Examples for generating TypeScript types of generated services / models.

https://github.com/badsyntax/grpc-js-typescript
