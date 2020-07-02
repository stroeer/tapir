FROM node:lts-stretch as installer

ARG PROTOC_VERSION
ARG NODE_GRPC_TOOLS_VERSION
ARG NODE_TS_PROTOC_GEN_VERSION
ARG JAVA_LIB_VERSION

# node plugins
RUN npm install grpc-tools@$NODE_GRPC_TOOLS_VERSION \
                ts-protoc-gen@$NODE_TS_PROTOC_GEN_VERSION

RUN mkdir /installer

# Download Java Lib
RUN wget https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/$JAVA_LIB_VERSION/protoc-gen-grpc-java-$JAVA_LIB_VERSION-linux-x86_64.exe -O /installer/protoc-gen-grpc-java

# protoc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -O protoc.zip
RUN mkdir -p /installer/protoc
RUN unzip -o protoc.zip -d /installer/protoc/

FROM node:lts-stretch-slim
USER root
COPY --from=installer /node_modules /node_modules
COPY --from=installer /installer/protoc /usr/bin/protoc
COPY --from=installer /installer/protoc-gen-grpc-java /usr/bin/

RUN chmod +x /usr/bin/protoc-gen-grpc-java

ENV PATH="/usr/bin/protoc/bin:${PATH}"
RUN protoc --version

ENTRYPOINT [ "protoc" ]
