<div align="center">
  <img src="docs_resources/tapir.png" height="200" alt="tapir"/>
   <h1>tapir</h1>
</div>

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/stroeer/tapir?color=%23f653a6&label=Release&style=flat-square)

The **T**-online **API** **R**epository contains the interface definitions of t-online APIs that support the [gRPC](https://grpc.io/) protocol. You can use these definitions with open source tools to generate client libraries, documentation and other artifacts.

T-online APIs use [Protocol Buffers](https://github.com/google/protobuf) version 3 (proto3) as their Interface Definition Language (IDL) to define the API interface and the structure of the payload messages.

## overview

### guidelines

* tapir provides an [IDL](https://en.wikipedia.org/wiki/Interface_description_language) and RCP services stubs to access editorial content and their configuration. This allows delivering various t-online products developed by independent teams
* RPC services and proto messages are optimized for an efficient development and delivery of those products: One of our internal API guideline demands that all the content required to render a page must depend on a single API call.
* different editorial content types and product features are modeled by the same proto messages using the same concepts. This allows to keep the APIs _clean_ and future proof. Examples:
    * articles, videos or galleries share the same message structure distinguished by a type field
    * attributes of an article are modeled as generic `<string, string>` maps
    * elements of an article like images, videos and their assets share the same proto messages and can be distinguished by a type field
    * [enumerations](https://developers.google.com/protocol-buffers/docs/proto3#enum) are only used for stable/rarely changing lists of pre-defined values like a content type.
    Volatile fields like layout types are modeled as `string` fields.
* proto message fields and entries of maps are optional unless commented otherwise. Clients must not break if an optional field or map entry is missing.


## generate gRPC source code

To generate gRPC source code for t-online APIs you need to install `protoc` and gRPC on your local machine,
or you can use our [protoc docker image](#protoc-docker-image) which includes all required plugins for `java`, `node` and `go` source code
generation.

Then you can run `make LANGUAGE=xxx` to generate the source code for a specific language.

It's also possible to generate gRPC source code for `java`, `node`, `node-legacy` and `go` at once: `make generate`.

### quality assurance

We use [buf](https://buf.build/) to lint our proto files and to detect breaking changes. In addition, we run some basic language specific tests to verify a
successful code generation for `java`, `node`, `node-legacy` and `go`.

Run `make check` to run all checks.

### client libraries

We generate packages for [java](https://github.com/stroeer/tapir/packages/235034) and [node](https://github.com/stroeer/tapir/packages/235031)
automatically for each new tag which can be integrated in your build system. 

In addition, a go module will be generated and pushed to [go-tapir](https://github.com/stroeer/go-tapir). 

## release a new tapir version

To create a new release run `make BUMP=[major|minor|patch] release` (defaults to `patch)` in your clean master branch.

## protoc docker image

We provide [stroeer/protoc-dockerized](https://github.com/orgs/stroeer/packages/container/package/protoc-dockerized) including `protoc` and all required
grpc plugins to generate source code for `java`, `node`, `node-legacy` and `go`. This docker image also supports generating a [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
reverse-proxy server.

### precondition

To access our docker image you need a valid GitHub PAT (personal access token) and login to `https://ghcr.io`. You have to do this only once.

1. create (or re-use) a GitHub PAT in your [GitHub settings](https://github.com/settings/tokens)
2. follow the login instructions via [GitHub docs](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-docker-registry#authenticating-with-a-personal-access-token)
3. :warning: login URL is `https://ghcr.io`. Example: `cat ~/TOKEN.txt | docker login https://ghcr.io -u [GITHUB-USERNAME] --password-stdin`

### release

1. bump versions

    - protoc version in `Makefile` - this version will be used as the docker image tag
    - go dependency versions in `go/go.mod` (run `go mod tidy -compat=1.17` afterwards)
    - java dependency versions in `build.gradle`
    - node dependency versions in `package.json`


2. Export actor and token for the [GitHub container registry](https://docs.github.com/en/packages/guides/about-github-container-registry)

   - `export GITHUB_ACTOR={yourusername}`
   - `export GITHUB_TOKEN={yourtoken}` (needs `read:packages`, `write:packages`, `delete:packages` permissions)


3. build and push protoc docker image
    - `make protoc-push`
