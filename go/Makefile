all: help

.PHONY: generate
generate: ## Generates go grpc/protobuf resources
	@echo "+ $@"
	protoc --proto_path=../proto --go_out=../ --go_opt=paths=source_relative content.proto
	protoc --proto_path=../proto --go_out=plugins=grpc:../ --go_opt=paths=source_relative service.proto

.PHONY: help
help: ## Display this help screen
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'	