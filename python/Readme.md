# Setup

```shell
$ pipenv install
$ find ../stroeer -name "*.proto" -exec python -m grpc_tools.protoc \
    -I.. \
    --python_out=. \
    --grpc_python_out=. \
    {} \;
    
# try it out
$ GRPC_ENDPOINT='<HOST>:<PORT>' python sample.py

main article
============
id: 87971076
type: CONTENT_TYPE_ARTICLE
sub_type: CONTENT_SUB_TYPE_NEWS
<output truncated>


related articles
================
[
  { 
    article {<output truncated>}
    source: RELATED_ARTICLE_SOURCE_EDITORIAL
  },
  { 
    article {<output truncated>}
    source: RELATED_ARTICLE_SOURCE_EDITORIAL
  }
]
  

```
