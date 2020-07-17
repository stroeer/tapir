<div align="center">
  <img src="doku/tapir.png" height="200" alt="tapir"/>
   <h1>tapir</h1>
</div>

![GitHub release (latest by date)](https://img.shields.io/github/v/release/stroeer/tapir?style=flat-square)

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs that support ~~both REST and~~ gRPC protocols. You can use these definitions with open source tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf) version 3 (proto3) as their Interface Definition Language (IDL) to define the API interface and the structure of the payload messages.

## building

We use [stroeer/protoc-dockerized](https://hub.docker.com/repository/docker/stroeer/protoc-dockerized) as `protoc` which currently supports generating code for

- `java`
- `node/TypeScript`
- `go`

(the `protoc` binary is configurable in the `Makefile` though)

### java

#### generate

For generating Java sources run:

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

For generating node sources run:

```bash
$ make LANGUAGE=node
```

#### release

TODO

(Login to github packages for `npm publish` is done via the actions/setup-node@v1 which writes a .npmrc file with the token from the secrets)

### go

Generating go code is currently not part of the CI and release process.

#### setup

```shell script
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go get github.com/grpc/grpc-go/cmd/protoc-gen-go-grpc
```

#### generate

```shell script
$ make LANGUAGE=go
```

### Docker Image for `stroeer/protoc-dockerized`

#### Requirements
- be member of https://hub.docker.com/u/stroeer
- login to hub.docker.com with docker login

#### Release new version
- look at Requirements above
- bump versions in `build_docker.sh`
- run `build_docker.sh`
- run `docker push stroeer/protoc-dockerized:latest`


