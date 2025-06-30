#!/usr/bin/env bash

set -ex

git config core.hooksPath .githooks

npm install -g yarn
yarn install && yarn build

go install github.com/go-task/task/v3/cmd/task@latest
task setup
task setup:jwks
task --list-all

echo "> Setup completed"
