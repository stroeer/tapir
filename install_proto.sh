JAVA_LIB_VERSION=1.29.0
URL=https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/$JAVA_LIB_VERSION/protoc-gen-grpc-java-$JAVA_LIB_VERSION-linux-x86_64.exe
wget $URL -O protoc-gen-grpc-java
chmod +x protoc-gen-grpc-java

PROTOC_VERSION=3.11.4
wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip -O protoc.zip
unzip -o protoc.zip
bin/protoc --version