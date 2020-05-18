DIR = $(shell pwd)
JAVA_DIR = ./java/src/main/java
NODE_DIR = ./node

OUTPUT ?= $(JAVA_DIR)
LANGUAGE ?= java
GRPCPLUGIN ?= /usr/bin/protoc-gen-grpc-$(LANGUAGE)
# with https://github.com/TheThingsIndustries/docker-protobuf/issues/9 we could use this image also for TypeScript
PROTOC ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) thethingsindustries/protoc

FLAGS+= --proto_path=.
ifeq ($(LANGUAGE),node)
FLAGS+= --plugin=protoc-gen-ts=$(NODE_DIR)/node_modules/.bin/protoc-gen-ts
FLAGS+= --plugin=protoc-gen-grpc=$(NODE_DIR)/node_modules/.bin/grpc_tools_node_protoc_plugin
FLAGS+= --js_out=import_style=commonjs,binary:$(NODE_DIR)
FLAGS+= --ts_out=service=grpc-node:$(NODE_DIR)
else
FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)
FLAGS+=	--plugin=protoc-gen-grpc=$(GRPCPLUGIN)
endif
FLAGS+= --grpc_out=$(OUTPUT)

all: clean stroeer/core stroeer/article

stroeer/%: $(OUTPUT)
	@echo "+ $@"
	$(PROTOC) $(FLAGS) $(shell find $@ -iname "*.proto" -exec echo -n '"{}" ' \; | tr '\n' ' ')

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR)  || true
	rm -R `find node -type d \( -iname "*" ! -iname "node_modules" ! -iname "tests" \) -mindepth 1 -maxdepth 1` || true

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
