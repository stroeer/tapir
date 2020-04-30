all: java ts grpc

OUTDIR = generated
PROTO_DIR = api

.PHONY: dir
dir: ## Creates outdir if necessary
	[ -d $(OUTDIR) ] || mkdir $(OUTDIR)

.PHONY: ts
ts: dir ## Generates node typescript resources
	@echo "+ $@"
	for filename in $(shell find api -iname "*.proto") ; do \
		protoc --proto_path=api --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs_strict,binary:node --ts_out=node $$filename ; \
	done
.PHONY: grpc
grpc: dir  ## Generates javascript grpc/protobuf resources
	@echo "+ $@"
	for filename in $(shell find api -iname "*.proto") ; do \
		protoc --proto_path=api --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin --js_out=import_style=commonjs_strict,binary:node --ts_out=service=grpc-node:node --grpc_out=node $$filename ; \
	done

.PHONY: java
java: dir ## Generates java resources
	@echo "+ $@"
	for filename in $(shell find api -iname "*.proto") ; do \
		protoc --proto_path=api --plugin=protoc-gen-grpc=protoc-gen-grpc-java --java_out=generated --grpc_out=generated $$filename ; \
	done


.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf article
	rm -rf section
	rm -rf generated
	rm -r node/article node/section

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'