#!/usr/bin/make

GO_BIN := $(shell command -v go 2> /dev/null)

.PHONY: build
build:
	@$(GO_BIN) build -o ./bin/testing-system ./cmd/testing-system/main.go