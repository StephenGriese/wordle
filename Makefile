NAME := messi
DESC := Kafka Messages to XMP ESS
PREFIX ?= usr/local
VERSION := $(shell git describe --tags --always --dirty --match "[0-9]*.[0-9]*.[0-9]*")
GOVERSION := 1.19.0
BUILDVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
PROJECT_URL := "https://github.com/comcast-pulse/$(NAME)"
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(BUILDVERSION)' \
           -X 'main.name=$(NAME)'
DOCKER_REPO := "atlas-registry.sys.comcast.net"

build: staticcheck lint test clean target/local

init:
	git config core.hooksPath .githooks

staticcheck:
	staticcheck ./...

lint:
	golangci-lint -v run ./...

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
	@test -n "$(WORDLE_DICTIONARY)" || (echo "WORDLE_DICTIONARY not set. Run: export WORDLE_DICTIONARY=./wordle.txt" && exit 1)
	@echo "Starting Wordle Helper server..."
	@go run ./cmd/server/main.go

run-server-dev:
	@export WORDLE_DICTIONARY=./wordle.txt && \
	export WORDLE_REMOVE=./words-to-remove && \
	export WORDLE_PORT=8080 && \
	go run ./cmd/server/main.go

.PHONY: clean lint modules build run-server run-server-dev target/server
