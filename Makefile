PACKAGE     := github.com/janpieper/jm
CURRENT_DIR  = $(shell pwd)
BINDIR      := $(CURRENT_DIR)/bin
BINNAME     ?= jm
SHELL        = /usr/bin/env bash
VERSION      = $(shell cat ${CURRENT_DIR}/VERSION)
BUILD_DATE   = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT   = $(shell git rev-parse HEAD)
GIT_SHA      = $(shell git rev-parse --short HEAD)
GIT_TAG      = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY    = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

.PHONY: all
all: build

.PHONY: build
build:
	go build -trimpath -o $(BINDIR)/$(BINNAME) \
		-ldflags="\
			-X $(PACKAGE)/common.version=$(VERSION) \
			-X $(PACKAGE)/common.buildDate=$(BUILD_DATE) \
			-X $(PACKAGE)/common.gitCommit=$(GIT_COMMIT) \
			-X $(PACKAGE)/common.gitTreeState=$(GIT_DIRTY) \
			-X $(PACKAGE)/common.gitTag=$(GIT_TAG)" \
		./main.go

.PHONY: format
format:
	go list -f '{{.Dir}}' ./... | xargs gofmt -w

.PHONY: clean
clean:
	@rm -r $(BINDIR)

.PHONY: info
info:
	@echo "Git Tag:        ${GIT_TAG}"
	@echo "Git Commit:     ${GIT_COMMIT}"
	@echo "Git Tree State: ${GIT_DIRTY}"
