# API Status and Error Definitions

https://grpc.io/docs/guides/error/

https://cloud.google.com/apis/design/errors#error_model

https://grpc.github.io/grpc/core/md_doc_statuscodes.html

## Article Page Service Status

| Error scenario | GRPC status | HTTP Status | Cache Headers | Description |
| --- |:---:| ---:| ---:| --- |
| Article found | OK | 200
| Article not found | NOT_FOUND | 404
| Article expired | NOT_FOUND | 410
| Article not published yet | NOT_FOUND | 404
| Article could not be loaded | FAILED_PRECONDITION | 503

## Section Page Service Status

| Error scenario | GRPC status | HTTP Status | Cache Headers | Description |
| --- |:---:| ---:| ---:| --- |
| Section path is empty| INVALID_ARGUMENT | 400 
| Section path invalid | INVALID_ARGUMENT | 400
| Section path unknown | INVALID_ARGUMENT | 400 
| Curation datasource is not available | FAILED_PRECONDITION | 503
| Curated stage not found | FAILED_PRECONDITION | 400 
| Curated stage empty | FAILED_PRECONDITION | 400 
| Article loading for a curated stage failed entirely | FAILED_PRECONDITION | 503 
| Automatic stage is empty | FAILED_PRECONDITION | ??? | | Elasticsearch does not return any data for the section path 
