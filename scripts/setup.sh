#!/usr/bin/env bash

set -ex

git config core.hooksPath .githooks

npm install -g yarn
npm install -g allure-commandline
yarn install && yarn build

go install honnef.co/go/tools/cmd/staticcheck@latest
go install gotest.tools/gotestsum@latest
go install github.com/go-task/task/v3/cmd/task@latest
task setup
task --list-all

echo "> Setup completed"
