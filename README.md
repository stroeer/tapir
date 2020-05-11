# tapir

<img src="doku/tapir.jpeg" height="200"/>

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs 
that support ~~both REST and~~ gRPC protocols. You can use these definitions with open source 
tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf)
version 3 (proto3) as their Interface Definition Language (IDL) to
define the API interface and the structure of the payload messages.

## Building

We currently support:

- Java
- Node (incl. typescript types)

## Installing Dependencies locally
### protoc
    brew install protobuf
### ts-protoc-gen
    npm --prefix node install
  
# Generating models and services with Tapir
```shell script
  make all
```

files will be generated in folders `src` for java and `node` for the node resources

## Continuous Integration
- Login to github packages for `npm publish` is done via the actions/setup-node@v1 which writes a .npmrc file with the token from the secrets
- login to github packages for `gradlew publish` is done via env vars which are accessed in the `build.gradle`


