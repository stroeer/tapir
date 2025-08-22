<div align="center">
  <img src="docs/tapir.png" height="200" alt="tapir"/>
   <h1>tapir</h1>
</div>

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/stroeer/tapir?color=%23f653a6&label=Release&style=flat-square) 
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/stroeer/tapir/proto.yaml?style=flat-square&label=build) 
![GitHub License](https://img.shields.io/github/license/stroeer/tapir?style=flat-square) 

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


## generate code

We recommend to use the generated SDKs from the [buf schema registry (BSR)](#buf-schema-registry) in your project. 

In addition, it's possible to generate code for interacting with the t-online APIs using the [buf cli](https://buf.build/docs/generate/overview). See [docs](https://buf.build/docs/installation)
for installation instructions.  

Run the following commands to generate code for `java`, `python`, `node` (deprecated), `protobuf-es` (`node-buf`) or `go`:

```shell
# java, python and go
make generate
# node (deprecated, use BSR sdk instead or node-buf if cjs is needed)
make generate TEMPLATE=buf.gen.node.yaml
# bufbuild/es generates cjs only
make generate TEMPLATE=buf.gen.node-buf.yaml
```

See corresponding `buf.gen.*.yaml` code generation configuration.

## testing and quality assurance

We use the [buf cli](https://buf.build/docs/ecosystem/cli-overview) for quality assurance of our proto files:

* `make lint`: linting all proto files with [buf lint](https://buf.build/docs/lint/overview) 
* `make fmt`: formatting all proto files with [buf fmt](https://buf.build/docs/format/style)
* `make breaking`: check for breaking changes against the main branch with [buf breaking](https://buf.build/docs/breaking/overview)

These tools are configured in the `buf.yaml` file. 

In addition, we run some basic language specific tests to verify a successful code generation for `java` and `node`.

## releases

We use [semantic versioning](https://semver.org/) for our releases. 

To create a new release run`make BUMP=[major|minor|patch] release` (defaults to `patch)` in your clean main branch. This will create 
a new tag and push it to the main branch. In addition, a new release will be created in GitHub if a
fine-grained personal access token is provided in the `GITHUB_TOKEN` environment variable.

### buf schema registry

We push to the [buf schema registry (BSR)](https://buf.build/stroeer/tapir) automatically for each new tag. The registry
provides SDKs for various languages to interact with the t-online APIs. It's recommended to 
use these SDKs instead of our custom [client libraries](#client-libraries).

### client libraries

> [!IMPORTANT]  
> It is recommended to use the [BSR SDKs](#buf-schema-registry) instead of the client libraries below. 
> The client libraries are deprecated and will not be maintained in the future.

In addition to the [BSR SDKs](https://buf.build/stroeer/tapir), we generate packages hosted on [GitHub](https://github.com/orgs/stroeer/packages?repo_name=tapiro) 
for the following languages:

* (`deprecated`) [node](https://github.com/stroeer/tapir/packages/235031)
* [protobuf-es (v2)](https://github.com/stroeer/tapir/pkgs/npm/tapir-buf-v1)

automatically for each new tag which can be integrated in your build system.

See our GitHub [workflow](.github/workflows/proto.yaml) for details. 



 
