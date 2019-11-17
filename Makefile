SHELL := /bin/bash

.PHONY: all
all: setup gen install

.PHONY: setup
setup:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u goa.design/goa/v3
	go get -u goa.design/goa/v3/...

.PHONY: gen
gen:
	goa gen github.com/baccenfutter/tictactoe/design
	goa example github.com/baccenfutter/tictactoe/design

.PHONY: install
install:
	CGO_ENABLED=0 GOOS=linux go install ./cmd/tictactoe
	CGO_ENABLED=0 GOOS=linux go install ./cmd/tictactoe-cli
