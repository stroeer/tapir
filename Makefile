all: generate test fmt lint breaking

TEMPLATE ?= buf.gen.yaml
.PHONY: generate
generate: ## Generates proto and grpc files using https://docs.buf.build/generate/overview
	@echo "+ $@"
	@buf generate --template $(TEMPLATE)

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

.PHONY: test
test: generate ## Runs all tests
	@echo "+ $@"
	@cd java && ./gradlew clean build
	@cd node && npm run checks

.PHONY: clean
clean: ## Deletes all generated files
	@echo "+ $@"
	rm -rf ./java/src/main/java || true
	rm -rf ./gen || true
	rm -rf `find ./node -type d \( -iname "*" ! -iname "node_modules" ! -iname "__tests__" ! -iname "src" \) -mindepth 1 -maxdepth 1` 2> /dev/null || true

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
	git commit --message "docs released for new tag: $(NEXT_TAG)" || true # in case nothing changes
	git push

introduction.md ::
	awk 'BEGINFILE{print "/**\n* @Article 00 Introduction"}{print "* " $$0} END{ print "*/"}' docs_resources/introduction.md > stroeer/$@
