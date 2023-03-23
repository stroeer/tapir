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

To generate gRPC source code for t-online APIs you need to install `protoc` and gRPC plugins on your local machine,
or you can use our [buf generate template](./buf.gen.yaml) to generate code for `java`, `node` and `go`
using buf [remote plugins](https://buf.build/docs/bsr/remote-plugins/usage/):

```shell
make generate
```

### quality assurance

We use [buf](https://buf.build/) to lint our proto files and to detect breaking changes. In addition, we run some basic language specific tests to verify a
successful code generation for `java`, `node` and `go`.

Run `make` to run all checks.

### client libraries

We generate packages for [java](https://github.com/stroeer/tapir/packages/235034) and [node](https://github.com/stroeer/tapir/packages/235031)
automatically for each new tag which can be integrated in your build system. 

In addition, a go module will be generated and pushed to [go-tapir](https://github.com/stroeer/go-tapir). 

## release a new tapir version

To create a new release run `make BUMP=[major|minor|patch] release` (defaults to `patch)` in your clean main branch.
