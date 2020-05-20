# This is a simple Makefile that generates client library source code
# for Tapir using Protocol Buffers and gRPC for any supported
# language. However, it does not compile the generated code into final
# libraries that can be directly used with application code.
#
# Supported (tested) languages: java, node (TypeScript), go
#
# Syntax example:
#		make LANGUAGE=node OUTPUT=./node
#

DIR = $(shell pwd)
JAVA_DIR = ./java/src/main/java
NODE_DIR = ./node
GO_DIR = ./go

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
FLAGS+= --grpc_out=$(OUTPUT)
else ifeq ($(LANGUAGE),go)
FLAGS+= --$(LANGUAGE)_out=plugins=grpc:$(OUTPUT)
else
# java, swift, php, js
FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)
FLAGS+=	--plugin=protoc-gen-grpc=$(GRPCPLUGIN)
FLAGS+= --grpc_out=$(OUTPUT)
endif

all: clean stroeer/*

$(OUTPUT):
	@echo "+ $@"
	[ -d $@ ] || mkdir -p $@

stroeer/%: $(OUTPUT)
	@echo "+ $@"
	$(PROTOC) $(FLAGS) $(shell find $@ -iname "*.proto" -exec echo -n '"{}" ' \; | tr '\n' ' ')

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR)  || true
	rm -rf $(GO_DIR)  || true
	rm -R `find node -type d \( -iname "*" ! -iname "node_modules" ! -iname "tests" \) -mindepth 1 -maxdepth 1` || true

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
