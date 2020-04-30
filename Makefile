.PHONY: help clean dir node java
all: clean java node

OUTDIR = generated
PROTO_DIR = api

dir: ## Creates outdir if necessary
	[ -d $(OUTDIR) ] || mkdir $(OUTDIR)

node: dir  ## Generates node resources
	@echo "+ $@"
	for filename in $(shell find api -iname "*.proto") ; do \
		protoc --proto_path=api --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin --js_out=import_style=commonjs_strict,binary:node --ts_out=service=grpc-node:node --grpc_out=node $$filename ; \
	done

java: dir ## Generates java resources
	@echo "+ $@"
	for filename in $(shell find api -iname "*.proto") ; do \
		protoc --proto_path=api --plugin=protoc-gen-grpc=protoc-gen-grpc-java --java_out=generated --grpc_out=generated $$filename ; \
	done


clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf generated  || true
	rm -R `ls -d node/*/` || true

help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'