# tapir

<img src="doku/tapir.jpeg" height="200"/>

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs 
that support ~~both REST and~~ gRPC protocols. You can use these definitions with open source 
tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf)
version 3 (proto3) as their Interface Definition Language (IDL) to
define the API interface and the structure of the payload messages.

## Building

The recommended way to build the API client libraries is through Gradle. We currently support:

- Java
- Node (incl. typescript types)

## Installing Dependencies locally
### ts-protoc-gen
    npm install
  
# Generating models and services with Tapir
```shell script
  ./gradlew build
```
files will be generated in  `generated/main`



