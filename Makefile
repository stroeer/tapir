# This is a simple Makefile that generates client library source code
# for Tapir using Protocol Buffers and gRPC for any supported
# language. However, it does not compile the generated code into final
# libraries that can be directly used with application code.
#
# Supported (tested) languages: java, node (TypeScript), go
#
# Syntax example:
#		make LANGUAGE=node
#

DIR = $(shell pwd)
JAVA_DIR = ./java/src/main/java
NODE_DIR = ./node
GO_DIR = ./go

OUTPUT ?= $(JAVA_DIR)
LANGUAGE ?= java
GRPCPLUGIN ?= /usr/bin/protoc-gen-grpc-$(LANGUAGE)
PROTOC ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) stroeer/protoc-dockerized

FLAGS+= --proto_path=$(DIR)
ifeq ($(LANGUAGE),node)
	ifeq ($(OUTPUT), $(JAVA_DIR))
		OUTPUT = $(NODE_DIR)
	endif
	FLAGS+= --plugin=protoc-gen-ts=/node_modules/.bin/protoc-gen-ts
	FLAGS+= --plugin=protoc-gen-grpc=/node_modules/.bin/grpc_tools_node_protoc_plugin
	FLAGS+= --js_out=import_style=commonjs,binary:$(OUTPUT)
	FLAGS+= --ts_out=service=grpc-node:$(OUTPUT)
	FLAGS+= --grpc_out=$(OUTPUT)
else ifeq ($(LANGUAGE),go)
	ifeq ($(OUTPUT), $(JAVA_DIR))
		OUTPUT = $(GO_DIR)
	endif
	FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)
	FLAGS+= --go_opt=paths=source_relative
	FLAGS+= --go-grpc_out=$(OUTPUT)
	FLAGS+= --go-grpc_opt=paths=source_relative
else
	FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)
	FLAGS+=	--plugin=protoc-gen-grpc=$(GRPCPLUGIN)
	FLAGS+= --grpc_out=$(OUTPUT)
endif

.PHONY: clean help stroeer/%

all: stroeer/*

$(OUTPUT):
	@echo "+ $@"
	[ -d $@ ] || mkdir -p $@

stroeer/%: $(OUTPUT)
	@echo "+ $@"
	$(PROTOC) $(FLAGS) $(shell find $@ -iname "*.proto" -exec echo -n '"{}" ' \; | tr '\n' ' ')

clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR) || true
	rm -rf `find $(GO_DIR) -type d \( -iname "*" ! -iname "go.mod" ! -iname "go.sum" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true
	rm -rf `find $(NODE_DIR) -type d \( -iname "*" ! -iname "node_modules" ! -iname "tests" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
