# API Status and Error Definitions

https://grpc.io/docs/guides/error/

https://cloud.google.com/apis/design/errors#error_model

https://grpc.github.io/grpc/core/md_doc_statuscodes.html

https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto

## Article Page Service Status

| scenario | description | gRPC status | gRPC error payload | HTTP Status | cacheable
|:---|:---|:---:|:---|:---:|:---:|
| found | article was found in the datastore and is published according to it's [Metadata](https://github.com/stroeer/tapir/blob/d99e155003a5f81a4f300d558d1fbc736dbda511/stroeer/core/v1/article.proto#L87)| OK | none | 200 | yes
| invalid id | article id is invalid | INVALID_ARGUMENT | [google.rpc.BadRequest](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L169) | 400 | yes
| not found | article was not found in the datastore | NOT_FOUND | [google.rpc.ResourceInfo](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L198) | 404 | yes
| not published yet | article was found in the datastore but it's `Metadata` [start_time](https://github.com/stroeer/tapir/blob/d99e155003a5f81a4f300d558d1fbc736dbda511/stroeer/core/v1/article.proto#L93) is not reached, yet | NOT_FOUND | [google.rpc.ResourceInfo](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L198) with info about the `start_time` | 404 | yes
| expired | article was found in the datastore but is expired according to it's `Metadata` [end_time](https://github.com/stroeer/tapir/blob/d99e155003a5f81a4f300d558d1fbc736dbda511/stroeer/core/v1/article.proto#L97) | NOT_FOUND | [google.rpc.ResourceInfo](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L198) with info about the `end_time` | 404/410 (TODO: clarify requirements with SEO) | yes
| deleted/archived (TODO: clarify requirements with SEO) | article was found in the datastore but it's [state](https://github.com/stroeer/tapir/blob/d99e155003a5f81a4f300d558d1fbc736dbda511/stroeer/core/v1/article.proto#L89) is marked as deleted/archived | FAILED_PRECONDITION | [google.rpc.PreconditionFailure](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L143) with info about the [state](https://github.com/stroeer/tapir/blob/d99e155003a5f81a4f300d558d1fbc736dbda511/stroeer/core/v1/article.proto#L89) | 410 | yes
| internal | internal error processing the article | INTERNAL | none | 500 | no
| timeout | timeout loading and processing the article | DEADLINE_EXCEEDED | none | 504 | no

**Redirect cases should be defined after requirements are clear**

## Section Page Service Status

| scenario | description | gRPC status | gRPC error payload | HTTP Status | cacheable |
|:---|:---|:---:|:---|:---:|:---:|
| found | all data for the section page was found | OK | none | none | yes |
| section path is empty | client did not provide a section path | INVALID_ARGUMENT | [google.rpc.BadRequest](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L169) | 400 | yes
| section path is invalid | client provided an invalid section path | INVALID_ARGUMENT | [google.rpc.BadRequest](https://github.com/googleapis/googleapis/blob/3e9b231b098d514c9a51ba8c59e416f45837b195/google/rpc/error_details.proto#L169) | 400 | yes
| section path is unknown | client provided an unknown section path | NOT_FOUND | none | 404 | yes
| internal | internal error while loading section data | INTERNAL | none | 500 | no
| timeout | timeout while loading section data | DEADLINE_EXCEEDED | none | 504 | no

**Scenarios about incomplete section data needs to be defined**
