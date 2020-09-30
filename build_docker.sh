docker build -t ghcr.io/stroeer/protoc-dockerized:latest \
--build-arg NODE_GRPC_TOOLS_VERSION="1.9.0" \
--build-arg NODE_TS_PROTOC_GEN_VERSION="0.12.0" \
--build-arg JAVA_LIB_VERSION="1.29.0" \
--build-arg PROTOC_VERSION="3.11.4" \
--build-arg GO_GRPC_TOOLS_VERSION="1.32.0" \
.
