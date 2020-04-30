#!/usr/bin/env bash
VERSION=1.29.0
URL=https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/$VERSION/protoc-gen-grpc-java-$VERSION-linux-x86_64.exe
wget $URL > protoc-gen-grpc-java
chmod +x protoc-gen-grpc-java