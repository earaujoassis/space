#!/usr/bin/env bash

set -ex


go get golang.org/x/lint/golint
test -z "$(golint ./...)"

yarn lint
