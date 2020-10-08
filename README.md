<div align="center">
  <img src="doku/tapir.png" height="200" alt="tapir"/>
   <h1>tapir</h1>
</div>

![GitHub release (latest by date)](https://img.shields.io/github/v/release/stroeer/tapir?style=flat-square)

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs that support ~~both REST and~~ gRPC protocols. You can use these definitions with open source tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf) version 3 (proto3) as their Interface Definition Language (IDL) to define the API interface and the structure of the payload messages.

## Build

We use [stroeer/protoc-dockerized](https://github.com/orgs/stroeer/packages/container/package/protoc-dockerized) as `protoc` which currently supports generating code for

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

#### generate

For generating go sources run:

```shell script
$ make LANGUAGE=go
```

#### release

Generating go code is currently not part of the release process. Go sources need
to be generated locally and added to pull requests.

## docker Image for `stroeer/protoc-dockerized`

### requirements

Login to `ghcr.io` with your Github user name and a Github personal access token having permissions to `read:packages`, `write:packages` and/or `delete:packages`:

```sh
cat ${TOKEN_FILE_LOCATION} | docker login ghcr.io -u ${GH_USERNAME} --password-stdin
```

### release new version

- bump versions:
  - protoc version in `Makefile`
  - go dependency versions in `go.mod`
  - java dependency versions in `build.gradle`
  - node depenedncy versions in `package.json`
- run `make image-build` (uses `$protoc_version` as well as `latest` as image tags)
- run `make image-release` to release new images
