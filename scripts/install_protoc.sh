#!/usr/bin/env bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip -O protoc.zip
unzip -o protoc.zip bin/protoc
bin/protoc --version