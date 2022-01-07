SHELL = bash

# Release tag
RELEASE_TAG := $(shell git tag -l --points-at HEAD)

# Version of current or last release
VERSION := $(shell cat VERSION)
GO_VERSION := $(shell go version)

.PHONY: all build run version

all:
	gofmt -s -w .
	go build -v --ldflags "-s -X github.com/vsoch/comp/version.Version=$(VERSION)" -o comp
	
build:
	go build -v --ldflags "-s -X github.com/vsoch/comp/version.Version=$(VERSION)" -o comp
	
run:
	go run main.go
	
version:
	@echo '$(VERSION)'

update:
	GO111MODULE=on go get -d -u -t ./...
	GO111MODULE=on go mod tidy

tidy:
	GO111MODULE=on go mod tidy
