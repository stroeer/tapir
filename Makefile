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

DIR 			= $(shell pwd)
JAVA_DIR		= ./java/src/main/java
NODE_DIR		= ./node
NODE_LEGACY_DIR		= ./node-legacy
GO_DIR			= ./gen/go
GATEWAY_DIR	= ./gen/gateway

LANGUAGE		?= java
GRPCPLUGIN		?= /usr/bin/protoc-gen-grpc-$(LANGUAGE)

PROTO_FILES		:= $(shell find stroeer -iname "*.proto" | sed "s/proto$$/$(TARGET_SUFFIX)/")
PROTOC_VERSION	?= 3.18.0
PROTOC			?= docker run --rm --volume $(DIR):$(DIR) --workdir $(DIR) ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION)

FLAGS			+= --proto_path=$(DIR)
ifeq ($(LANGUAGE),node)
	OUTPUT		:= $(NODE_DIR)
	FLAGS		+= --plugin=protoc-gen-ts=/node_modules/.bin/protoc-gen-ts
	FLAGS		+= --plugin=protoc-gen-grpc=/node_modules/.bin/grpc_tools_node_protoc_plugin
	FLAGS		+= --js_out=import_style=commonjs,binary:$(OUTPUT)
	FLAGS		+= --ts_out=service=grpc-node,mode=grpc-js:$(OUTPUT)
	FLAGS		+= --grpc_out=grpc_js:$(OUTPUT)
else ifeq ($(LANGUAGE),node-legacy)
	OUTPUT		:= $(NODE_LEGACY_DIR)
	FLAGS		+= --plugin=protoc-gen-ts=/node_modules/.bin/protoc-gen-ts
	FLAGS		+= --plugin=protoc-gen-grpc=/node_modules/.bin/grpc_tools_node_protoc_plugin
	FLAGS		+= --js_out=import_style=commonjs,binary:$(OUTPUT)
	FLAGS		+= --ts_out=service=grpc-node:$(OUTPUT)
	FLAGS		+= --grpc_out=$(OUTPUT)
else ifeq ($(LANGUAGE),go)
	OUTPUT		:= $(GO_DIR)
	FLAGS		+= --$(LANGUAGE)_out=$(OUTPUT)
	FLAGS		+= --$(LANGUAGE)_opt=module=github.com/stroeer/go-tapir
	FLAGS		+= --$(LANGUAGE)-grpc_out=$(OUTPUT)
	FLAGS		+= --$(LANGUAGE)-grpc_opt=module=github.com/stroeer/go-tapir
else
	OUTPUT		:= $(JAVA_DIR)
	FLAGS		+= --$(LANGUAGE)_out=$(OUTPUT)
	FLAGS		+=	--plugin=protoc-gen-grpc=$(GRPCPLUGIN)
	FLAGS		+= --grpc_out=$(OUTPUT)
endif

all: $(PROTO_FILES)

# Generate source files for a specific language like 'java'
%.$(TARGET_SUFFIX) : %.proto
	@echo "+ $^ ($(LANGUAGE))"
	@mkdir -p $(OUTPUT)
	$(PROTOC) $(FLAGS) $*.proto

LANGUAGES := java node node-legacy go
.PHONY: generate java node node-legacy go
generate: $(LANGUAGES) ## Generates source files for all supported languages

$(LANGUAGES):
	$(MAKE) LANGUAGE="$@"

.PHONY: lint
lint: ## Lints all proto files using https://docs.buf.build/lint-overview
	@echo "+ $@"
	@buf lint || exit 1

.PHONY: breaking
breaking: ## Detects breaking changes using https://docs.buf.build/breaking-overview
	@echo "+ $@"
	@buf breaking --against 'https://github.com/stroeer/tapir.git#branch=master' --config buf.yaml || true

.PHONY: test
test: generate go-mod ## Runs all tests
	@echo "+ $@"
	cd java && ./gradlew clean build
#	cd node && nvm use && npm run checks
#	cd node-legacy && nvm use && npm run checks
	cd go && go test -v .

.PHONY: go-mod
go-mod: ## Create tapir go module for generated code
	@echo "+ $@"
	cd $(GO_DIR) && go mod init github.com/stroeer/go-tapir && go mod tidy

.PHONY: check
check: lint breaking test ## Runs all checks

GATEWAYS := article section
$(GATEWAYS):
	@echo "+ $@"
	mkdir -p $(GATEWAY_DIR)
	$(PROTOC) --proto_path=$(DIR) --grpc-gateway_out $(GATEWAY_DIR) \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt module=github.com/stroeer/go-tapir \
		--grpc-gateway_opt standalone=true \
		--grpc-gateway_opt grpc_api_configuration=stroeer/page/$@/v1/api_config_http.yaml \
    stroeer/page/$@/v1/$@_page_service.proto

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR) || true
	rm -rf $(GO_DIR) || true
	rm -rf `find $(NODE_DIR) -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#######################
# protoc docker image #
#######################

ACTOR = $(GITHUB_ACTOR)
TOKEN = $(GITHUB_TOKEN)
.PHONY: ghcr-login
ghcr-login:
	@echo "+ $@"
	@echo $(TOKEN) | docker login ghcr.io -u $(ACTOR) --password-stdin

.PHONY: protoc-build
protoc-build: ## Build protoc docker image
	@echo "+ $@"
	@docker build \
		--tag ghcr.io/stroeer/protoc-dockerized:latest \
		--tag ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION) \
		--build-arg PROTOC_VERSION=$(PROTOC_VERSION) .

.PHONY: protoc-push
protoc-push: ghcr-login protoc-build ## Push protoc docker image to https://github.com/orgs/stroeer/packages/container/package/protoc-dockerized
	@echo "+ $@"
	@docker push --all-tags ghcr.io/stroeer/protoc-dockerized

###########
# release #
###########

DESCRIBE			= $(eval DESCRIBE := $(shell git fetch --all > /dev/null && git describe --match "v*" --always --tags))$(DESCRIBE)

DESCRIBE_PARTS		= $(subst -, ,$(DESCRIBE))
# 'v0.2.0'
VERSION_TAG			= $(word 1,$(DESCRIBE_PARTS))
# '0.2.0'
VERSION				= $(subst v,,$(VERSION_TAG))
# '0 2 0'
VERSION_PARTS		= $(subst ., ,$(VERSION))

MAJOR				= $(word 1,$(VERSION_PARTS))
MINOR				= $(word 2,$(VERSION_PARTS))
PATCH				= $(word 3,$(VERSION_PARTS))

BUMP ?= patch
ifeq ($(BUMP), major)
NEXT_VERSION		= $(shell echo $$(($(MAJOR)+1)).0.0)
else ifeq ($(BUMP), minor)
NEXT_VERSION		= $(shell echo $(MAJOR).$$(($(MINOR)+1)).0)
else
NEXT_VERSION		= $(shell echo $(MAJOR).$(MINOR).$$(($(PATCH)+1)))
endif
NEXT_TAG 			= v$(NEXT_VERSION)

.PHONY: check-git-clean
check-git-clean: ## Verifies clean working directory
	@echo "+ $@"
	git status
	git diff-index --exit-code --name-only HEAD || (echo "There are uncomitted changes"; exit 1)

.PHONY: check-git-branch
check-git-branch: check-git-clean
	@echo "+ $@"
	git fetch --all --tags --prune
	git checkout master

.PHONY: release
release: clean check check-git-branch fundoc ## Releases new version of gRPC source code packages
	@echo "+ $@ $(NEXT_TAG)"
	git tag -a $(NEXT_TAG) -m "$(NEXT_TAG)"
	git push origin $(NEXT_TAG)

.PHONY:
release-local-java: ## Releases generated Java code to your local maven repository
	cd java && ./gradlew clean build publishToMavenLocal

###########
# docs	  #
###########

.PHONY: fundoc
# to test locally, install fundoc via `cargo install fundoc`
fundoc: introduction.md ## Generate, Commit and Push documentation.
	@echo "+ $@"
  ifeq (, $(shell which fundoc))
	@echo "fundoc not installed natively, running docker build. When running locally, consider installing it."
		docker build --tag fundoc docs_resources
		docker run --rm --volume ${CURDIR}/:/opt fundoc
  else
		@echo "local fundoc installation found!"
		fundoc
  endif
	cp docs_resources/highlight.js docs
	@rm stroeer/introduction.md || true
	git add --all docs/
	git commit --message "updated docs"
	git push

introduction.md ::
	awk 'BEGINFILE{print "/**\n* @Article 00 Introduction"}{print "* " $$0} END{ print "*/"}' docs_resources/introduction.md > stroeer/$@
