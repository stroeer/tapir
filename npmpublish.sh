#!/usr/bin/env bash
ROOT=generated/main/ts
cp package.json $ROOT
cd $ROOT || exit 99
npm pack