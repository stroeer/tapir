#!/usr/bin/env bash
ROOT=generated/main/node
RELEASE_TAG=1.0.1
cp package.json $ROOT
cp .npmrc $ROOT
cd $ROOT || exit 99
npm pack --tag $RELEASE_TAG