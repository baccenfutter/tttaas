SHELL := /bin/bash

.PHONY: all
all: setup gen build

.PHONY: setup
setup:
	go get -u goa.design/goa/v3
	go get -u goa.design/goa/v3/...

.PHONY: gen
gen:
	goa gen github.com/baccenfutter/tictactoe/design
	goa example github.com/baccenfutter/tictactoe/design

.PHONY: build
build:
	go install ./cmd/tictactoe
	go install ./cmd/tictactoe-cli
