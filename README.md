<div align="center">
  <img src="docs/tapir.png" height="200" alt="tapir"/>
   <h1>tapir</h1>
</div>

![GitHub release (latest by date)](https://img.shields.io/github/v/release/stroeer/tapir?style=flat-square)

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs that support the [gRPC](https://grpc.io/) protocol. You can use these definitions with open source tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf) version 3 (proto3) as their Interface Definition Language (IDL) to define the API interface and the structure of the payload messages.

## Guidelines

* tapir provides an [IDL](https://en.wikipedia.org/wiki/Interface_description_language) and RCP services stubs to access editorial content and their configuration. This allows delivering various t-online products developed by independent teams
* RPC services and proto messages are optimized for an efficient development and delivery of those products: One of our internal API guideline demands that all the content required to render a page must depend on a single API call.
* different editorial content types and product features are modeled by the same proto messages using the same concepts. This allows to keep the APIs _clean_ and future proof. Examples:
    * articles, videos or galleries share the same message structure distinguished by a type field
    * attributes of an article are modeled as generic `<string, string>` maps
    * elements of an article like images, videos and their assets share the same proto messages and can be distinguished by a type field
    * [enumerations](https://developers.google.com/protocol-buffers/docs/proto3#enum) are only used for stable/rarely changing lists of pre-defined values like a content type.
    Volatile fields like layout types are modeled as `string` fields.   
* proto message fields and entries of maps are optional unless commented otherwise. Clients must not break if an optional field or map entry is missing.
 

## Code generation

We provide [stroeer/protoc-dockerized](https://github.com/orgs/stroeer/packages/container/package/protoc-dockerized) as `protoc` 
and a Makefile tested to generate code for `java`, `node (TypeScript)` and `go`:

```shell script
$ make LANGUAGE=[java,node,go]
``` 

## Client libraries

[Releases](https://github.com/stroeer/tapir/releases) include client libraries as `java` and `npm` [packages](https://github.com/orgs/stroeer/packages?repo_name=tapir). 

Generating go code is currently not part of the release process. Go sources need to be generated locally and added to pull requests.

## Docker image

### Version management

bump versions in a PR and merge into `master`:

- protoc version in `Makefile`
- go dependency versions in `go.mod`
- java dependency versions in `build.gradle`
- node dependency versions in `package.json`

### Release

Login to `ghcr.io` with your Github user name and a Github personal access token having permissions to `read:packages`, `write:packages` and/or `delete:packages`:

```sh
cat ${TOKEN_FILE_LOCATION} | docker login ghcr.io -u ${GH_USERNAME} --password-stdin
```

- run `make image-build` (tags as [`$protoc_version`, `latest`])
- run `make image-release` to push the new image
