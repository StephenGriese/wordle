NAME := wordle
DESC := Wordle web app
VERSION := $(shell git describe --tags --always --dirty --match "[0-9]*.[0-9]*.[0-9]*")
BUILDVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(BUILDVERSION)' \
           -X 'main.name=$(NAME)'

# Sentinel file — proves `make init` has been run on this clone.
HOOKS_SENTINEL=.git/hooks/.githooks-installed

.PHONY: init checks staticcheck lint test build clean modules run-server run-server-dev

# Guard target — aborts with a helpful message if init was never run.
$(HOOKS_SENTINEL):
	@echo ""
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "⚠️  Run 'make init' first to configure git hooks."
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@exit 1

init:
	git config core.hooksPath .githooks
	@touch $(HOOKS_SENTINEL)
	@echo "✓ Git hooks configured. You're ready to develop."

build: checks clean target/local

checks: $(HOOKS_SENTINEL) staticcheck lint test

staticcheck:
	staticcheck ./...

lint:
	golangci-lint -v run --fix ./...

test:
	go test ./...

target/local: modules
	mkdir -p target/local/bin && go build -ldflags "$(LDFLAGS)" -o target/local/bin/wordle ./cmd/cli

target/server: modules
	mkdir -p target/local/bin && go build -ldflags "$(LDFLAGS)" -o target/local/bin/wordle-server ./cmd/server

modules:
	go mod tidy

clean-target:
	rm -rf target

clean: clean-target

run-server:
	@test -n "$(WORDLE_DICTIONARY)" || (echo "WORDLE_DICTIONARY not set. Run: export WORDLE_DICTIONARY=./american-english" && exit 1)
	@echo "Starting Wordle Helper server..."
	@go run ./cmd/server/main.go

run-server-dev:
	@export WORDLE_DICTIONARY=./american-english && \
	export WORDLE_REMOVE=./words-to-remove && \
	export WORDLE_PORT=8080 && \
	go run ./cmd/server/main.go

