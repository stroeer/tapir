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

LANGUAGE		?= java
GRPCPLUGIN		?= /usr/bin/grpc_$(LANGUAGE)_plugin

PROTOC_VERSION								?= 21.7
GRPC_VERSION 									= 1.49.1
GRPC_JAVA_VERSION 						= 1.50.0
PROTOBUF_JS_VERSION 					= 3.21.2
GO_PROTOC_GEN_GO_VERSION 			= v1.28.1
GO_PROTOC_GEN_GO_GRPC_VERSION = v1.2.0
GRPC_GATEWAY_VERSION 					= v2.11.3

PROTO_FILES	:= $(shell find stroeer -iname "*.proto" | sed "s/proto$$/$(TARGET_SUFFIX)/")
PROTOC ?= docker run --rm --volume $(DIR):$(DIR) --workdir $(DIR) ghcr.io/stroeer/protoc-dockerized:$(PROTOC_VERSION)

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
	@buf breaking --against 'https://github.com/stroeer/tapir.git#branch=main' --config buf.yaml || true

.PHONY: test
test: generate ## Runs all tests
	@echo "+ $@"
	cd java && ./gradlew clean build
#	cd node && nvm use && npm run checks
#	cd node-legacy && nvm use && npm run checks

.PHONY: check
check: lint breaking test ## Runs all checks

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf $(JAVA_DIR) || true
	rm -rf $(GO_DIR) || true
	rm -rf `find $(NODE_DIR) -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" ! -iname "src" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true
	rm -rf `find $(NODE_LEGACY_DIR) -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" ! -iname "src" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

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
		--build-arg PROTOC_VERSION=$(PROTOC_VERSION) \
		--build-arg GRPC_JAVA_VERSION=$(GRPC_JAVA_VERSION) \
		--build-arg PROTOBUF_JS_VERSION=$(PROTOBUF_JS_VERSION) \
		--build-arg GRPC_VERSION=$(GRPC_VERSION) \
		--build-arg GO_PROTOC_GEN_GO_GRPC_VERSION=$(GO_PROTOC_GEN_GO_GRPC_VERSION) \
		--build-arg GO_PROTOC_GEN_GO_VERSION=$(GO_PROTOC_GEN_GO_VERSION) \
		--build-arg GRPC_GATEWAY_VERSION=$(GRPC_GATEWAY_VERSION) .

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
	git checkout main

.PHONY: release
release: clean check check-git-branch commit_fundoc ## Releases new version of gRPC source code packages
	@echo "+ $@ $(NEXT_TAG)"
	git tag -a $(NEXT_TAG) -m "$(NEXT_TAG)"
	git push origin $(NEXT_TAG)
	if [ ! -z "${GITHUB_TOKEN}" ] ; then 								\
    		curl 																					\
    			-H "Authorization: token ${GITHUB_TOKEN}" 	\
    			-X POST 																		\
    			-H "Accept: application/vnd.github.v3+json"	\
    			https://api.github.com/repos/stroeer/tapir/releases \
    			-d "{\"tag_name\":\"$(NEXT_TAG)\",\"generate_release_notes\":true}"; \
    fi;

.PHONY:
release-local-java: java ## Releases generated Java code to your local maven repository
	cd java && ./gradlew clean build publishToMavenLocal

###########
# docs	  #
###########

.PHONY: fundoc
# to test locally, install fundoc via `cargo install fundoc`
fundoc :: introduction.md ## Generate, Commit and Push documentation.
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

commit_fundoc :: fundoc
	git add --all docs/
	git commit --message "updated docs" || true # in case nothing changes
	git push

introduction.md ::
	awk 'BEGINFILE{print "/**\n* @Article 00 Introduction"}{print "* " $$0} END{ print "*/"}' docs_resources/introduction.md > stroeer/$@
