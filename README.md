<div align="center">
  <img src="doku/tapir.png" height="200" alt="tapir"/>
   <h1>tapir</h1>
</div>

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs
that support ~~both REST and~~ gRPC protocols. You can use these definitions with open source
tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf) version 3 (proto3) 
as their Interface Definition Language (IDL) to define the API interface and the structure of the 
payload messages.

## building

we currently support generating and packaging code for `java` and `node/TypeScript`.

### java

#### generate
 
We use [docker-protobuf](https://github.com/TheThingsIndustries/docker-protobuf) as `protoc`. For generating
Java sources run:

```bash
$ make
```

#### release

The recommended way to build and publish the Java client libraries is through gradle:

```shell script
$ export GITHUB_USERNAME=foo
$ export GITHUB_TOKEN=bar
$ cd java && ./gradlew publish
``` 

### node

#### generate

Since [docker-protobuf](https://github.com/TheThingsIndustries/docker-protobuf) doesn't support TypeScript 
([yet](https://github.com/TheThingsIndustries/docker-protobuf/issues/9)), you need to install the required 
binaries for your environment first. For MacOs run:

```bash
$ brew install protobuf
$ npm --prefix node install
```

Then run:

```bash
$ make LANGUAGE=node OUTPUT=./node PROTOC=protoc
```

#### release

TODO

(Login to github packages for `npm publish` is done via the actions/setup-node@v1 which writes a .npmrc file with the token from the secrets)


