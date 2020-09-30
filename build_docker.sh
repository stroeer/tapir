# https://github.com/grpc/grpc-node/releases
# https://github.com/improbable-eng/ts-protoc-gen/releases
# https://github.com/grpc/grpc-java/releases
# https://github.com/protocolbuffers/protobuf/releases
# https://github.com/grpc/grpc-go

docker build -t stroeer/protoc-dockerized:latest \
--build-arg NODE_GRPC_TOOLS_VERSION="1.9.1" \
--build-arg NODE_TS_PROTOC_GEN_VERSION="0.13.0" \
--build-arg JAVA_LIB_VERSION="1.32.1" \
--build-arg PROTOC_VERSION="3.13.0" \
--build-arg GO_GRPC_TOOLS_VERSION="1.32.0" \
.
