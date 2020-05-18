# mark targets to avoid confusion with files
.PHONY: clean dir help java node
all: clean java node

JAVA_DIR = ./java/src/main/gen
PROTO_DIR = ./api
NODE_DIR = ./node
PROTOC_BIN ?= protoc
ALL_PROTO_FILES = $(shell find api -iname "*.proto" -exec echo -n '"{}" ' \; | tr '\n' ' ')

dir: ## Creates outdir if necessary
	[ -d $(JAVA_DIR) ] || mkdir -p $(JAVA_DIR)

node:  ## Generates node resources
	@echo "+ $@"
	$(PROTOC_BIN) --proto_path=$(PROTO_DIR) \
		--plugin=protoc-gen-ts=$(NODE_DIR)/node_modules/.bin/protoc-gen-ts \
		--plugin=protoc-gen-grpc=$(NODE_DIR)/node_modules/.bin/grpc_tools_node_protoc_plugin \
		--js_out=import_style=commonjs,binary:$(NODE_DIR) \
		--ts_out=service=grpc-node:$(NODE_DIR) \
		--grpc_out=$(NODE_DIR) $(ALL_PROTO_FILES)


java: dir ## Generates java resources
	@echo "+ $@"
	$(PROTOC_BIN) --proto_path=$(PROTO_DIR) \
		--plugin=protoc-gen-grpc=protoc-gen-grpc-java \
		--java_out=$(JAVA_DIR) \
		--grpc_out=$(JAVA_DIR) $(ALL_PROTO_FILES)

clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR)  || true
	rm -R `find node -type d \( -iname "*" ! -iname "node_modules" ! -iname "tests" \) -mindepth 1 -maxdepth 1` || true

help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
