#!/usr/bin/env bash
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org,direct

export ENV=development

go build -v -o $(ROOT)/bin/ar-cli $(ROOT)/cmd/ar-cli/*.go
