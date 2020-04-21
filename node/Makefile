all: help

.PHONY: generate
generate: ## Generates node grpc/protobuf resources
	@echo "+ $@"
	protoc --proto_path=../proto content.proto \
		--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
		--js_out=import_style=commonjs_strict,binary:. \
		--ts_out=.
	protoc --proto_path=../proto service.proto \
		--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
		--plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
		--js_out=import_style=commonjs_strict,binary:. \
		--ts_out=service=grpc-node:. \
		--grpc_out=.

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -f *.js *.d.ts

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'