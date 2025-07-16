all: build generate test fmt lint breaking

TEMPLATE ?= buf.gen.yaml
ifeq ($(findstring node,$(TEMPLATE)),node)
	# Also generate all imports (except for Well-Known Types) like google/api/http.proto
  FLAGS = --include-imports
else
  FLAGS =
endif

.PHONY: generate
generate: ## Generates proto and grpc files using https://docs.buf.build/generate/overview
	@echo "+ $@"
	@buf generate --template $(TEMPLATE) $(FLAGS)

.PHONY: lint
lint: ## Lints all proto files using https://docs.buf.build/lint-overview
	@echo "+ $@"
	@buf lint || exit 1

.PHONY: breaking
breaking: ## Detects breaking changes using https://docs.buf.build/breaking-overview
	@echo "+ $@"
	@buf breaking --against 'https://github.com/stroeer/tapir.git#branch=main' --config buf.yaml || true

.PHONY: fmt
fmt: ## Formats all proto files using https://docs.buf.build/format/style
	@echo "+ $@"
	@buf format -w

.PHONY: build
build: ## Builds buf image, see https://buf.build/docs/reference/images
	@echo "+ $@"
	@buf build -o tapir.pb

.PHONY: push
push: build ## Pushes tapir to the buf schema registry, see https://buf.build/docs/bsr/introduction and https://buf.build/docs/bsr/module/publish#pushing-with-labels
	@echo "+ $@"
	@buf push --git-metadata

.PHONY: test
test: generate ## Runs all tests
	@echo "+ $@"
	@cd java && ./gradlew clean build
	@cd node && npm run checks

.PHONY: node-examples
node-examples:  ## Tests generated node client stubs against API
	@echo "+ $@"
	@cd node && npm run examples:build  && npm run examples:grpc-js:run

.PHONY: node-proto-examples
node-proto-examples:  ## Tests generated node-proto client stubs against API
	@echo "+ $@"
	@cd node-proto && npm run examples:build  && npm run examples:protobuf-es:run

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf ./java/src/main/java || true
	rm -rf ./gen || true
	rm -rf `find ./node -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" ! -iname "examples" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true
	rm -rf `find ./node-proto -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" ! -iname "examples" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true
	rm -rf `find ./node-buf -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" ! -iname "examples" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

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
release: clean check-git-branch commit_fundoc ## Releases new version of gRPC source code packages
	@gh --version >/dev/null 2>&1 || (echo "ERROR: gh CLI is required (try running 'brew install gh')."; exit 1)
	@gh auth status >/dev/null 2>&1 || (echo "ERROR: gh auth status is NOT OK (try running 'gh auth status')."; exit 1)
	@echo "+ $@ $(NEXT_TAG)"
	git tag -a $(NEXT_TAG) -m "$(NEXT_TAG)"
	git push origin $(NEXT_TAG)
	gh release create --generate-notes --notes --target $(NEXT_TAG) --verify-tag

.PHONY:
release-local-java: java ## Releases generated Java code to your local maven repository
	cd java && ./gradlew clean build publishToMavenLocal

postman :: ## Generate postman root.proto than can be imported into postman v10 gRPC APIs
	echo 'syntax = "proto3";' > root.proto
	echo 'package stroeer;' >> root.proto
	find stroeer -name '*service.proto' -type f -exec echo "import \"{}\";" >> root.proto \;
