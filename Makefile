SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

# go builds are fast enough that we can just build on demand instead of trying to do any fancy
# change detection
build: clean modfit
.PHONY: build

modfit:
	go build ./cmd/modfit

clean:
	rm -f ./modfit
.PHONY: clean

test:
	go test ./...
.PHONY: test
