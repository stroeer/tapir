# API Status and Error Definitions

https://grpc.io/docs/guides/error/

https://cloud.google.com/apis/design/errors#error_model

https://grpc.github.io/grpc/core/md_doc_statuscodes.html

## Article Page Service Status

| Error scenario | GRPC status | HTTP Status | Cache Headers | Description |
| --- |:---:| ---:| ---:| --- |
| Article found | OK | 200 | max-age=60
| Article not found | NOT_FOUND | 404 | max-age=60
| Article not published yet | NOT_FOUND | 404 | max-age=60
| Article expired | NOT_FOUND | 410 | max-age=60
| Article could not be loaded | FAILED_PRECONDITION | 503 | max-age=0, no-cache, no-store, must-revalidate
| Article loading timeout | FAILED_PRECONDITION | 503 | max-age=0, no-cache, no-store, must-revalidate

## Section Page Service Status

| Error scenario | GRPC status | HTTP Status | Cache Headers | Description |
| --- |:---:| ---:| ---:| --- |
| Section path is empty| INVALID_ARGUMENT | 400 | max-age=60
| Section path invalid | INVALID_ARGUMENT | 400 | max-age=60
| Section path unknown | FAILED_PRECONDITION | 404 | max-age=60
| Stage data not found | FAILED_PRECONDITION | 404 | max-age=60
| Stage data is empty | FAILED_PRECONDITION | 404 | max-age=60
| Stage datasource is not available | FAILED_PRECONDITION | 503 | max-age=0, no-cache, no-store, must-revalidate
| Stage loading timeout | FAILED_PRECONDITION | 503 | max-age=0, no-cache, no-store, must-revalidate
