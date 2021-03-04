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

PROTO_FILES := $(shell find stroeer -iname "*.proto" | sed "s/proto$$/$(SUFFIX)/")
PROTOC_VERSION ?= 3.15.0
PROTOC ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION)

FLAGS+= --proto_path=$(DIR)
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
	FLAGS+=	--plugin=protoc-gen-grpc=$(GRPCPLUGIN)
	FLAGS+= --grpc_out=$(OUTPUT)
endif

.PHONY: clean help image-build image-release gateway lint java go node

all: $(PROTO_FILES)

# Generate source files for a specific language like 'java'
%.$(SUFFIX): %.proto
	@echo "+ $^ ($(LANGUAGE))"
	@mkdir -p $(OUTPUT)
	$(PROTOC) $(FLAGS) $*.proto

LANGUAGES ?= java node go
generate: $(LANGUAGES) ## Generates source files for all supported languages

$(LANGUAGES):
	$(MAKE) LANGUAGE="$@"

lint: ## Lints all proto files using https://docs.buf.build/lint-overview
	@echo "+ $@"
	@buf lint || exit 1

breaking: ## Detects breaking changes using https://docs.buf.build/breaking-overview
	@echo "+ $@"
	@buf breaking --against 'https://github.com/stroeer/tapir.git#branch=master' --config buf.yaml || true

test: generate # Runs all tests
	@echo "+ $@"
	cd java && ./gradlew clean build
	cd node && npm run checks
	cd go && go test -v .

check: lint breaking test ## Runs all checks

gateway: ## Generates grpc-gateway resources
	@echo "+ $@"
	$(PROTOC) -I . --grpc-gateway_out $(GO_DIR) \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt grpc_api_configuration=stroeer/page/article/v1/api_config_http.yaml \
		stroeer/page/article/v1/article_page_service.proto

clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR) || true
	rm -rf `find $(GO_DIR) -type d \( -iname "*" ! -iname "go.mod" ! -iname "go.sum" ! -iname "*_test.go" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true
	rm -rf `find $(NODE_DIR) -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

ACTOR = $(GITHUB_ACTOR)
TOKEN = $(GITHUB_TOKEN)
ghcr-login:
	@echo "+ $@"
	@echo $(TOKEN) | docker login ghcr.io -u $(ACTOR) --password-stdin

protoc-build: ## Build protoc docker image
	@echo "+ $@"
	@docker build \
		-t ghcr.io/stroeer/protoc-dockerized:latest \
		-t ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION) \
		--build-arg PROTOC_VERSION=$(PROTOC_VERSION) .

protoc-push: ghcr-login protoc-build ## Push protoc docker image to https://github.com/orgs/stroeer/packages/container/package/protoc-dockerized
	@echo "+ $@"
	@docker push ghcr.io/stroeer/protoc-dockerized:latest
	@docker push ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION)

help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
