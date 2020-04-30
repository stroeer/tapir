#!/usr/bin/env bash
ROOT=generated/main
cp package.json $ROOT
cp index.js $ROOT
cp .npmrc $ROOT
cd $ROOT || exit 99
npm pack --tag $RELEASE_TAG