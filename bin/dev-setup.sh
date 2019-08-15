#!/usr/bin/env bash

set -ex

git config core.hooksPath .githooks
npm install -g yarn
yarn install && yarn build
go get github.com/mattn/goreman

echo "> Setup completed. You may run 'go run main.go serve' to start the server"
