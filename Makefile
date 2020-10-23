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
DOCS_DIR = ./api-docs

OUTPUT ?= $(JAVA_DIR)
LANGUAGE ?= java
GRPC_PLUGIN ?= /usr/bin/protoc-gen-grpc-$(LANGUAGE)

PROTOC_VERSION ?= 3.13.0

PROTOC ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION)

FLAGS+= --proto_path=$(DIR)
FLAGS+= --doc_opt=markdown,package.md
ifeq ($(LANGUAGE),node)
	ifeq ($(OUTPUT), $(JAVA_DIR))
		OUTPUT = $(NODE_DIR)
	endif
	FLAGS+= --plugin=protoc-gen-ts=/node_modules/.bin/protoc-gen-ts
	FLAGS+= --plugin=protoc-gen-grpc=/node_modules/.bin/grpc_tools_node_protoc_plugin
	FLAGS+= --js_out=import_style=commonjs,binary:$(OUTPUT)
	FLAGS+= --ts_out=service=grpc-node,mode=grpc-js:$(OUTPUT)
	FLAGS+= --grpc_out=grpc_js:$(OUTPUT)
else ifeq ($(LANGUAGE),go)
	ifeq ($(OUTPUT), $(JAVA_DIR))
		OUTPUT = $(GO_DIR)
	endif
	FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)
	FLAGS+= --$(LANGUAGE)_opt=paths=source_relative
	FLAGS+= --$(LANGUAGE)-grpc_out=$(OUTPUT)
	FLAGS+= --$(LANGUAGE)-grpc_opt=paths=source_relative
else
	FLAGS+= --$(LANGUAGE)_out=$(OUTPUT)
	FLAGS+=	--plugin=protoc-gen-grpc=$(GRPC_PLUGIN)
	FLAGS+= --grpc_out=$(OUTPUT)
endif

.PHONY: clean help image-build image-release stroeer/%

all: stroeer/*

doc-folders: clean-docs
	@echo "+ docsdir $@"
	find stroeer/* -type d -depth 0 | xargs -L1 -I{} mkdir -p $(DOCS_DIR)/"{}"

$(OUTPUT):
	@echo "+ mkdir $@"
	[ -d $@ ] || mkdir -p $@


stroeer/%: $(OUTPUT) doc-folders
	@echo "+ compile: $@"
	$(PROTOC) $(FLAGS) --doc_out=$(DOCS_DIR)/$@ $(shell find $@ -iname "*.proto" -exec echo -n '"{}" ' \; | tr '\n' ' ')

clean: ## Deletes all generated files
	@echo "+ $@"
	@rm -rf $(JAVA_DIR) || true
	@rm -rf `find $(GO_DIR) -type d \( -iname "*" ! -iname "go.mod" ! -iname "go.sum" ! -iname "*_test.go" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true
	@rm -rf `find $(NODE_DIR) -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

clean-docs:
	@if [[ $(DOCS_DIR)/* == /* ]]; then echo "var DOCS_DIR is unset, prevented 'rm -rf /*'"; else rm -rf $(DOCS_DIR)/*; fi

docker-build: ## Build new docker image
	docker build \
		-t ghcr.io/stroeer/protoc-dockerized:latest \
		-t ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION) \
		--build-arg PROTOC_VERSION=$(PROTOC_VERSION) .

docker-push: image-build ## Release new docker image to GHCR
	docker push ghcr.io/stroeer/protoc-dockerized:latest
	docker push ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION)

help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dev-release-java: clean java stroeer/*
	cd java; ./gradlew publishToMavenLocal